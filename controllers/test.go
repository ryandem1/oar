package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"oar/models"
	"oar/utils"
)

var tests []*models.Test // Temp store, will implement DB later

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func CreateTest(c *gin.Context) {
	test, err := utils.DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	test.ID = len(tests) + 1
	test.Clean()

	if err := test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	tests = append(tests, test)
	c.JSON(http.StatusCreated, test)
}

// PatchTest will perform a patch (partial update) operation on an existing test if it exists. Because of the nature of
// the Test enrichment process, I imagine this will be used more than a PUT would be.
func PatchTest(c *gin.Context) {
	test, err := utils.DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if test.ID == 0 {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(
			errors.New("must define an id of an existing test to update")),
		)
		return
	}

	iTest := slices.IndexFunc(tests, func(t *models.Test) bool {
		return t.ID == test.ID
	})
	if iTest == -1 {
		err := fmt.Errorf("cannot update test, test with ID: %d does not exist", test.ID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if err := test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	// Perform partial update and doc merge
	existingTest := tests[iTest]
	if test.Summary != "" {
		existingTest.Summary = test.Summary
	}
	if test.Outcome != "" {
		existingTest.Outcome = test.Outcome
	}
	if test.Analysis != "" {
		existingTest.Analysis = test.Analysis
	}
	if test.Resolution != "" {
		existingTest.Resolution = test.Resolution
	}
	if test.Doc != nil && len(test.Doc) > 0 {
		for k, v := range test.Doc {
			existingTest.Doc[k] = v
		}
	}

	c.JSON(http.StatusOK, test)
}

func GetTests(c *gin.Context) {
	c.JSON(http.StatusOK, tests)
}
