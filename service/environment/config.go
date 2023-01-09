package environment

import (
	"github.com/jackc/pgx"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type PGConfig struct {
	Host        string        `mapstructure:"PG_HOST"`
	Port        uint16        `mapstructure:"PG_PORT"`
	DB          string        `mapstructure:"PG_DB"`
	User        string        `mapstructure:"PG_USER"`
	Pass        string        `mapstructure:"PG_PASS"`
	LogLevel    pgx.LogLevel  `mapstructure:"PG_LL"`
	PoolSize    int           `mapstructure:"PG_POOL_SIZE"`
	PollTimeout time.Duration `mapstructure:"PG_POOL_TIMEOUT"`
}

type Config struct {
	PG *PGConfig
}

func NewConfig() (*Config, error) {
	SetDefaultConfigValues()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	config := &Config{}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("could not find a config file! File must be named 'config.<yaml or toml or json>'")
		} else {
			log.Fatalf("fatal error config file: %s", err.Error())
		}
	}

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SetDefaultConfigValues will set all the default configuration variables with viper
func SetDefaultConfigValues() {
	viper.SetDefault("PG_HOST", "oar-postgres")
	viper.SetDefault("PG_PORT", 5432)
	viper.SetDefault("PG_DB", "oar")
	viper.SetDefault("PG_USER", "postgres")
	viper.SetDefault("PG_PASS", "postgres")
	viper.SetDefault("PG_LL", pgx.LogLevelInfo) // Postgres Log Level
	viper.SetDefault("PG_POOL_SIZE", 4)         // Max number of pool connections
	viper.SetDefault("PG_POLL_TIMEOUT", 30)     // Time to wait for a connection to be freed up
}
