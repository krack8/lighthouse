package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
)

// CreateUserHandler handles the creation of a new user.
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := services.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}

// GetUserHandler handles fetching a user by ID.
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, err := services.GetUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// GetUserHandler handles fetching a user by ID.
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userList, err := services.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, userList)
}

// UpdateUserHandler handles updating a user's information.
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := services.UpdateUser(id, &updatedData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write([]byte("User updated successfully"))
	utils.RespondWithJSON(w, http.StatusOK, nil)
}

// DeleteUserHandler handles deleting a user by ID.
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := services.DeleteUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write([]byte("User deleted successfully"))
	utils.RespondWithJSON(w, http.StatusOK, nil)
}
