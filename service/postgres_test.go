package main

import (
	"github.com/jackc/pgx"
	"testing"
)

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
