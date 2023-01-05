package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"oar/models"
	"oar/utils"
)

// SetAnalysis will set the analysis of a test. A test can only have 1 analysis at a time.
func SetAnalysis(c *gin.Context) {
	ta := &models.TestAnalysis{}

	if err := c.BindJSON(ta); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	iTest := slices.IndexFunc(tests, func(test *models.Test) bool {
		return test.ID == ta.TestID
	})
	if iTest == -1 {
		err := fmt.Errorf("cannot set analysis, test with ID: %d does not exist", ta.TestID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	test := tests[iTest]

	if err := ta.Validate(test); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	analyses[ta.TestID] = ta
	c.JSON(http.StatusAccepted, ta)
}

// GetAnalyses will return all analyses
func GetAnalyses(c *gin.Context) {
	c.JSON(http.StatusOK, analyses)
	return
}
