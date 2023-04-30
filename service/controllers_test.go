package main

import (
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/magiconair/properties/assert"
	"net/http"
	"strconv"
	"testing"
	"time"
)

// TestTestController_CreateTest will ensure that the CreateTest controller accepts valid Test objects and does not
// accept invalid tests
func TestTestController_CreateTest(t *testing.T) {
	controller := Fake.testController()

	t.Run("valid test returns valid response", func(t *testing.T) {
		c, w := Fake.ginContext()

		c.Request = Fake.testRequest(http.MethodPost, Fake.test(), "/test")
		controller.CreateTest(c)

		assert.Equal(t, 201, w.Code)
	})

	outcome := Fake.testOutcome()
	analysis := Fake.testAnalysis(&outcome)
	resolution := Fake.testResolution()

	invalidTests := map[string]*Test{
		"blank summary": {
			ID:         Fake.testID(),
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid outcome": {
			ID:         Fake.testID(),
			Summary:    Fake.testSummary(),
			Outcome:    "Skipped",
			Analysis:   analysis,
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid analysis": {
			ID:         Fake.testID(),
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   "Some Analysis",
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid resolution": {
			ID:         Fake.testID(),
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: "Some resolution",
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"doc field defined": {
			ID:         Fake.testID(),
			Summary:    Fake.testSummary(),
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: Fake.testResolution(),
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc: map[string]any{
				"doc": "something",
			},
		},
	}
	for scenario, invalidTest := range invalidTests {
		t.Run(scenario, func(t *testing.T) {
			c, w := Fake.ginContext()

			c.Request = Fake.testRequest(http.MethodPost, invalidTest, "/test")
			controller.CreateTest(c)

			assert.Equal(t, 400, w.Code)
		})
	}
}

// TestTestController_DeleteTests will ensure that you can delete tests and that deleting non-existing tests does not
// throw an error
func TestTestController_DeleteTests(t *testing.T) {
	controller := Fake.testController()

	testID, err := InsertTest(Fake.pgPool(), Fake.test())
	testID2, err := InsertTest(Fake.pgPool(), Fake.test())

	query := TestQuery{
		IDs:            []uint64{testID, testID2},
		Summaries:      nil,
		Outcomes:       nil,
		Analyses:       nil,
		Resolutions:    nil,
		CreatedBefore:  nil,
		CreatedAfter:   nil,
		ModifiedBefore: nil,
		ModifiedAfter:  nil,
		Docs:           nil,
	}
	encodedQuery, err := encodeToBase64(query)
	if err != nil {
		t.Error(err)
	}

	t.Run("delete tests that exist", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c, w := Fake.ginContext()

		c.Request = req
		controller.DeleteTests(c)
		assert.Equal(t, w.Code, 200)
	})

	t.Run("delete tests that don't exist", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c, w := Fake.ginContext()

		c.Request = req

		controller.DeleteTests(c)

		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("delete with no query", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/tests", nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c, w := Fake.ginContext()

		c.Request = req
		controller.DeleteTests(c)

		assert.Equal(t, w.Code, 400)
	})
}

