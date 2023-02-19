package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestTestController_CreateTest will ensure that the CreateTest controller accepts valid Test objects and does not
// accept invalid tests
func TestTestController_CreateTest(t *testing.T) {
	controller := Fake.testController()

	t.Run("valid test returns valid response", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = Fake.testRequest("POST", Fake.test())
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
	}
	for scenario, invalidTest := range invalidTests {
		t.Run(scenario, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = Fake.testRequest("POST", invalidTest)
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
	body := []gin.H{
		{"ID": testID},
		{"ID": testID2},
	}
	jsonValue, err := json.Marshal(body)
	if err != nil {
		t.Error("setup error", err)
	}

	t.Run("delete tests that exist", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/tests", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Error("setup error", err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = req
		controller.DeleteTests(c)

		assert.Equal(t, w.Code, 200)
	})

	t.Run("delete tests that don't exist", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/tests", bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Error("setup error", err)
		}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = req
		controller.DeleteTests(c)

		t.Log(w.Code)
		assert.Equal(t, w.Code, 200)
	})
}
