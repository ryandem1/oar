package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/controllers"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	pgConnConfig := pgx.ConnConfig{
		Host:     "192.168.0.2",
		Port:     5432,
		Database: "oar",
		User:     "postgres",
		Password: "postgres",
		LogLevel: pgx.LogLevelInfo,
	}
	pgConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConnConfig,
		MaxConnections: 4,
		AfterConnect:   nil,
		AcquireTimeout: 30,
	}
	pgPool, err := pgx.NewConnPool(pgConnPoolConfig)
	if err != nil {
		log.Fatal(err)
	}
	testController := controllers.TestController{DBPool: pgPool}
	r.GET("/tests", testController.GetTests)
	r.DELETE("/tests", testController.DeleteTests)

	r.POST("/test", testController.CreateTest)
	r.PATCH("/test", testController.PatchTest)

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
