package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var EnvConfig = GetConfig()

// GetConfig will return the Config from the environment or panic if something goes wrong.
func GetConfig() *Config {
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func GetRouter() *gin.Engine {
	pgPool, err := NewPGPool(EnvConfig.PG)
	if err != nil {
		log.Fatal(err)
	}
	testController := TestController{DBPool: pgPool}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/tests", testController.GetTests)
	r.DELETE("/tests", testController.DeleteTests)
	r.POST("/test", testController.CreateTest)
	r.PATCH("/test", testController.PatchTest)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "healthy"})
		return
	})

	r.POST("/query", EncodeSearchQuery)
	return r
}

func main() {
	r := GetRouter()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
