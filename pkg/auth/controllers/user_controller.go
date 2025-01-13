package controllers

import (
	"encoding/json"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Generate unique ID and timestamp
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()

	if err := services.CreateUser(&user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}
