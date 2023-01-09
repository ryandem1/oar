package drivers

import "github.com/jackc/pgx"

// NewPGPool will establish a new connection with postgres and return a pointer to a  connection pool
func NewPGPool() (*pgx.ConnPool, error) {
	pgConnConfig := pgx.ConnConfig{
		Host:     "oar-postgres",
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
		return nil, err
	}

	return pgPool, nil
}
