package main

import (
	"log"
	"os"
	"backend/config"
	"backend/services"
	"github.com/gin-gonic/gin"
	"backend/api"
)

func main() {
	cfg := config.LoadConfig()
	mongoClient := services.InitMongo(cfg)
	defer mongoClient.Disconnect(nil)
	vectorClient := services.InitMilvus(cfg)
	defer vectorClient.Close()

	r := gin.Default()
	api.RegisterRoutes(r, mongoClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":"+port); err != nil {
		log.Fatal(err)
	}
}