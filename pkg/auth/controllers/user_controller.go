package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

// CreateUserHandler handles the creation of a new user.
func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := uc.UserService.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdUser)
}

// GetUserHandler handles fetching a user by ID.
func (uc *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	user, err := uc.UserService.GetUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

// GetUserHandler handles fetching a user by ID.
func (uc *UserController) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userList, err := uc.UserService.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, userList)
}

// UpdateUserHandler handles updating a user's information.
func (uc *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := uc.UserService.UpdateUser(id, &updatedData)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write([]byte("User updated successfully"))
	utils.RespondWithJSON(w, http.StatusOK, nil)
}

// DeleteUserHandler handles deleting a user by ID.
func (uc *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	err := uc.UserService.DeleteUser(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Write([]byte("User deleted successfully"))
	utils.RespondWithJSON(w, http.StatusOK, nil)
}
