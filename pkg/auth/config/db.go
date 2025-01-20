package config

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB collections
var (
	PermissionCollection *mongo.Collection
	RoleCollection       *mongo.Collection
	UserCollection       *mongo.Collection
)

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
	PermissionCollection = client.Database(dbName).Collection("permissions")
	RoleCollection = client.Database(dbName).Collection("roles")

	log.Println("Successfully connected to MongoDB")
	return client, ctx, nil
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
			Password:     utils.HashPassword("admin123"), // Use a hashed password here
			UserType:     "ADMIN",
			Roles:        []models.Role{},
			UserIsActive: true,
			IsVerified:   true,
			Phone:        "1234567890",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		_, err := UserCollection.InsertOne(context.Background(), defaultUser)
		if err != nil {
			log.Fatalf("Error creating default user: %v", err)
		}

		log.Println("Default user created successfully.")
	} else {
		log.Println("Users already exist. No default user created.")
	}
}

func InitRBAC() {
	permissionCount, err := PermissionCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Error counting documents in users collection: %v", err)
	}

	roleCount, err := RoleCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("Error counting documents in users collection: %v", err)
	}

	var defaultPermissions []models.Permission

	if permissionCount == 0 {
		// Example permissions
		defaultPermissions = []models.Permission{
			{
				ID:          primitive.NewObjectID(),
				Name:        "view_dashboard",
				Description: "Permission to view the dashboard",
				Route:       "/dashboard",
				Method:      "GET",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			{
				ID:          primitive.NewObjectID(),
				Name:        "edit_profile",
				Description: "Permission to edit user profile",
				Route:       "/profile/edit",
				Method:      "POST",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
		}

		// Convert the []models.Permission to []interface{}
		var permissionsInterface []interface{}
		for _, perm := range defaultPermissions {
			permissionsInterface = append(permissionsInterface, perm)
		}

		// Insert the Default Permissions into the collection
		_, err := PermissionCollection.InsertMany(context.Background(), permissionsInterface)
		if err != nil {
			log.Printf("Error inserting permissions: %v", err)
		}
	}

	if roleCount == 0 {
		// Example role with permissions
		defaultRole := models.Role{
			ID:          primitive.NewObjectID(),
			Name:        "Admin",
			Description: "Administrator role with all permissions",
			Permissions: defaultPermissions,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		// Insert the Default Role into the collection
		_, err = RoleCollection.InsertOne(context.Background(), defaultRole)
		if err != nil {
			log.Printf("Error inserting role: %v", err)
		}
	}
}
