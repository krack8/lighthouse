package services

import (
	"context"
	"github.com/go-errors/errors"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// CreateUser inserts a new user into the database.
func CreateUser(user *models.User) (*models.User, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return nil, err
	}
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

// GetUser retrieves a user by ID.
func GetUser(id string) (*models.User, error) {
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	filter := bson.M{"_id": objectID}
	err = db.UserCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers() ([]models.User, error) {
	var users []models.User

	// Define an empty filter to fetch all documents
	filter := bson.M{}

	// Use Find to retrieve multiple documents
	cursor, err := db.UserCollection.Find(context.Background(), filter)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into the `users` slice
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Printf("Error decoding user: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	// Check if the cursor encountered any errors during iteration
	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		return nil, err
	}

	return users, nil
}

// UpdateUser updates an existing user's information.
func UpdateUser(id string, updatedData *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	updatedData.UpdatedAt = time.Now()
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": updatedData}

	_, err = db.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

// DeleteUser deletes a user by ID.
func DeleteUser(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid user ID")
	}

	filter := bson.M{"_id": objectID}
	_, err = db.UserCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}

// GetUserByUsername fetches a user by username from the database
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
