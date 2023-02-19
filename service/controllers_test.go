package main

import (
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

// TestTestController_CreateTest will
func TestTestController_CreateTest(t *testing.T) {
	controller := Fake.testController()

	t.Run("valid test returns valid response", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request = Fake.testRequest("POST")
		controller.CreateTest(c)

		t.Log(w.Body)
		assert.Equal(t, 201, w.Code)
	})
}
