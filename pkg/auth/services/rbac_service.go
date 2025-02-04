package services

import (
	"context"
	"errors"
	"fmt"
	db "github.com/krack8/lighthouse/pkg/auth/config"
	"github.com/krack8/lighthouse/pkg/auth/dto"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

// RBAC service handles rbac-related business logic
type RbacService struct {
	collection Collection
}

// NewRbacService creates a new RbacService instance
func NewRbacService(collection Collection) *RbacService {
	return &RbacService{
		collection: collection,
	}
}

// CreatePermission creates a new permission
func (r *RbacService) CreatePermission(permission models.Permission) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.PermissionCollection.InsertOne(ctx, permission)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// CreateRole creates a new role
func (r *RbacService) CreateRole(role models.Role) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Additional validation if needed
	if len(role.Permissions) == 0 {
		return primitive.NilObjectID, errors.New("permissions cannot be empty")
	}

	// Validate each permission
	for _, perm := range role.Permissions {
		if strings.TrimSpace(perm.Name) == "" {
			return primitive.NilObjectID, errors.New("invalid permission name")
		}
	}

	result, err := db.RoleCollection.InsertOne(ctx, role)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// AssignRole assigns roles to a user.
func (r *RbacService) AssignRole(username string, roleIds []string) error {
	// Find user by username
	var user models.User
	err := db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	// Loop through roleNames and add roles
	var roles []models.Role
	for _, roleId := range roleIds {
		objectID, err := primitive.ObjectIDFromHex(roleId)
		var role models.Role
		err = db.RoleCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&role)
		if err != nil {
			return fmt.Errorf("role '%s' not found with Id", objectID)
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
	// Normalize the input
	method = strings.ToUpper(method)
	permissionKey := route + ":" + method

	// Check for exact matches first
	for _, perm := range permissions {
		if strings.EqualFold(perm, permissionKey) {
			return true
		}
	}

	return false
}

// GetAllRoles retrieves all roles
func (r *RbacService) GetAllRoles() ([]models.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all roles
	cursor, err := db.RoleCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve roles: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	// Slice to store roles
	var roles []models.Role

	// Decode all roles
	if err = cursor.All(ctx, &roles); err != nil {
		return nil, fmt.Errorf("failed to decode roles: %v", err)
	}

	return roles, nil
}

// GetRoleByID retrieves a specific role by its ID
func (r *RbacService) GetRoleByID(roleID string) (*models.Role, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return nil, fmt.Errorf("invalid role ID format: %v", err)
	}

	// Find the role by ID
	var role models.Role
	err = db.RoleCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&role)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve role: %v", err)
	}

	return &role, nil
}

// DeleteRoleByID deletes a role by its ID
func (r *RbacService) DeleteRoleByID(roleID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(roleID)
	if err != nil {
		return fmt.Errorf("invalid role ID format: %v", err)
	}

	// Delete the role
	result, err := db.RoleCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return fmt.Errorf("failed to delete role: %v", err)
	}

	// Check if a role was actually deleted
	if result.DeletedCount == 0 {
		return fmt.Errorf("no role found with ID: %s", roleID)
	}

	return nil
}

// GetAllPermissions retrieves all permissions
func (r *RbacService) GetAllPermissions() ([]models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find all permissions
	cursor, err := db.PermissionCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve permissions: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	// Slice to store permissions
	var permissions []models.Permission

	// Decode all permissions
	if err = cursor.All(ctx, &permissions); err != nil {
		return nil, fmt.Errorf("failed to decode permissions: %v", err)
	}

	return permissions, nil
}

// GetPermissionByID retrieves a specific permission by its ID
func (r *RbacService) GetPermissionByID(permissionID string) (*models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(permissionID)
	if err != nil {
		return nil, fmt.Errorf("invalid permission ID format: %v", err)
	}

	// Find the permission by ID
	var permission models.Permission
	err = db.PermissionCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&permission)
	if err != nil {

		return nil, fmt.Errorf("failed to retrieve permission: %v", err)
	}

	return &permission, nil
}