// TestTestController_PatchTest will ensure that PatchTest works with valid tests and rejects invalid tests
func TestTestController_PatchTest(t *testing.T) {
	controller := Fake.testController()
	testID, err := InsertTest(Fake.pgPool(), Fake.test())
	if err != nil {
		t.Error("setup error", err)
	}

	tests, err := SelectTests(Fake.pgPool(), "select * from oar_tests where id=$1", testID)
	if err != nil {
		t.Error("setup error", err)
	}
	test := tests[0]

	query := TestQuery{
		IDs:            []uint64{testID},
		Summaries:      nil,
		Outcomes:       nil,
		Analyses:       nil,
		Resolutions:    nil,
		CreatedBefore:  nil,
		CreatedAfter:   nil,
		ModifiedBefore: nil,
		ModifiedAfter:  nil,
		Docs:           nil,
	}
	encodedQuery, err := encodeToBase64(query)
	if err != nil {
		t.Error(err)
	}

	t.Run("valid test returns valid response", func(t *testing.T) {
		c, w := Fake.ginContext()

		c.Request = Fake.testRequest(http.MethodPatch, test, "/tests?query="+encodedQuery)
		controller.PatchTests(c)

		assert.Equal(t, w.Code, 200)
	})

	outcome := Fake.testOutcome()
	analysis := Fake.testAnalysis(&outcome)
	resolution := Fake.testResolution()

	invalidTests := map[string]*Test{
		"blank summary": {
			ID:         test.ID,
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid outcome": {
			ID:         test.ID,
			Summary:    Fake.testSummary(),
			Outcome:    "Skipped",
			Analysis:   analysis,
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid analysis": {
			ID:         test.ID,
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   "Some Analysis",
			Resolution: resolution,
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid resolution": {
			ID:         test.ID,
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: "Some resolution",
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"invalid ID": {
			ID:         0,
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: "Some resolution",
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc:        nil,
		},
		"doc field added": { // Tests cannot define the "doc" field due to how the double bind works
			ID:         0,
			Summary:    "    ",
			Outcome:    outcome,
			Analysis:   analysis,
			Resolution: "Some resolution",
			Created:    time.Now(),
			Modified:   time.Now(),
			Doc: map[string]any{
				"doc": "something",
			},
		},
	}
	for scenario, invalidTest := range invalidTests {
		t.Run(scenario, func(t *testing.T) {
			c, w := Fake.ginContext()

			c.Request = Fake.testRequest(http.MethodPatch, invalidTest, "/tests?query="+encodedQuery)
			controller.PatchTests(c)

			assert.Equal(t, w.Code, 400)
		})
	}

	t.Run("no update returns 304", func(t *testing.T) {
		query = TestQuery{
			IDs:            nil,
			Summaries:      []string{"I am fairly certain this summary will not exist"},
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err = encodeToBase64(query)
		if err != nil {
			t.Error(err)
		}

		c, w := Fake.ginContext()

		testPatch := Test{
			Summary: "Update",
		}
		c.Request = Fake.testRequest(http.MethodPatch, &testPatch, "/tests?query="+encodedQuery)
		controller.PatchTests(c)

		assert.Equal(t, w.Code, 304)
	})

	t.Run("invalid pool test", func(t *testing.T) {
		pgConnConfig := pgx.ConnConfig{
			Host:     EnvConfig.PG.Host,
			Port:     EnvConfig.PG.Port,
			Database: "postgres",
			User:     EnvConfig.PG.User,
			Password: EnvConfig.PG.Pass,
			LogLevel: EnvConfig.PG.LogLevel,
		}
		pgConnPoolConfig := pgx.ConnPoolConfig{
			ConnConfig:     pgConnConfig,
			MaxConnections: EnvConfig.PG.PoolSize,
			AfterConnect:   nil,
			AcquireTimeout: EnvConfig.PG.PollTimeout,
		}
		badPool, err := pgx.NewConnPool(pgConnPoolConfig)
		if err != nil {
			t.Error("setup error", err)
		}
		controller = &TestController{DBPool: badPool}
		c, w := Fake.ginContext()

		c.Request = Fake.testRequest(http.MethodPatch, test, "/tests?query="+encodedQuery)
		controller.PatchTests(c)

		assert.Equal(t, w.Code, 400)
	})
}

// TestTestController_GetTests will ensure that the GetTests controller is querying correctly
func TestTestController_GetTests(t *testing.T) {
	controller := Fake.testController()
	numTests := 5 // Amount of tests to make for the GetTests test

	generatedTests := make([]*Test, 0, numTests)
	testIDs := make([]uint64, 0, numTests)

	for i := 0; i < numTests; i++ {
		generatedTests = append(generatedTests, Fake.test())
		testID, err := InsertTest(Fake.pgPool(), generatedTests[i])
		if err != nil {
			t.Error("setup error", err)
		}
		testIDs = append(testIDs, testID)
	}

	t.Run("valid query returns valid response", func(t *testing.T) {
		c, w := Fake.ginContext()

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		returnedTestIDs := make([]uint64, numTests, numTests)
		for i, test := range queryResponse.Tests {
			returnedTestIDs[len(queryResponse.Tests)-1-i] = test.ID
		}

		assert.Equal(t, returnedTestIDs, testIDs)
	})

	t.Run("limit and offset work", func(t *testing.T) {
		c, w := Fake.ginContext()

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(
			http.MethodGet,
			"/tests?query="+encodedQuery+"&limit="+strconv.Itoa(numTests-2)+"&offset=3",
			nil,
		)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, len(queryResponse.Tests), numTests-3)
		assert.Equal(t, queryResponse.Count, uint64(numTests-3))
	})

	t.Run("oar filtering works", func(t *testing.T) {
		c, w := Fake.ginContext()

		outcomeFilter := generatedTests[0].Outcome
		analysisFilter := generatedTests[0].Analysis
		resolutionFilter := generatedTests[0].Resolution

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       []string{string(outcomeFilter)},
			Analyses:       []string{string(analysisFilter)},
			Resolutions:    []string{string(resolutionFilter)},
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		for _, test := range queryResponse.Tests {
			assert.Equal(t, test.Outcome, outcomeFilter)
			assert.Equal(t, test.Analysis, analysisFilter)
			assert.Equal(t, test.Resolution, resolutionFilter)
		}
	})

	t.Run("oar summary filtering works", func(t *testing.T) {
		c, w := Fake.ginContext()

		summaryFilter := generatedTests[0].Summary

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      []string{generatedTests[0].Summary},
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		for _, test := range queryResponse.Tests {
			assert.Equal(t, test.Summary, summaryFilter)
		}
	})

	t.Run("filter limit works", func(t *testing.T) {
		c, w := Fake.ginContext()

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?limit=1001&query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 400)
	})

	t.Run("created before filter work", func(t *testing.T) {
		c, w := Fake.ginContext()

		tomorrow := time.Now().AddDate(0, 0, 1)

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  &tomorrow,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, len(queryResponse.Tests), numTests)
		assert.Equal(t, queryResponse.Count, uint64(numTests))
	})

	t.Run("created after filter work", func(t *testing.T) {
		c, w := Fake.ginContext()

		yesterday := time.Now().AddDate(0, 0, -1)

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   &yesterday,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, len(queryResponse.Tests), numTests)
		assert.Equal(t, queryResponse.Count, uint64(numTests))
	})

	t.Run("modified before filter work", func(t *testing.T) {
		c, w := Fake.ginContext()

		tomorrow := time.Now().AddDate(0, 0, 1)

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: &tomorrow,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, len(queryResponse.Tests), numTests)
		assert.Equal(t, queryResponse.Count, uint64(numTests))
	})

	t.Run("modified after filter work", func(t *testing.T) {
		c, w := Fake.ginContext()

		yesterday := time.Now().AddDate(0, 0, -1)

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  &yesterday,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, len(queryResponse.Tests), numTests)
		assert.Equal(t, queryResponse.Count, uint64(numTests))
	})

	t.Run("docs filter work", func(t *testing.T) {
		c, w := Fake.ginContext()

		genTest := generatedTests[0]
		genTest2 := generatedTests[1]

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   nil,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           []map[string]any{genTest.Doc, genTest2.Doc},
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		for _, test := range queryResponse.Tests {
			passed := fmt.Sprint(test.Doc) == fmt.Sprint(genTest.Doc) || fmt.Sprint(test.Doc) == fmt.Sprint(genTest2.Doc)
			assert.Equal(t, passed, true)
		}
	})

	t.Run("empty query doesn't fail", func(t *testing.T) {
		c, w := Fake.ginContext()

		tomorrow := time.Now().AddDate(0, 0, 1)

		query := TestQuery{
			IDs:            testIDs,
			Summaries:      nil,
			Outcomes:       nil,
			Analyses:       nil,
			Resolutions:    nil,
			CreatedBefore:  nil,
			CreatedAfter:   &tomorrow,
			ModifiedBefore: nil,
			ModifiedAfter:  nil,
			Docs:           nil,
		}
		encodedQuery, err := encodeToBase64(query)
		if err != nil {
			t.Error("setup error")
		}

		req, err := http.NewRequest(http.MethodGet, "/tests?query="+encodedQuery, nil)
		if err != nil {
			t.Error("setup error", err)
		}

		c.Request = req

		controller.GetTests(c)
		assert.Equal(t, w.Code, 200)

		var queryResponse TestQueryResponse

		err = json.Unmarshal(w.Body.Bytes(), &queryResponse)
		if err != nil {
			t.Error("response error", err)
		}

		assert.Equal(t, queryResponse.Count, uint64(0))
	})
}
