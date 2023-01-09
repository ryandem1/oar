package drivers

import (
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"time"
)

// NewPGPool will establish a new connection with postgres and return a pointer to a  connection pool
func NewPGPool() (*pgx.ConnPool, error) {
	pgConnConfig := pgx.ConnConfig{
		Host:     viper.GetString("PGHost"),
		Port:     uint16(viper.GetInt("PGPort")),
		Database: viper.GetString("PGDatabase"),
		User:     viper.GetString("PGUser"),
		Password: viper.GetString("PGPass"),
		LogLevel: pgx.LogLevel(viper.GetInt("PGLogLevel")),
	}
	pgConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConnConfig,
		MaxConnections: viper.GetInt("PGPoolSize"),
		AfterConnect:   nil,
		AcquireTimeout: time.Duration(viper.GetInt("PGPoolTimeout")),
	}
	pgPool, err := pgx.NewConnPool(pgConnPoolConfig)
	if err != nil {
		return nil, err
	}

	return pgPool, nil
}
