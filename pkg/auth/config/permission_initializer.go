package config

import (
	"context"
	"fmt"
	"github.com/krack8/lighthouse/pkg/auth/enum"
	"github.com/krack8/lighthouse/pkg/auth/models"
	"github.com/krack8/lighthouse/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

// PermissionInitializer manages permission initialization
type PermissionInitializer struct {
	permissionCollection *mongo.Collection
	endpointRegistry     map[enum.PermissionName]func() []models.Endpoint
	mu                   sync.RWMutex
}

// NewPermissionInitializer creates a new initializer instance
func NewPermissionInitializer(collection *mongo.Collection) *PermissionInitializer {
	pi := &PermissionInitializer{
		permissionCollection: collection,
		endpointRegistry:     make(map[enum.PermissionName]func() []models.Endpoint),
	}
	pi.registerEndpoints()
	return pi
}

// registerEndpoints maps permission names to their endpoint functions
func (pi *PermissionInitializer) registerEndpoints() {
	pi.endpointRegistry = map[enum.PermissionName]func() []models.Endpoint{
		enum.VIEW_NAMESPACE:   utils.GetViewNamespaceEndpoints,
		enum.MANAGE_NAMESPACE: utils.GetManageEndpointsEndpoints,
		// Add more mappings
	}
}

// InitializePermissions initializes all permissions
func (pi *PermissionInitializer) InitializePermissions(ctx context.Context) error {
	for permName, def := range enum.PermissionDefinitions {
		endpointFunc, exists := pi.endpointRegistry[permName]
		if !exists {
			return fmt.Errorf("no endpoint function registered for permission: %s", permName)
		}

		err := pi.initializePermission(ctx, models.Permission{
			Name:         string(permName),
			Description:  string(def.Description),
			Category:     def.Category,
			EndpointList: endpointFunc(),
			CreatedBy:    "SYSTEM",
			UpdatedBy:    "SYSTEM",
		})
		if err != nil {
			return fmt.Errorf("failed to initialize permission %s: %v", permName, err)
		}
	}
	return nil
}

// initializePermission handles initialization of a single permission
func (pi *PermissionInitializer) initializePermission(ctx context.Context, perm models.Permission) error {
	pi.mu.Lock()
	defer pi.mu.Unlock()

	exists, err := pi.permissionExists(ctx, perm.Name)
	if err != nil {
		return err
	}

	if !exists {
		perm.ID = primitive.NewObjectID()
		perm.CreatedAt = time.Now()
		perm.UpdatedAt = time.Now()

		_, err = pi.permissionCollection.InsertOne(ctx, perm)
		return err
	}

	return pi.updatePermissionEndpoints(ctx, perm.Name, perm.EndpointList)
}

// permissionExists checks if a permission already exists
func (pi *PermissionInitializer) permissionExists(ctx context.Context, name string) (bool, error) {
	count, err := pi.permissionCollection.CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// updatePermissionEndpoints updates endpoints for an existing permission
func (pi *PermissionInitializer) updatePermissionEndpoints(ctx context.Context, name string, endpoints []models.Endpoint) error {
	filter := bson.M{"name": name}
	update := bson.M{
		"$set": bson.M{
			"endpoint-list": endpoints,
			"updated_at":    time.Now(),
			"updated_by":    "SYSTEM",
		},
	}
	_, err := pi.permissionCollection.UpdateOne(ctx, filter, update)
	return err
}

// GetPermissions returns all initialized permissions
func (pi *PermissionInitializer) GetPermissions(ctx context.Context) ([]models.Permission, error) {
	var permissions []models.Permission
	cursor, err := pi.permissionCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	return permissions, cursor.All(ctx, &permissions)
}
