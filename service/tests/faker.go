package tests

import (
	"fmt"
	"github.com/ryandem1/oar/models"
	"math/rand"
	"time"
)

// Faker is a structure that can generate randomized fake data
type Faker struct {
	seed int64
}

// newFaker will generate a new Faker object and seed the rand package.
func newFaker() *Faker {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return &Faker{seed: seed}
}

// integer will return an integer from the rand min - max. Max is not inclusive
func (fake *Faker) integer(min int, max int) int {
	return min + rand.Intn(max-min)
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
func (fake *Faker) testOutcome() models.Outcome {
	outcomes := []models.Outcome{
		models.Passed,
		models.Failed,
	}
	outcome := outcomes[fake.integer(0, len(outcomes))]
	return outcome
}

// testAnalysis will return a random test analysis. Takes in an outcome, because valid analyses vary based on the
// outcome of the test. Pass in a nil outcome to get any random testAnalysis
func (fake *Faker) testAnalysis(outcome *models.Outcome) models.Analysis {
	var validAnalyses []models.Analysis

	if *outcome == models.Passed {
		validAnalyses = []models.Analysis{
			models.NotAnalyzed,
			models.TrueNegative,
			models.FalseNegative,
		}
	} else if *outcome == models.Failed {
		validAnalyses = []models.Analysis{
			models.NotAnalyzed,
			models.TruePositive,
			models.FalsePositive,
		}
	} else if outcome == nil {
		validAnalyses = []models.Analysis{
			models.NotAnalyzed,
			models.TruePositive,
			models.FalsePositive,
			models.TrueNegative,
			models.FalseNegative,
		}
	} else {
		panic(fmt.Errorf("error with testAnalysis parameter, must be a valid outcome or nil! Got %s", *outcome))
	}
	analysis := validAnalyses[fake.integer(0, len(validAnalyses))]
	return analysis
}

// testResolution will return a random test resolution
func (fake *Faker) testResolution() models.Resolution {
	analyses := []models.Resolution{
		models.Unresolved,
		models.NotNeeded,
		models.QuickFix,
		models.TicketCreated,
		models.TestFixed,
		models.TestDisabled,
		models.KnownIssue,
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
func (fake *Faker) test() *models.Test {
	outcome := fake.testOutcome()
	analysis := fake.testAnalysis(&outcome)
	resolution := fake.testResolution()

	test := &models.Test{
		ID:         int64(fake.integer(1, 1000000)),
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

// multiple will call a specific fakeMethod function n times and return the results as a slice
func multiple[T any](n int, fakerMethod func() T) []T {
	sl := make([]T, n)
	for i := 0; i < n; i++ {
		sl[i] = fakerMethod()
	}
	return sl
}
