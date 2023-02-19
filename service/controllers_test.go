package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
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
