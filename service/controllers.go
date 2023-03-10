package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"net/http"
)

// TestController will maintain a database pool for all test controllers
type TestController struct {
	DBPool *pgx.ConnPool
}

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func (tc *TestController) CreateTest(c *gin.Context) {
	test, err := DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	test.Clean()
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	testID, err := InsertTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	c.JSON(http.StatusCreated, testID)
}

// PatchTest will perform a patch (partial update) operation on an existing test if it exists. Because of the nature of
// the Test enrichment process, I imagine this will be used more than a PUT would be.
func (tc *TestController) PatchTest(c *gin.Context) {
	testPatch, err := DoubleBindTest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	if testPatch.ID == 0 {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(
			errors.New("must define an id of an existing testPatch to update")),
		)
		return
	}

	// Perform partial update on copy of existing testPatch and doc merge
	tests, err := SelectTests(tc.DBPool, "SELECT * FROM OAR_TESTS WHERE id=$1", testPatch.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	test := tests[0]
	test.Merge(testPatch)

	// Validate after update to ensure testPatch is still okay
	if err = test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}

	// Update in DB
	err = UpdateTest(tc.DBPool, test)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
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
	var testsToDelete []Test

	if err := c.BindJSON(&testsToDelete); err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
		return
	}
	var testIDsToDelete []uint64
	for _, testToDelete := range testsToDelete {
		testIDsToDelete = append(testIDsToDelete, testToDelete.ID)
	}

	testsDeleted, err := DeleteTests(tc.DBPool, testIDsToDelete)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConvertErrToGinH(err))
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
