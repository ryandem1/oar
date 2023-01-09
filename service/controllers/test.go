package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/drivers"
	"github.com/ryandem1/oar/models"
	"github.com/ryandem1/oar/utils"
	"net/http"
)

// TestController will maintain a database pool for all test controllers
type TestController struct {
	DBPool *pgx.ConnPool
}

// GetTests will retrieve test objects from the database. Will take queries/limit/offset
func (tc *TestController) GetTests(c *gin.Context) {
	tests, err := drivers.SelectTests(tc.DBPool, "SELECT * FROM tests")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	c.JSON(http.StatusOK, tests)
}

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func (tc *TestController) CreateTest(c *gin.Context) {
	test, err := utils.DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	test.Clean()
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	err = drivers.InsertTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	c.Status(http.StatusCreated)
}

// PatchTest will perform a patch (partial update) operation on an existing test if it exists. Because of the nature of
// the Test enrichment process, I imagine this will be used more than a PUT would be.
func (tc *TestController) PatchTest(c *gin.Context) {
	testPatch, err := utils.DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if testPatch.ID == 0 {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(
			errors.New("must define an id of an existing testPatch to update")),
		)
		return
	}

	// Perform partial update on copy of existing testPatch and doc merge
	tests, err := drivers.SelectTests(tc.DBPool, "SELECT * FROM tests WHERE id=$1", testPatch.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	if len(tests) < 1 {
		c.JSON(
			http.StatusBadRequest,
			fmt.Errorf("cannot update, testPatch with ID: %d does not exist", testPatch.ID),
		)
		return
	} else if len(tests) > 1 {
		c.JSON(
			http.StatusInternalServerError,
			fmt.Errorf("found >1 testPatch with ID: %d, data is corrupted", testPatch.ID),
		)
		return
	}

	test := tests[0]
	if testPatch.Summary != "" {
		test.Summary = testPatch.Summary
	}
	if testPatch.Outcome != "" {
		test.Outcome = testPatch.Outcome
	}
	if testPatch.Analysis != "" {
		test.Analysis = testPatch.Analysis
	}
	if testPatch.Resolution != "" {
		test.Resolution = testPatch.Resolution
	}
	if testPatch.Doc != nil && len(testPatch.Doc) > 0 {
		for k, v := range testPatch.Doc {
			test.Doc[k] = v
		}
	}

	// Validate after update to ensure testPatch is still okay
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	// Update in DB
	err = drivers.UpdateTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	c.Status(http.StatusOK)
}

// DeleteTests is a bulk endpoint that can mark tests as deleted. These tests will no longer be visible from the UI, but
// will still remain in the DB unless manually removed just-in-case they need to be recovered.
// DeleteTests will silently ignore if the caller passes in test IDs that already don't exist.
// DeleteTests will respond with a http.StatusNotModified (304) status code if it does not delete a single test.
// DeleteTests will respond with a http.StatusOK (200) status code if it deletes at least 1 test.
func (tc *TestController) DeleteTests(c *gin.Context) {
	var testsToDelete []models.Test

	if err := c.BindJSON(&testsToDelete); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	var testIDsToDelete []int64
	for _, testToDelete := range testsToDelete {
		testIDsToDelete = append(testIDsToDelete, testToDelete.ID)
	}

	testsDeleted, err := drivers.DeleteTests(tc.DBPool, testIDsToDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if testsDeleted < 0 {
		c.Status(http.StatusInternalServerError)
	} else if testsDeleted == 0 {
		c.Status(http.StatusNotModified)
	} else {
		c.Status(http.StatusOK)
	}
}
