package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/services"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RbacController struct {
	RbacService *services.RbacService
}

// CreatePermissionHandler handles the creation of a new permission
func (rbac *RbacController) CreatePermissionHandler(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&permission); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create Permission
	permissionID, err := rbac.RbacService.CreatePermission(permission)
	if err != nil {
		http.Error(w, "Error creating permission", http.StatusInternalServerError)
		return
	}

	// Respond with the ID of the created permission
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]primitive.ObjectID{"permission_id": permissionID})
}

// CreateRoleHandler handles the creation of a new role
func (rbac *RbacController) CreateRoleHandler(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create Role
	roleID, err := rbac.RbacService.CreateRole(role)
	if err != nil {
		http.Error(w, "Error creating role", http.StatusInternalServerError)
		return
	}

	// Respond with the ID of the created role
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]primitive.ObjectID{"role_id": roleID})
}

// AssignRolesHandler assigns multiple roles to a user.
func (rbac *RbacController) AssignRolesHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Username string   `json:"username"`
		Roles    []string `json:"roles"`
	}

	// Parse the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service to assign the roles
	err := rbac.RbacService.AssignRole(request.Username, request.Roles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Roles %v assigned to user '%s'\n", request.Roles, request.Username)
}
