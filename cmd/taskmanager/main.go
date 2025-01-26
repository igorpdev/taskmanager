package main

import (
	"fmt"
	"log"
	"taskmanager/internal/config"
	"taskmanager/internal/controller"
	"taskmanager/internal/database"
	"taskmanager/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig("configs/config.yaml")

	dbConfig := config.AppConfig.Database
	err := database.Connect(dbConfig.URI, uint64(dbConfig.MinPoolSize), uint64(dbConfig.MaxPoolSize), dbConfig.Timeout)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func() {
		if err := database.Disconnect(); err != nil {
			log.Fatalf("Error disconnecting from the database: %v", err)
		}
	}()

	controller.InitTaskCollection()

	r := gin.Default()
	router.SetupRoutes(r)

	port := config.AppConfig.Server.Port
	log.Printf("Server is running on port %d", port)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
