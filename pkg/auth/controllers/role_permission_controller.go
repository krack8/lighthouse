package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/dto"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RbacController struct {
	RbacService *services.RbacService
}

func NewRbacController(rbacService *services.RbacService) *RbacController {
	return &RbacController{
		RbacService: rbacService,
	}
}

// CreatePermissionHandler handles the creation of a new permission
func (rbac *RbacController) CreatePermissionHandler(c *gin.Context) {
	var permission models.Permission
	if err := c.ShouldBindJSON(&permission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create Permission
	permissionID, err := rbac.RbacService.CreatePermission(permission)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating permission"})
		return
	}

	// Respond with the ID of the created permission
	c.JSON(http.StatusCreated, gin.H{"permission_id": permissionID})
}

// Helper function to convert permission ID strings to Permission slice
func convertPermissionsToModels(permissionIDs []string) ([]models.Permission, error) {
	// Create a slice to store ObjectIDs
	objectIDs := make([]primitive.ObjectID, 0, len(permissionIDs))

	// Convert string IDs to ObjectIDs
	for _, idStr := range permissionIDs {
		if strings.TrimSpace(idStr) == "" {
			continue
		}
		objID, err := primitive.ObjectIDFromHex(strings.TrimSpace(idStr))
		if err != nil {
			return nil, fmt.Errorf("invalid permission ID: %s - %v", idStr, err)
		}
		objectIDs = append(objectIDs, objID)
	}

	// If no valid IDs, return empty result
	if len(objectIDs) == 0 {
		return []models.Permission{}, nil
	}

	// Create filter for MongoDB query
	filter := bson.M{
		"_id":    bson.M{"$in": objectIDs},
		"status": enum.VALID, // Assuming you want only active permissions
	}

	// Find permissions
	cursor, err := db.PermissionCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fetching permissions: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	// Decode results into Permission models
	var permissions []models.Permission
	if err := cursor.All(context.Background(), &permissions); err != nil {
		return nil, fmt.Errorf("error decoding permissions: %v", err)
	}

	return permissions, nil
}

// CreateRoleHandler handles the creation of a new role
func (rbac *RbacController) CreateRoleHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context.Please Enable AUTH"})
		return
	}
	requester := username.(string)
	var roleDTO dto.RoleDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate permissions slice
	if len(roleDTO.Permissions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permissions cannot be empty"})
		return
	}

	PermissionList, e := convertPermissionsToModels(roleDTO.Permissions)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to Convert Permission Model"})
		return
	}

	// Convert DTO to model
	role := models.Role{
		ID:          primitive.NewObjectID(),
		Name:        roleDTO.Name,
		Description: roleDTO.Description,
		Permissions: PermissionList,
		Status:      enum.VALID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   requester,
		UpdatedBy:   requester,
	}

	// Create Role
	roleID, err := rbac.RbacService.CreateRole(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating role"})
		return
	}
	// Respond with the ID of the created role
	c.JSON(http.StatusCreated, gin.H{"role_id": roleID})
}

// AssignRolesHandler assigns multiple roles to a user.
func (rbac *RbacController) AssignRolesHandler(c *gin.Context) {
	var request struct {
		Username string   `json:"username"`
		RoleIds  []string `json:"roleIds"`
	}

	// Parse the JSON request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service to assign the roles
	err := rbac.RbacService.AssignRole(request.Username, request.RoleIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Roles %v assigned to user '%s'", request.RoleIds, request.Username)})
}

// GetAllRolesHandler retrieves all roles
func (rbac *RbacController) GetAllRolesHandler(c *gin.Context) {
	// Call the service to get all roles
	roles, err := rbac.RbacService.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving roles"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"roles": roles})
}

// GetRoleByIDHandler retrieves a specific role by its ID
func (rbac *RbacController) GetRoleByIDHandler(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	// Call the service to get the role by ID
	role, err := rbac.RbacService.GetRoleByID(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving role"})
		return
	}

	if role == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
}

// DeleteRoleHandler handles the deletion of a role by its ID
func (rbac *RbacController) DeleteRoleHandler(c *gin.Context) {
	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	// Call the service to delete the role
	err := rbac.RbacService.DeleteRoleByID(roleID)
	if err != nil {
		if strings.Contains(err.Error(), "no role found") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Role %s deleted successfully", roleID)})
}

// GetAllPermissionsHandler retrieves all permissions
func (rbac *RbacController) GetAllPermissionsHandler(c *gin.Context) {
	// Call the service to get all permissions
	permissions, err := rbac.RbacService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

// GetPermissionByIDHandler retrieves a specific permission by its ID
func (rbac *RbacController) GetPermissionByIDHandler(c *gin.Context) {
	permissionID := c.Param("id")
	if permissionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission ID is required"})
		return
	}

	// Call the service to get the permission by ID
	permission, err := rbac.RbacService.GetPermissionByID(permissionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permission"})
		return
	}

	if permission == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permission": permission})
}

// GetPermissionsByCategoryHandler retrieves permissions by their category
func (rbac *RbacController) GetPermissionsByCategoryHandler(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category is required"})
		return
	}

	// Call the service to get permissions by category
	permissions, err := rbac.RbacService.GetPermissionsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	if len(permissions) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No permissions found for the given category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

// GetUserPermissionsHandler handles the permission request for a user
func (rbac *RbacController) GetUserPermissionsHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context.Please Enable AUTH"})
		return
	}
	// username is of type interface{}, so cast it to string
	usernameStr := username.(string)
	// Get permissions
	permissions, err := rbac.RbacService.GetPermissionsByUser(usernameStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving permissions"})
		return
	}

	// Return permissions directly without the "data" wrapper
	c.JSON(http.StatusOK, permissions)
}

func (rbac *RbacController) UpdateRoleHandler(c *gin.Context) {
	// Get role ID from URL parameter
	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	// Get username from context
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username not found in context. Please Enable AUTH"})
		return
	}
	requester := username.(string)

	// Bind input data
	var roleDTO dto.RoleDTO
	if err := c.ShouldBindJSON(&roleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate permissions
	if len(roleDTO.Permissions) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permissions cannot be empty"})
		return
	}

	// Convert permission IDs to models
	permissionList, err := convertPermissionsToModels(roleDTO.Permissions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to Convert Permission Model"})
		return
	}

	// Convert ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}

	// Create update model
	updateRole := models.Role{
		ID:          objectID,
		Name:        roleDTO.Name,
		Description: roleDTO.Description,
		Permissions: permissionList,
		Status:      enum.VALID,
		UpdatedAt:   time.Now(),
		UpdatedBy:   requester,
	}

	// Call service to update
	err = rbac.RbacService.UpdateRole(updateRole)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating role"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role updated successfully"})
}

func (rbac *RbacController) GetUsersByRoleIDHandler(c *gin.Context) {
	// Get role ID from URL parameter
	roleID := c.Param("id")
	if roleID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	// Convert ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID format"})
		return
	}

	// Default values
	const (
		defaultPage  = 1
		defaultLimit = 10
		maxLimit     = 100
	)

	// Get page parameter with default value
	page := defaultPage
	pageStr := c.Query("page")
	if pageStr != "" {
		parsedPage, err := strconv.Atoi(pageStr)
		if err != nil || parsedPage < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}
		page = parsedPage
	}

	// Get limit parameter with default value
	limit := defaultLimit
	limitStr := c.Query("limit")
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}
		limit = parsedLimit
	}

	// Enforce maximum limit
	if limit > maxLimit {
		limit = maxLimit
	}

	// Call service to get users
	users, total, err := rbac.RbacService.GetUsersByRoleID(objectID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
