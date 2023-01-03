package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testable/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/test", controllers.CreateTest)
	r.GET("/tests", controllers.GetTests)
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
