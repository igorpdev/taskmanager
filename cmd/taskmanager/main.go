package main

import (
	"fmt"
	"log"
	"os"
	"taskmanager/internal/controller"
	"taskmanager/internal/database"
	"taskmanager/internal/router"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		URI         string        `yaml:"uri"`
		Name        string        `yaml:"name"`
		MinPoolSize uint64        `yaml:"minPoolSize"`
		MaxPoolSize uint64        `yaml:"maxPoolSize"`
		Timeout     time.Duration `yaml:"timeout"`
	} `yaml:"database"`
	Metrics struct {
		Enabled        bool   `yaml:"enabled"`
		Endpoint       string `yaml:"endpoint"`
		ScrapeInterval string `yaml:"scrapeInterval"`
	} `yaml:"metrics"`
}

func loadConfig() (*Config, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := database.Connect(cfg.Database.URI, cfg.Database.MinPoolSize, cfg.Database.MaxPoolSize, cfg.Database.Timeout); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func() {
		if err := database.Disconnect(); err != nil {
			log.Fatalf("Failed to disconnect from database: %v", err)
		}
	}()

	controller.InitTaskCollection()

	r := gin.Default()

	if cfg.Metrics.Enabled {
		log.Printf("Metrics enabled at %s", cfg.Metrics.Endpoint)
		router.SetupRoutes(r, cfg.Metrics.Endpoint)
	} else {
		router.SetupRoutes(r, "")
	}

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
