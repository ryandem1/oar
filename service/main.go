package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/controllers"
	"github.com/ryandem1/oar/drivers"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("could not find a config file! File must be named 'config.<yaml or toml or json>'")
		} else {
			log.Fatalf("fatal error config file: %s", err.Error())
		}
	}
	viper.SetDefault("PGHost", "oar-postgres")
	viper.SetDefault("PGPort", 5432)
	viper.SetDefault("PGDatabase", "oar")
	viper.SetDefault("PGUser", "postgres")
	viper.SetDefault("PGPass", "postgres")
	viper.SetDefault("PGLogLevel", pgx.LogLevelInfo)
	viper.SetDefault("PGPoolSize", 4)     // Max number of pool connections
	viper.SetDefault("PGPoolTimeout", 30) // Time to wait for a connection to be freed up

	pgPool, err := drivers.NewPGPool()
	if err != nil {
		log.Fatal(err)
	}
	testController := controllers.TestController{DBPool: pgPool}

	r := gin.Default()
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
