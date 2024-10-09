package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"yayawallet-webhook/controller"
	"yayawallet-webhook/repository"
	"yayawallet-webhook/services"
	"yayawallet-webhook/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Import godotenv
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	mongoURI := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")

	// Ensure environment variables are loaded
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	// Set up Gin and MongoDB connection
	r := gin.Default()
	fmt.Println("Connecting to MongoDB at", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	webhook_collection := client.Database("yayawallet").Collection("payloads")
	repo := repository.NewWebhookRepository(webhook_collection, context.TODO())
	services := services.NewWebhookService(repo)
	uc := usecases.NewWebhookUseCase(services)
	handler := controller.NewWebhookHandler(uc)

	// Register the webhook endpoint
	r.POST("/webhook", handler.HandleWebhook)

	r.Run(":" + port)
}
