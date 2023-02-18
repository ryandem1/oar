package main

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
	"time"
)

type PGConfig struct {
	Host        string        `mapstructure:"HOST"`
	Port        uint16        `mapstructure:"PORT"`
	DB          string        `mapstructure:"DB"`
	User        string        `mapstructure:"USER"`
	Pass        string        `mapstructure:"PASS"`
	LogLevel    pgx.LogLevel  `mapstructure:"LL"`
	PoolSize    int           `mapstructure:"POOL_SIZE"`
	PollTimeout time.Duration `mapstructure:"POOL_TIMEOUT"`
}

// NewPGPool will establish a new connection with postgres and return a pointer to a  connection pool
func NewPGPool(config *PGConfig) (*pgx.ConnPool, error) {
	pgConnConfig := pgx.ConnConfig{
		Host:     config.Host,
		Port:     config.Port,
		Database: config.DB,
		User:     config.User,
		Password: config.Pass,
		LogLevel: config.LogLevel,
	}
	pgConnPoolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgConnConfig,
		MaxConnections: config.PoolSize,
		AfterConnect:   nil,
		AcquireTimeout: config.PollTimeout,
	}
	pgPool, err := pgx.NewConnPool(pgConnPoolConfig)
	if err != nil {
		return nil, err
	}

	return pgPool, nil
}

// InsertTest will insert a new models.Test object into the postgres DB
func InsertTest(pgPool *pgx.ConnPool, test *Test) error {
	conn, err := pgPool.Acquire()
	if err != nil {
		return err
	}
	defer pgPool.Release(conn)

	exec, err := conn.Exec(
		"INSERT INTO OAR_TESTS (summary, outcome, analysis, resolution, doc) VALUES ($1, $2, $3, $4, $5)",
		test.Summary,
		test.Outcome,
		test.Analysis,
		test.Resolution,
		test.Doc,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return fmt.Errorf("rows affected: %d != 1", exec.RowsAffected())
	}

	return nil
}

// UpdateTest will update an existing test in the postgres DB by ID
func UpdateTest(pgPool *pgx.ConnPool, test *Test) error {
	conn, err := pgPool.Acquire()
	if err != nil {
		return err
	}
	defer pgPool.Release(conn)

	if err != nil {
		return err
	}
	exec, err := conn.Exec(
		"UPDATE OAR_TESTS SET summary=$1, outcome=$2, analysis=$3, resolution=$4, doc=$5 WHERE id=$6",
		test.Summary,
		test.Outcome,
		test.Analysis,
		test.Resolution,
		test.Doc,
		test.ID,
	)
	if err != nil {
		return err
	}
	if exec.RowsAffected() != 1 {
		return fmt.Errorf("rows affected: %d != 1", exec.RowsAffected())
	}

	return nil
}

// SelectTests will take in a query that returns rows that are in the models.Test schema, deserialize them, and return
// pointers to the models.
// args will be passed down to Conn.query
func SelectTests(pgPool *pgx.ConnPool, query string, args ...any) ([]*Test, error) {
	conn, err := pgPool.Acquire()
	if err != nil {
		return nil, err
	}
	defer pgPool.Release(conn)
	var tests []*Test

	rows, err := conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		test := &Test{}
		err = rows.Scan(
			&test.ID,
			&test.Summary,
			&test.Outcome,
			&test.Analysis,
			&test.Resolution,
			&test.Created,
			&test.Modified,
			&test.Doc,
		)
		if err != nil {
			return nil, err
		}
		tests = append(tests, test)
	}

	return tests, nil
}

// DeleteTests will take in a slice of test IDs and attempt to delete all tests with those IDs. Will return the amount
// of rows deleted and any error that occurred. If an error occurred, it will return -1 rows deleted, which is invalid.
func DeleteTests(pgPool *pgx.ConnPool, testIDs []uint64) (int64, error) {
	conn, err := pgPool.Acquire()
	if err != nil {
		return -1, err
	}
	defer pgPool.Release(conn)

	pgTestIDs := &pgtype.Int8Array{}
	err = pgTestIDs.Set(testIDs)
	if err != nil {
		return -1, err
	}

	exec, err := conn.Exec("DELETE FROM OAR_TESTS WHERE ID = ANY($1)", pgTestIDs)
	if err != nil {
		return -1, err
	}

	return exec.RowsAffected(), nil
}
