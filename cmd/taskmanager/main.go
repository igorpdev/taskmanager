package main

import (
	"fmt"
	"log"
	"os"
	"taskmanager/internal/controller"
	"taskmanager/internal/database"
	"taskmanager/internal/router"

	"gopkg.in/yaml.v3"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		URI  string `yaml:"uri"`
		Name string `yaml:"name"`
	} `yaml:"database"`
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

	if err := database.Connect(cfg.Database.URI); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	controller.InitTaskCollection()

	r := gin.Default()
	router.SetupRoutes(r)

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
