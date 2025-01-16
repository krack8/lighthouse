package controllers

import (
	"encoding/json"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody map[string]string
	json.NewDecoder(r.Body).Decode(&requestBody)

	username := requestBody["username"]
	password := requestBody["password"]

	accessToken, refreshToken, err := services.Login(db, username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

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
	claims, err := utils.ValidateToken(refreshToken, "your-refresh-secret-key")
	if err != nil {
		http.Error(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	// Generate a new access token
	accessToken, err := utils.GenerateToken(claims.Username, "your-secret-key", time.Hour*1) // 1-hour expiry
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