// GetPermissionsByCategory retrieves permissions by their category
func (r *RbacService) GetPermissionsByCategory(category string) ([]models.Permission, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find permissions by category
	cursor, err := db.PermissionCollection.Find(ctx, bson.M{"category": category})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve permissions by category: %v", err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, ctx)

	// Slice to store permissions
	var permissions []models.Permission

	// Decode all permissions
	if err = cursor.All(ctx, &permissions); err != nil {
		return nil, fmt.Errorf("failed to decode permissions: %v", err)
	}

	return permissions, nil
}

func (r *RbacService) GetPermissionsByUserType(username string) (*dto.PermissionResponse, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	user, _ := GetUserByUsername(username)

	if user == nil {
		return nil, errors.New("user do not exists")
	}

	// Initialize response
	response := &dto.PermissionResponse{
		DEFAULT:  make([]dto.PermissionDTO, 0),
		CLUSTER:  make([]dto.PermissionDTO, 0),
		HelmApps: make([]dto.PermissionDTO, 0),
	}

	// Create filter for Valid permissions
	filter := bson.M{
		"status": "V",
	}

	// Fetch permissions from database
	cursor, err := db.PermissionCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {

		}
	}(cursor, context.Background())

	// Process permissions
	var permissions []models.Permission
	if err := cursor.All(context.Background(), &permissions); err != nil {
		return nil, err
	}

	// Group permissions by category
	for _, perm := range permissions {
		dto := dto.PermissionDTO{
			ID:          perm.ID,
			Name:        perm.Name,
			Description: perm.Description,
		}

		switch perm.Category {
		case enum.DEFAULT:
			response.DEFAULT = append(response.DEFAULT, dto)
		case enum.CLUSTER:
			response.CLUSTER = append(response.CLUSTER, dto)
		case enum.HELM:
			response.HelmApps = append(response.HelmApps, dto)
		}
		//add category here
	}

	return response, nil
}

func (r *RbacService) UpdateRole(role models.Role) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Validate permissions
	if len(role.Permissions) == 0 {
		return errors.New("permissions cannot be empty")
	}

	// Create update filter
	filter := bson.M{
		"_id":    role.ID,
		"status": enum.VALID,
	}

	// Create update document
	update := bson.M{
		"$set": bson.M{
			"name":        role.Name,
			"description": role.Description,
			"permissions": role.Permissions,
			"updated_at":  role.UpdatedAt,
			"updated_by":  role.UpdatedBy,
		},
	}

	// Perform update
	result, err := db.RoleCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	// Check if document was found and updated
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (r *RbacService) GetUsersByRoleID(roleID primitive.ObjectID, page, limit int) ([]models.User, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Calculate skip value for pagination
	skip := (page - 1) * limit

	// Create match stage for aggregation
	matchStage := bson.D{
		{"$match", bson.D{
			{"roles._id", roleID},
			{"status", enum.VALID},
			{"user_is_active", true},
		}},
	}

	// Create pagination stages
	paginationStage := bson.D{
		{"$skip", skip},
	}
	limitStage := bson.D{
		{"$limit", limit},
	}

	// Execute count query
	countPipeline := mongo.Pipeline{matchStage}
	countCursor, err := db.UserCollection.Aggregate(ctx, append(countPipeline, bson.D{
		{"$count", "total"},
	}))
	if err != nil {
		return nil, 0, err
	}
	defer countCursor.Close(ctx)

	// Get total count
	var countResult []bson.M
	if err := countCursor.All(ctx, &countResult); err != nil {
		return nil, 0, err
	}

	total := int64(0)
	if len(countResult) > 0 {
		total = countResult[0]["total"].(int64)
	}

	// Execute main query
	pipeline := mongo.Pipeline{
		matchStage,
		paginationStage,
		limitStage,
	}

	cursor, err := db.UserCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode results
	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
