package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	URI         string        `mapstructure:"uri"`
	Name        string        `mapstructure:"name"`
	MinPoolSize int           `mapstructure:"minPoolSize"`
	MaxPoolSize int           `mapstructure:"maxPoolSize"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

type MetricsConfig struct {
	Enabled        bool          `mapstructure:"enabled"`
	Endpoint       string        `mapstructure:"endpoint"`
	ScrapeInterval time.Duration `mapstructure:"scrapeInterval"`
}

type CircuitBreakerConfig struct {
	MaxRequests int           `mapstructure:"maxRequests"`
	Interval    time.Duration `mapstructure:"interval"`
	Timeout     time.Duration `mapstructure:"timeout"`
}

type Config struct {
	Server         struct{ Port int }   `mapstructure:"server"`
	Database       DatabaseConfig       `mapstructure:"database"`
	Metrics        MetricsConfig        `mapstructure:"metrics"`
	CircuitBreaker CircuitBreakerConfig `mapstructure:"circuitBreaker"`
}

var AppConfig Config

func LoadConfig(configPath string) {
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading configuration file: %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshalling configuration: %s", err)
	}
}
