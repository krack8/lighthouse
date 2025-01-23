package controllers

import (
	"encoding/json"
	"errors"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"os"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username := requestBody["username"]
	password := requestBody["password"]
	// Ensure both username and password are provided
	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := services.Login(username, password)
	if err != nil {
		// Return structured error in JSON format
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// Return the access and refresh tokens as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshTokenHandler handles token refresh requests
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	refreshToken := request["refresh_token"]
	if refreshToken == "" {
		http.Error(w, "Missing refresh token", http.StatusBadRequest)
		return
	}

	// Validate the refresh token
	claims, err := utils.ValidateToken(refreshToken, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Load expiry durations from environment variables
	accessTokenExpiry, err := parseDurationFromEnv("ACCESS_TOKEN_EXPIRY")
	if err != nil {
		http.Error(w, "Failed to load access token expiry", http.StatusInternalServerError)
		return
	}

	// Generate a new access token
	accessToken, err := utils.GenerateToken(claims.Username, os.Getenv("JWT_SECRET"), accessTokenExpiry)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Send response
	response := map[string]string{
		"access_token": accessToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
