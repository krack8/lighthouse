package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

// CreateUserHandler handles the creation of a new user.
func (uc *UserController) CreateUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := uc.UserService.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, createdUser)
}

// GetUserHandler handles fetching a user by ID.
func (uc *UserController) GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := uc.UserService.GetUser(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, user)
}

// GetAllUsersHandler handles fetching all users.
func (uc *UserController) GetAllUsersHandler(c *gin.Context) {
	userList, err := uc.UserService.GetAllUsers()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, userList)
}

// UpdateUserHandler handles updating a user's information.
func (uc *UserController) UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")

	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := uc.UserService.UpdateUser(id, &updatedData)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUserHandler handles deleting a user by ID.
func (uc *UserController) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	err := uc.UserService.DeleteUser(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
