package main

import (
	"github.com/gin-gonic/gin"
	"oar/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	r.GET("/tests", controllers.GetTests)
	r.DELETE("/tests", controllers.DeleteTests)

	r.POST("/test", controllers.CreateTest)
	r.PATCH("/test", controllers.PatchTest)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
