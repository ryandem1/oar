package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ryandem1/oar/controllers"
	"github.com/ryandem1/oar/drivers"
	"github.com/ryandem1/oar/models"
	"log"
	"time"
)

func main() {
	config, err := models.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgPool, err := drivers.NewPGPool(config.PG)
	if err != nil {
		log.Fatal(err)
	}
	testController := controllers.TestController{DBPool: pgPool}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", controllers.Health)
	r.GET("/tests", testController.GetTests)
	r.DELETE("/tests", testController.DeleteTests)
	r.POST("/test", testController.CreateTest)
	r.PATCH("/test", testController.PatchTest)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
