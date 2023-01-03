package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testable/helpers"
	"testable/models"
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

	tests = append(tests, test)
	c.JSON(http.StatusCreated, test)
}

func GetTests(c *gin.Context) {
	c.JSON(http.StatusOK, tests)
}
