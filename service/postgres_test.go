// I am not a fan of mocks, so these tests will need a local instance of postgres

package main

import (
	"fmt"
	"github.com/jackc/pgx"
	"testing"
)

// getPoolFromConfig will return a new pgx.ConnPool after getting a NewConfig from the environment
func getPoolFromConfig() (*pgx.ConnPool, error) {
	config, err := NewConfig() // Gets config from environment
	if err != nil {
		return nil, fmt.Errorf("error obtaining config. error: %s", err)
	}

	pgPool, err := NewPGPool(config.PG)
	if err != nil {
		return nil, fmt.Errorf("error obtaining pg connection. error: %s", err)
	}
	return pgPool, err
}

// TestNewPGPoolPositive ensures NewPGPool works with a valid config
func TestNewPGPoolPositive(t *testing.T) {
	config := &PGConfig{
		Host:        "localhost",
		Port:        5432,
		DB:          "oar",
		User:        "postgres",
		Pass:        "postgres",
		LogLevel:    pgx.LogLevelInfo,
		PoolSize:    4,
		PollTimeout: 60,
	}
	pool, err := NewPGPool(config)
	if err != nil {
		t.Error(err)
	}
	if pool == nil {
		t.Error("no error was thrown, but pool returned was nil")
	}
}

// TestNewPGPoolNegative ensures NewPGPool will throw an error and a nil pool for invalid configs
func TestNewPGPoolNegative(t *testing.T) {
	invalidConfigs := map[string]*PGConfig{
		"0 pool size": {
			Host:        "localhost",
			Port:        5432,
			DB:          "oar",
			User:        "postgres",
			Pass:        "postgres",
			LogLevel:    pgx.LogLevelInfo,
			PoolSize:    0,
			PollTimeout: 60,
		},
		"empty config": {
			Host:        "",
			Port:        0,
			DB:          "",
			User:        "",
			Pass:        "",
			LogLevel:    0,
			PoolSize:    0,
			PollTimeout: 0,
		},
	}
	for scenarioName, invalidConfig := range invalidConfigs {
		t.Run(scenarioName, func(t *testing.T) {
			pool, err := NewPGPool(invalidConfig)
			if err == nil {
				t.Error("no error was returned for an invalid config")
			}
			if pool != nil {
				t.Error("pool was not nil when an invalid config was passed")
			}
		})
	}
}

// TestSelectCreateTests will ensure that we can select valid tests that are in postgres
func TestSelectCreateTests(t *testing.T) {
	amountOfTests := 5 // number of tests to create/read

	pgPool, err := getPoolFromConfig()
	if err != nil {
		t.Error("error obtaining pg connection", err)
	}
	validTests := multiple(amountOfTests, Fake.test)

	for _, validTest := range validTests {
		err = InsertTest(pgPool, validTest)
		if err != nil {
			t.Error("error during data setup", err)
		}
	}

	selectedTests, err := SelectTests(
		pgPool,
		"select * from oar_tests order by created desc limit $1",
		amountOfTests,
	)

	if err != nil {
		t.Error(err)
	}

	for _, selectedTest := range selectedTests {
		err = selectedTest.Validate()
		if err != nil {
			t.Error(err)
		}
	}
}

// TestDeleteTests will ensure that we can delete tests with DeleteTests. Also ensures that if you attempt to delete
// tests that are already deleted, it will just return 0 rows affected with no errors.
func TestDeleteTests(t *testing.T) {
	amountOfTests := 5 // number of tests to create/read

	pgPool, err := getPoolFromConfig()
	if err != nil {
		t.Error("error obtaining pg connection", err)
	}

	validTests := multiple(amountOfTests, Fake.test)

	for _, validTest := range validTests {
		err = InsertTest(pgPool, validTest)
		if err != nil {
			t.Error("error during data setup", err)
		}
	}

	testsToDelete, err := SelectTests(
		pgPool,
		"select * from oar_tests order by created desc limit $1",
		amountOfTests,
	)

	var testIDsToDelete []uint64

	for _, testToDelete := range testsToDelete {
		testIDsToDelete = append(testIDsToDelete, testToDelete.ID)
	}

	rowsDeleted, err := DeleteTests(pgPool, testIDsToDelete)
	if err != nil {
		t.Error(err)
	}

	if rowsDeleted != int64(amountOfTests) {
		t.Error("not all tests were deleted")
	}

	rowsDeleted, err = DeleteTests(pgPool, testIDsToDelete)
	if err != nil {
		t.Error(err)
	}

	if rowsDeleted != 0 {
		t.Error("consecutive delete call deleted more rows")
	}
}
