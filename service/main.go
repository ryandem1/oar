package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var EnvConfig = GetConfig()

// GetConfig will return the Config from the environment or panic if something goes wrong
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

	r.DELETE("/tests", testController.DeleteTests)
	r.POST("/test", testController.CreateTest)
	r.PATCH("/test", testController.PatchTest)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"health": "healthy"})
		return
	})
	return r
}

func main() {
	r := GetRouter()
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
