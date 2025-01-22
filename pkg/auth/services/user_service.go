package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/config"
	"time"

	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection defines an interface for MongoDB operations
type Collection interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
}

// UserService struct for user operations
type UserService struct {
	UserCollection Collection
	Context        context.Context
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := config.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUser(userID string) (*models.User, error) {
	// Validate dependencies
	if s.UserCollection == nil {
		return nil, errors.New("user collection not initialized")
	}
	if s.Context == nil {
		return nil, errors.New("context not initialized")
	}

	// Validate input
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	// Create filter
	filter := bson.M{"_id": objectID}

	// Find the user
	var user models.User
	result := config.UserCollection.FindOne(context.Background(), filter)
	if err := result.Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}

// GetAllUsers retrieves all users
func (s *UserService) GetAllUsers() ([]models.User, error) {
	cursor, err := config.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.Context)

	var users []models.User
	for cursor.Next(s.Context) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser updates a user by their ID
func (s *UserService) UpdateUser(userID string, updatedUser *models.User) error {
	if updatedUser == nil {
		return errors.New("updated user cannot be nil")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"firstname": updatedUser.FirstName,
			"lastname":  updatedUser.LastName,
			"updatedat": time.Now(),
		},
	}

	result, err := config.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// DeleteUser deletes a user by their ID
func (s *UserService) DeleteUser(userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID")
	}

	filter := bson.M{"_id": objectID}
	result, err := config.UserCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	filter := bson.M{"username": username}
	// Check if filter is not nil
	if filter == nil {
		// Handle the error
		return nil, errors.New("Filter is nil")
	}

	// FindOne with error handling
	result := config.UserCollection.FindOne(context.Background(), filter)

	var user models.User
	if err := result.Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
