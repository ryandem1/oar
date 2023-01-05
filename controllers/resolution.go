package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
	"net/http"
	"oar/models"
	"oar/utils"
)

// SetResolution will set a Test's resolution
func SetResolution(c *gin.Context) {
	tr := &models.TestResolution{}

	if err := c.BindJSON(tr); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	iTest := slices.IndexFunc(tests, func(test *models.Test) bool {
		return test.ID == tr.TestID
	})
	if iTest == -1 {
		err := fmt.Errorf("cannot set analysis, test with ID: %d does not exist", tr.TestID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	if err := tr.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	resolutions[tr.TestID] = tr
	c.JSON(http.StatusAccepted, tr)
}

// GetResolutions will get all resolutions
func GetResolutions(c *gin.Context) {
	c.JSON(http.StatusOK, resolutions)
	return
}
