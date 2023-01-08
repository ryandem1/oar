package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health will display the health status of the application
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "healthy"})
	return
}
