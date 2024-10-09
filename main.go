package main

import (
	"context"
	"os"
	"yayawallet-webhook/controller"
	"yayawallet-webhook/repository"
	"yayawallet-webhook/services"
	"yayawallet-webhook/usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	mongoURI := os.Getenv("MONGO_URI")
	port := os.Getenv("PORT")
	// Set up repository
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	webhook_collection := client.Database("yayawallet").Collection("webhooks")
	repo := repository.NewWebhookRepository(webhook_collection, context.TODO())
	services := services.NewWebhookService(repo)
	uc := usecases.NewWebhookUseCase(services)
	handler := controller.NewWebhookHandler(uc)

	// Register the webhook endpoint
	r.POST("/webhook", handler.HandleWebhook)

	r.Run(":" + port)
}
