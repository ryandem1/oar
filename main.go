package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"log"
	"oar/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/health", controllers.Health)

	pgConnConfig := pgx.ConnConfig{
		Host:                 "",
		Port:                 0,
		Database:             "",
		User:                 "",
		Password:             "",
		TLSConfig:            nil,
		UseFallbackTLS:       false,
		FallbackTLSConfig:    nil,
		Logger:               nil,
		LogLevel:             0,
		Dial:                 nil,
		RuntimeParams:        nil,
		OnNotice:             nil,
		CustomConnInfo:       nil,
		CustomCancel:         nil,
		PreferSimpleProtocol: false,
		TargetSessionAttrs:   "",
	}
	pgConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConnConfig,
		MaxConnections: 0,
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
