package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/dto"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// CreateUserHandler handles the creation of a new user.
func (uc *UserController) CreateUserHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context.Please Enable AUTH"})
		return
	}
	requester := username.(string)

	var userDTO dto.UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Convert DTO to User model
	user, err := uc.convertDTOToUser(c, userDTO, requester)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := uc.UserService.CreateUser(user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, createdUser)
}

func (uc *UserController) convertDTOToUser(ctx context.Context, userDTO dto.UserDTO, requester string) (*models.User, error) {
	// Convert role IDs to Role objects
	roles, err := uc.UserService.GetRolesByIds(ctx, userDTO.RoleIds)
	if err != nil {
		return nil, err
	}

	if userDTO.UserType != string(models.AdminUser) {
		if len(roles) == 0 {
			roles, err = services.GetRoleByName("DEFAULT_ROLE")
			if err != nil {
				return nil, err
			}
		}
	}

	return &models.User{
		ID:            primitive.NewObjectID(),
		Username:      userDTO.Username,
		FirstName:     userDTO.FirstName,
		LastName:      userDTO.LastName,
		Password:      utils.HashPassword(userDTO.Password),
		UserType:      models.UserType(userDTO.UserType),
		Roles:         roles,
		ClusterIdList: userDTO.ClusterIdList,
		UserIsActive:  userDTO.UserIsActive,
		IsVerified:    userDTO.IsVerified,
		Phone:         userDTO.Phone,
		Status:        enum.VALID,
		CreatedBy:     requester,
		UpdatedBy:     requester,
	}, nil
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

// GetUserProfileInfoHandler fetch user details by token.
func (uc *UserController) GetUserProfileInfoHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context.Please Enable AUTH"})
		return
	}
	// username is of type interface{}, so cast it to string
	usernameStr := username.(string)
	user, err := uc.UserService.GetUserProfileInfo(usernameStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, user)
}

// ResetPasswordHandler handles the password reset with old password verification
func (uc *UserController) ResetPasswordHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context.Please Enable AUTH"})
		return
	}
	// username is of type interface{}, so cast it to string
	usernameStr := username.(string)

	userID := c.Param("id")
	if userID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "user ID is required")
		return
	}

	var req dto.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid request payload")
		return
	}

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid user ID format")
		return
	}

	err = uc.UserService.ResetPassword(objectID, req.CurrentPassword, req.NewPassword, usernameStr)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"message": "Password updated successfully",
	})
}

// ForgotPasswordHandler handles initiating the forgot password process
func (uc *UserController) ForgotPasswordHandler(c *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalid request payload")
		return
	}

	err := uc.UserService.InitiateForgotPassword(req.Username)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"message": "Password reset link sent to email",
	})
}
