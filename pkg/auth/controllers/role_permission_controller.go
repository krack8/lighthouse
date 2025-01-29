package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"net/http"
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

// CreateRoleHandler handles the creation of a new role
func (rbac *RbacController) CreateRoleHandler(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
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
		Roles    []string `json:"roles"`
	}

	// Parse the JSON request body
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call the service to assign the roles
	err := rbac.RbacService.AssignRole(request.Username, request.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Roles %v assigned to user '%s'", request.Roles, request.Username)})
}
