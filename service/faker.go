package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"time"
)

var Fake = newFaker() // Tests can access this instance directly

// Faker is a structure that can generate randomized fake data
type Faker struct {
	seed      int64
	envConfig *Config
}

// newFaker will generate a new Faker object and seed the rand package.
func newFaker() *Faker {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	config, err := NewConfig() // Gets config from environment
	if err != nil {
		panic(err)
	}
	return &Faker{seed: seed, envConfig: config}
}

// integer will return an integer from the rand min - max. Max is not inclusive
func (fake *Faker) integer(min int, max int) int {
	return min + rand.Intn(max-min)
}

// func testID will return a random test ID
func (fake *Faker) testID() uint64 {
	return uint64(fake.integer(1, 1000000))
}

// testSummary will return a random test summary
func (fake *Faker) testSummary() string {
	summaries := []string{
		"Ensures the /metadata endpoint is functional",
		"Checks that a valid input produces a valid output",
		"User service load test",
		"Navbar component link positive test",
		"Ensures that publishing a valid Kafka event gets consumed correctly downstream",
		"Ensures a bad input returns a correct error message",
		"Verifies that bad data does not get forwarded downstream",
		"Test user insert query is functional",
	}
	summary := summaries[fake.integer(0, len(summaries))]
	return summary
}

// testOutcome will return a random test outcome
func (fake *Faker) testOutcome() Outcome {
	outcomes := []Outcome{
		Passed,
		Failed,
	}
	outcome := outcomes[fake.integer(0, len(outcomes))]
	return outcome
}

// testAnalysis will return a random test analysis. Takes in an outcome, because valid analyses vary based on the
// outcome of the test. Pass in a nil outcome to get any random testAnalysis
func (fake *Faker) testAnalysis(outcome *Outcome) Analysis {
	var validAnalyses []Analysis

	if *outcome == Passed {
		validAnalyses = []Analysis{
			NotAnalyzed,
			TrueNegative,
			FalseNegative,
		}
	} else if *outcome == Failed {
		validAnalyses = []Analysis{
			NotAnalyzed,
			TruePositive,
			FalsePositive,
		}
	} else if outcome == nil {
		validAnalyses = []Analysis{
			NotAnalyzed,
			TruePositive,
			FalsePositive,
			TrueNegative,
			FalseNegative,
		}
	} else {
		panic(fmt.Errorf("error with testAnalysis parameter, must be a valid outcome or nil! Got %s", *outcome))
	}
	analysis := validAnalyses[fake.integer(0, len(validAnalyses))]
	return analysis
}

// testResolution will return a random test resolution
func (fake *Faker) testResolution() Resolution {
	analyses := []Resolution{
		Unresolved,
		NotNeeded,
		QuickFix,
		TicketCreated,
		TestFixed,
		TestDisabled,
		KnownIssue,
	}
	resolution := analyses[fake.integer(0, len(analyses))]
	return resolution
}

// testDoc will generate a random testDoc. For simplicity, these docs are finite and hard-coded.
func (fake *Faker) testDoc() map[string]any {
	docs := []map[string]any{
		{
			"app":   "user-service",
			"type":  "integration",
			"owner": "Patrick Star",
			"testPayload": map[string]any{
				"id":            1,
				"username":      "someUser48",
				"accountStatus": "lock",
			},
			"testResponse": map[string]any{
				"responseCode": 200,
				"responseBody": nil,
			},
		},
		{
			"owner":         "Sandy Cheeks",
			"type":          "UI",
			"browsers":      []string{"chrome", "firefox", "edge"},
			"screenshotURL": "https://some-s3-bucket-that-doesnt-exist.com/714029473432412",
		},
		{
			"owner":   "Squidward Tentacles",
			"type":    "load",
			"maxRPS":  300,
			"service": "application-service",
			"samplePayloads": []map[string]any{
				{
					"app_id": "47324033",
					"status": "APPROVED",
				},
				{
					"app_id": "9948302",
					"status": "REJECTED",
				},
			},
			"runtime": "10m",
			"latency (ms)": map[string]float32{
				"p50": 254.33,
				"p75": 332.45,
				"p95": 501.99,
				"p99": 676.51,
			},
		},
	}
	doc := docs[fake.integer(0, len(docs))]
	return doc
}

// test will generator a random, valid models.Test object and return a pointer to it.
func (fake *Faker) test() *Test {
	outcome := fake.testOutcome()
	analysis := fake.testAnalysis(&outcome)
	resolution := fake.testResolution()

	test := &Test{
		ID:         fake.testID(),
		Summary:    fake.testSummary(),
		Outcome:    outcome,
		Analysis:   analysis,
		Resolution: resolution,
		Created:    time.Now(),
		Modified:   time.Now(),
		Doc:        fake.testDoc(),
	}
	return test
}

// ginContext will return a pointer to a fake gin.Context object for testing
func (fake *Faker) ginContext() *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, err := gin.CreateTestContext(w)
	if err != nil {
		panic(err)
	}
	return c
}

// pgPool will return a real pgx.ConnPool because I do not think it is valuable to mock it out
func (fake *Faker) pgPool() *pgx.ConnPool {
	pgPool, err := NewPGPool(fake.envConfig.PG)
	if err != nil {
		panic(err)
	}
	return pgPool
}

// testController will return a fake TestController with a pgPool connection
func (fake *Faker) testController() *TestController {
	controller := &TestController{DBPool: fake.pgPool()}
	return controller
}

// testRequest will return a fake Test http.Request that can be sent through a testController
func (fake *Faker) testRequest(method string, test *Test, withID bool) *http.Request {
	body := gin.H{
		"summary":    test.Summary,
		"outcome":    test.Outcome,
		"analysis":   test.Analysis,
		"resolution": test.Resolution,
	}
	if withID {
		body["ID"] = test.ID
	}

	// Right-merges doc
	if test.Doc != nil && len(test.Doc) > 0 {
		for k, v := range test.Doc {
			body[k] = v
		}
	}

	jsonValue, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(method, "/test", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	return req
}

// multiple will call a specific fakeMethod function n times and return the results as a slice
func multiple[T any](n int, fakerMethod func() T) []T {
	sl := make([]T, n)
	for i := 0; i < n; i++ {
		sl[i] = fakerMethod()
	}
	return sl
}
