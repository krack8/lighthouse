package services

import (
	"context"
	"github.com/go-errors/errors"
	db "github.com/krack8/lighthouse/pkg/auth/db"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

// CreateUser inserts a new user in the database.
func CreateUser(user *models.User) (*models.User, error) {

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := db.CreateUser(user)
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
