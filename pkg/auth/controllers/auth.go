package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"os"
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
