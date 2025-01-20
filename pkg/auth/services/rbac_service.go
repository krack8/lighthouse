package services

import (
	"context"
	"errors"
	"fmt"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"time"
)

// CreatePermission creates a new permission
func CreatePermission(permission models.Permission) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.PermissionCollection.InsertOne(ctx, permission)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// CreateRole creates a new role
func CreateRole(role models.Role) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.RoleCollection.InsertOne(ctx, role)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// AssignRole assigns roles to a user.
func AssignRole(username string, roleNames []string) error {
	// Find user by username
	var user models.User
	err := db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	// Loop through roleNames and add roles
	var roles []models.Role
	for _, roleName := range roleNames {
		var role models.Role
		err := db.RoleCollection.FindOne(context.Background(), bson.M{"name": roleName}).Decode(&role)
		if err != nil {
			return fmt.Errorf("role '%s' not found", roleName)
		}
		roles = append(roles, role)
	}

	// Update roles for the user
	_, err = db.UserCollection.UpdateOne(
		context.Background(),
		bson.M{"username": username},
		bson.M{"$set": bson.M{"roles": roles}},
	)
	if err != nil {
		return err
	}

	return nil
}

// CheckPermission checks if a user has a specific permission for a route and method
func CheckPermission(permissions []string, route, method string) bool {
	permissionKey := route + ":" + method
	for _, perm := range permissions {
		if strings.EqualFold(perm, permissionKey) {
			return true
		}
	}
	return false
}

// Helper function to get the role by ID
func getRoleByID(roleID string) (*models.Role, error) {
	var role models.Role
	err := db.RoleCollection.FindOne(context.Background(), bson.M{"_id": roleID}).Decode(&role)
	if err != nil {
		return nil, errors.New("role not found")
	}
	return &role, nil
}
