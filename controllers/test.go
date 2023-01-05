package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oar/helpers"
	"oar/models"
)

var tests []*models.Test

func CreateTest(c *gin.Context) {
	test := &models.Test{}

	if err := c.BindJSON(test); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ConvertErrToGinH(err))
		return
	}
	if err := test.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ConvertErrToGinH(err))
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
