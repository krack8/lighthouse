package services

import (
	"context"
	"errors"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"
)

func Login(db *mongo.Database, username, password string) (string, string, error) {
	collection := db.Collection("users")

	var user models.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", "", errors.New("invalid credentials")
	}

	// Collect all permissions of the user's roles
	var permissions []string
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			permissions = append(permissions, perm.Route+":"+perm.Method)
		}
	}

	// Load expiry durations from environment variables
	accessTokenExpiry, err := parseDurationFromEnv("ACCESS_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	refreshTokenExpiry, err := parseDurationFromEnv("REFRESH_TOKEN_EXPIRY")
	if err != nil {
		return "", "", err
	}

	// Generate JWT tokens with the username and permissions
	accessToken, err := utils.GenerateToken(username, permissions, os.Getenv("JWT_SECRET"), accessTokenExpiry)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(username, permissions, os.Getenv("JWT_REFRESH_SECRET"), refreshTokenExpiry)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Helper function to parse durations from environment variables
func parseDurationFromEnv(envKey string) (time.Duration, error) {
	value := os.Getenv(envKey)
	if value == "" {
		return 0, errors.New(envKey + " is not set in environment variables")
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return 0, errors.New("invalid duration format for " + envKey + ": " + value)
	}

	return parsed, nil
}
