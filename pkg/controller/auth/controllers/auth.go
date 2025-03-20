package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/auth/config"
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
	"github.com/krack8/lighthouse/pkg/controller/auth/services"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	api2 "github.com/krack8/lighthouse/pkg/controller/rest/api"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"os"
	"strings"
	"time"
)

// LoginHandler handles user login requests
func LoginHandler(c *gin.Context) {
	var requestBody map[string]string

	// Bind the request body to the map
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	username := requestBody["username"]
	password := requestBody["password"]

	// Ensure both username and password are provided
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// Call the login service to get the tokens
	accessToken, refreshToken, err := services.Login(username, password)
	if err != nil {
		// Return structured error in JSON format
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the access and refresh tokens as JSON
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshTokenHandler handles token refresh requests
func RefreshTokenHandler(c *gin.Context) {
	var request map[string]string

	// Bind the request body to the map
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	refreshToken := request["refresh_token"]
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing refresh token"})
		return
	}

	// Validate the refresh token
	claims, err := utils.ValidateToken(refreshToken, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	// Load expiry durations from environment variables
	accessTokenExpiry, err := parseDurationFromEnv("ACCESS_TOKEN_EXPIRY")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load access token expiry"})
		return
	}

	// Generate a new access token
	accessToken, err := utils.GenerateToken(claims.Username, os.Getenv("JWT_SECRET"), accessTokenExpiry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	// Send response with the new access token
	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
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

func LogoutHandler(c *gin.Context) {
	var token string

	token, exists := c.GetQuery("token")
	if exists == false {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		} else {
			// Remove "Bearer " prefix from the token if present
			token = strings.TrimPrefix(authHeader, "Bearer ")
		}
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
		c.Abort()
		return
	}

	// Validate token and extract claims
	claims, err := utils.ValidateToken(token, os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token"})
		c.Abort()
		return
	}

	filter := bson.M{"username": claims.Username}
	// Check if filter is not nil
	if filter == nil {
		// Handle the error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	// FindOne with error handling
	result := config.UserCollection.FindOne(context.Background(), filter)

	var user models.User
	if err := result.Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		c.Abort()
		return
	}

	go api2.PodController().ClearAllPodExecConnection(user.ID.Hex())

	// Send response with the success message
	c.JSON(http.StatusOK, gin.H{})
}
