package drivers

import (
	"github.com/jackc/pgx"
	"github.com/ryandem1/oar/environment"
	"time"
)

// NewPGPool will establish a new connection with postgres and return a pointer to a  connection pool
func NewPGPool(config *environment.PGConfig) (*pgx.ConnPool, error) {
	pgConnConfig := pgx.ConnConfig{
		Host:     config.Host,
		Port:     config.Port,
		Database: config.DB,
		User:     config.User,
		Password: config.Pass,
		LogLevel: pgx.LogLevel(config.LogLevel),
	}
	pgConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConnConfig,
		MaxConnections: config.PoolSize,
		AfterConnect:   nil,
		AcquireTimeout: time.Duration(config.PollTimeout),
	}
	pgPool, err := pgx.NewConnPool(pgConnPoolConfig)
	if err != nil {
		return nil, err
	}

	return pgPool, nil
}
