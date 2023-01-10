package models

import (
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"log"
	"strings"
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

type Config struct {
	PG *PGConfig
}

func NewConfig() (*Config, error) {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	SetDefaultConfigValues()

	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("no config file loaded; if using, file must be named 'config.<yaml or toml or json>'")
		} else {
			log.Fatalf("fatal error config file: %s", err.Error())
		}
	}

	viper.AutomaticEnv()
	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SetDefaultConfigValues will set all the default configuration variables with viper
func SetDefaultConfigValues() {
	viper.SetDefault("PG.HOST", "oar-postgres")
	viper.SetDefault("PG.PORT", 5432)
	viper.SetDefault("PG.DB", "oar")
	viper.SetDefault("PG.USER", "postgres")
	viper.SetDefault("PG.PASS", "postgres")
	viper.SetDefault("PG.LL", pgx.LogLevelInfo) // Postgres Log Level
	viper.SetDefault("PG.POOL_SIZE", 4)         // Max number of pool connections
	viper.SetDefault("PG.POLL_TIMEOUT", 30)     // Time to wait for a connection to be freed up
}
