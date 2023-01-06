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

// SetAnalysis will set the analysis of a test. A test can only have 1 analysis at a time.
func SetAnalysis(c *gin.Context) {
	ta := &models.Test{}

	if err := c.BindJSON(ta); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if ta.ID == 0 {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(
			errors.New("must define an id to set the analysis for")),
		)
		return
	}

	iTest := slices.IndexFunc(tests, func(test *models.Test) bool {
		return test.ID == ta.ID
	})
	if iTest == -1 {
		err := fmt.Errorf("cannot set analysis, test with ID: %d does not exist", ta.ID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	test := *tests[iTest]
	test.Analysis = ta.Analysis

	if err := test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	tests[iTest] = &test
	c.JSON(http.StatusAccepted, test)
}
