package services

import (
	"context"
	"errors"
	"fmt"
	db "github.com/krack8/lighthouse/pkg/auth/config"
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

// UserService handles user-related business logic
type UserService struct {
	collection Collection
}

// NewUserService creates a new UserService instance
func NewUserService(collection Collection) *UserService {
	return &UserService{
		collection: collection,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	data, _ := GetUserByUsername(user.Username)

	if data != nil {
		return nil, errors.New("user already exists")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(userID string) (*models.User, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format: %w", err)
	}

	var user models.User
	filter := bson.M{"_id": objectID}
	result := db.UserCollection.FindOne(context.Background(), filter)
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
	cursor, err := db.UserCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return users, nil
}

// UpdateUser updates a user by ID
func (s *UserService) UpdateUser(userID string, updatedUser *models.User) error {
	if updatedUser == nil {
		return errors.New("updated user cannot be nil")
	}

	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"firstname":  updatedUser.FirstName,
			"lastname":   updatedUser.LastName,
			"username":   updatedUser.Username,
			"password":   updatedUser.Password,
			"usertype":   updatedUser.UserType,
			"roles":      updatedUser.Roles,
			"isactive":   updatedUser.UserIsActive,
			"isverified": updatedUser.IsVerified,
			"phone":      updatedUser.Phone,
			"updatedat":  time.Now(),
		},
	}

	result, err := db.UserCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(userID string) error {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return fmt.Errorf("invalid user ID format: %w", err)
	}

	filter := bson.M{"_id": objectID}
	result, err := db.UserCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// GetUserByUsername retrieves a user by username
func GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	var user models.User
	filter := bson.M{"username": username}
	if err := db.UserCollection.FindOne(context.Background(), filter).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}
