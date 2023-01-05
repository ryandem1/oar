package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"oar/models"
	"oar/utils"
	"strings"
)

var tests []*models.Test                     // Temp store, will implement DB later
var analyses = make(map[int]models.Analysis) // Temp store, will implement DB later

// CreateTest will create a new test from a Summary, Outcome, and optional Doc
func CreateTest(c *gin.Context) {
	test := &models.Test{}

	// We must copy our request body for the second unmarshal because the bind operation will consume it
	byteBody, err := utils.CopyRequestBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	// First bind binds the test information
	if err := c.BindJSON(test); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	// Check if Doc is manually defined, it should not be. If it is, it causes all sorts of conflicts
	if test.Doc != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(
			fmt.Errorf("'Doc' is reserved! Cannot use that key")),
		)
		return
	}

	// Second bind will move all dynamic fields to test.Doc
	if err := json.Unmarshal(byteBody, &test.Doc); err != nil {
		panic(err)
	}
	// Removes the keys that are from the first binding
	for key := range test.Doc {
		if slices.Contains([]string{"summary", "id", "outcome"}, strings.ToLower(key)) {
			delete(test.Doc, key)
		}
	}

	if err := test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	test.ID = len(tests) + 1
	test.Clean()
	tests = append(tests, test)
	c.JSON(http.StatusCreated, test)
}

func GetTests(c *gin.Context) {
	c.JSON(http.StatusOK, tests)
}

// SetAnalysis will set the analysis of a test. A test can only have 1 analysis at a time.
func SetAnalysis(c *gin.Context) {
	ta := &models.TestAnalysis{}

	if err := c.BindJSON(ta); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	testExists := slices.ContainsFunc(tests, func(existingTest *models.Test) bool {
		return ta.TestID == existingTest.ID
	})

	if !testExists {
		err := fmt.Errorf("cannot set analysis, test with ID: %d does not exist", ta.TestID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if err := ta.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	analyses[ta.TestID] = ta.Analysis
	c.JSON(http.StatusAccepted, ta)
}
