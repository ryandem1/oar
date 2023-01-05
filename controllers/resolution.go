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
	tr := &models.Test{}

	if err := c.BindJSON(tr); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	iTest := slices.IndexFunc(tests, func(test *models.Test) bool {
		return test.ID == tr.ID
	})
	if iTest == -1 {
		err := fmt.Errorf("cannot set analysis, test with ID: %d does not exist", tr.ID)
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}
	test := tests[iTest]
	test.Resolution = tr.Resolution

	if err := tr.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, utils.ConvertErrToGinH(err))
		return
	}

	c.JSON(http.StatusAccepted, tr)
}
