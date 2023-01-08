package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/models"
	"github.com/ryandem1/oar/utils"
	"golang.org/x/exp/slices"
	"net/http"
)

var tests []*models.Test // Temp store, will implement DB later

// TestController will maintain a database pool for all test controllers
type TestController struct {
	DBPool *pgx.ConnPool
}

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func (tc *TestController) CreateTest(c *gin.Context) {
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
func (tc *TestController) PatchTest(c *gin.Context) {
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

	// Perform partial update on copy of existing test and doc merge
	existingTest := *tests[iTest]
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

	// Validate after update to ensure test is still okay
	if err := existingTest.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	tests[iTest] = &existingTest
	c.JSON(http.StatusOK, existingTest)
}

// DeleteTests is a bulk endpoint that can mark tests as deleted. These tests will no longer be visible from the UI, but
// will still remain in the DB unless manually removed just-in-case they need to be recovered.
// DeleteTests will silently ignore if the caller passes in test IDs that already don't exist.
// DeleteTests will respond with a http.StatusNotModified (304) status code if it does not delete a single test.
// DeleteTests will respond with a http.StatusOK (200) status code if it deletes at least 1 test.
func (tc *TestController) DeleteTests(c *gin.Context) {
	var testsToDelete *[]models.Test

	if err := c.BindJSON(testsToDelete); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	statusCode := http.StatusNotModified

	for _, test := range *testsToDelete {
		iTest := slices.IndexFunc(tests, func(t *models.Test) bool {
			return t.ID == test.ID
		})
		// Silently ignore tests that do not exist
		if iTest == -1 {
			continue
		}

		statusCode = http.StatusOK // If we have deleted at least 1 item, we use status code 200 instead of 304
		tests = slices.Delete(tests, iTest, iTest+1)
	}

	c.Status(statusCode)
}

func (tc *TestController) GetTests(c *gin.Context) {
	c.JSON(http.StatusOK, tests)
}
