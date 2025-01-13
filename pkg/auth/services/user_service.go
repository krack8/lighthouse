package services

import (
	"context"
	"github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

var userCollection *mongo.Collection

func init() {
	client, _, _ := config.ConnectDB()
	userCollection = client.Database(os.Getenv("DB_NAME")).Collection("users")
}

// CreateUser inserts a new user into the MongoDB collection.
func CreateUser(user *models.User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}
