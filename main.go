package main

import (
	"github.com/gin-gonic/gin"
	"oar/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	r.GET("/tests", controllers.GetTests)
	r.POST("/test", controllers.CreateTest)

	r.GET("/analyses", controllers.GetAnalyses)
	r.PUT("/analysis", controllers.SetAnalysis)

	r.PUT("/resolution", controllers.SetResolution)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
