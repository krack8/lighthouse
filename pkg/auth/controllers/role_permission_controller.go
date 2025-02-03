package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"net/http"
	"strings"
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
