package db

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB collections
var UserCollection *mongo.Collection
var CounterCollection *mongo.Collection

// ConnectDB initializes the MongoDB client and collections.
func ConnectDB() (*mongo.Client, context.Context, error) {
	ctx := context.Background()

	// Get the MongoDB URI and database name from environment variables
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable is not set")
		return nil, nil, fmt.Errorf("MONGO_URI not set")
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, nil, err
	}

	// Get the DB name from environment variables
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB_NAME environment variable is not set")
		return nil, nil, fmt.Errorf("DB_NAME not set")
	}

	// Initialize the collections
	UserCollection = client.Database(dbName).Collection("users")

	log.Println("Successfully connected to MongoDB")
	return client, ctx, nil
}

// CreateUser inserts a new user into the database.
func CreateUser(user *models.User) error {
	// Insert the user into the collection
	_, err := UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}
	return nil
}

// InitializeDefaultUser creates a default user if no users exist.
func InitializeDefaultUser() {
	count, err := UserCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Error counting documents in users collection: %v", err)
	}

	if count == 0 {
		defaultUser := models.User{
			Username:     "admin@default.com",
			FirstName:    "Admin",
			LastName:     "User",
			Password:     "admin123", // Use a hashed password here
			UserType:     "ADMIN",
			UserIsActive: true,
			IsVerified:   true,
			Phone:        "1234567890",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		err := CreateUser(&defaultUser)
		if err != nil {
			log.Fatalf("Error creating default user: %v", err)
		}

		log.Println("Default user created successfully.")
	} else {
		log.Println("Users already exist. No default user created.")
	}
}
