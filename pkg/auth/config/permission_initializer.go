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

// registerEndpoints maps permission names to their endpoint functions. register additional
func (pi *PermissionInitializer) registerEndpoints() {
	pi.endpointRegistry = map[enum.PermissionName]func() []models.Endpoint{
		enum.DEFAULT_PERMISSION: utils.GetDefaultEndpoints,
		enum.VIEW_USER:          utils.GetUserEndpoints,
		enum.MANAGE_USER:        utils.GetManageUserEndpoints,
		enum.VIEW_ROLE:          utils.GetRolesEndpoints,
		enum.MANAGE_ROLE:        utils.GetManageRolesEndpoints,
		//enum.MANAGE_NAMESPACE:                  utils.GetManageEndpointsEndpoints,
		enum.CREATE_NAMESPACE:                  utils.GetCreateNamespaceEndpoints,
		enum.VIEW_NAMESPACE:                    utils.GetViewNamespaceEndpoints,
		enum.UPDATE_NAMESPACE:                  utils.GetUpdateNamespaceEndpoints,
		enum.DELETE_NAMESPACE:                  utils.GetDeleteNamespaceEndpoints,
		enum.VIEW_DEPLOYMENT:                   utils.GetViewDeploymentEndpoints,
		enum.VIEW_REPLICA_SET:                  utils.GetViewReplicaSetEndpoints,
		enum.MANAGE_POD:                        utils.GetManagePodEndpoints,
		enum.VIEW_POD:                          utils.GetViewPodEndpoints,
		enum.MANAGE_DEPLOYMENT:                 utils.GetManageDeploymentEndpoints,
		enum.MANAGE_REPLICA_SET:                utils.GetManageReplicaSetEndpoints,
		enum.VIEW_STATEFUL_SET:                 utils.GetViewStatefulSetEndpoints,
		enum.MANAGE_STATEFUL_SET:               utils.GetManageStatefulSetEndpoints,
		enum.VIEW_DAEMON_SET:                   utils.GetViewDaemonSetEndpoints,
		enum.MANAGE_DAEMON_SET:                 utils.GetManageDaemonSetEndpoints,
		enum.VIEW_SECRET:                       utils.GetViewSecretEndpoints,
		enum.MANAGE_SECRET:                     utils.GetManageSecretEndpoints,
		enum.VIEW_CONFIG_MAP:                   utils.GetViewConfigMapEndpoints,
		enum.MANAGE_CONFIG_MAP:                 utils.GetManageConfigMapEndpoints,
		enum.VIEW_SERVICE_ACCOUNT:              utils.GetViewServiceAccountEndpoints,
		enum.MANAGE_SERVICE_ACCOUNT:            utils.GetManageServiceAccountEndpoints,
		enum.VIEW_SERVICE:                      utils.GetViewServiceEndpoints,
		enum.MANAGE_SERVICE:                    utils.GetManageServiceEndpoints,
		enum.VIEW_INGRESS:                      utils.GetViewIngressEndpoints,
		enum.MANAGE_INGRESS:                    utils.GetManageIngressEndpoints,
		enum.VIEW_CERTIFICATE:                  utils.GetViewCertificateEndpoints,
		enum.MANAGE_CERTIFICATE:                utils.GetManageCertificateEndpoints,
		enum.VIEW_NAMESPACE_ROLE:               utils.GetViewNamespaceRoleEndpoints,
		enum.MANAGE_NAMESPACE_ROLE:             utils.GetManageNamespaceRoleEndpoints,
		enum.VIEW_NAMESPACE_ROLE_BINDING:       utils.GetViewNamespaceRoleBindingEndpoints,
		enum.MANAGE_NAMESPACE_ROLE_BINDING:     utils.GetManageNamespaceRoleBindingEndpoints,
		enum.VIEW_JOB:                          utils.GetViewJobEndpoints,
		enum.MANAGE_JOB:                        utils.GetManageJobEndpoints,
		enum.VIEW_CRON_JOB:                     utils.GetViewCronJobEndpoints,
		enum.MANAGE_CRON_JOB:                   utils.GetManageCronJobEndpoints,
		enum.VIEW_NAMESPACE_NETWORK_POLICY:     utils.GetViewNetworkPolicyEndpoints,
		enum.MANAGE_NAMESPACE_NETWORK_POLICY:   utils.GetManageNetworkPolicyEndpoints,
		enum.VIEW_NAMESPACE_RESOURCE_QUOTA:     utils.GetViewResourceQuotaEndpoints,
		enum.MANAGE_RESOURCE_QUOTA:             utils.GetManageResourceQuotaEndpoints,
		enum.VIEW_PERSISTENT_VOLUME:            utils.GetViewPersistentVolumeEndpoints,
		enum.MANAGE_PERSISTENT_VOLUME:          utils.GetManagePersistentVolumeEndpoints,
		enum.VIEW_PERSISTENT_VOLUME_CLAIM:      utils.GetViewPersistentVolumeClaimEndpoints,
		enum.MANAGE_PERSISTENT_VOLUME_CLAIM:    utils.GetManagePersistentVolumeClaimEndpoints,
		enum.VIEW_GATEWAY:                      utils.GetViewGatewayEndpoints,
		enum.MANAGE_GATEWAY:                    utils.GetManageGatewayEndpoints,
		enum.VIEW_VIRTUAL_SERVICE:              utils.GetViewVirtualServiceEndpoints,
		enum.MANAGE_VIRTUAL_SERVICE:            utils.GetManageVirtualServiceEndpoints,
		enum.VIEW_NODES:                        utils.GetViewNodeEndpoints,
		enum.MANAGE_NODE_TAINT:                 utils.GetManageNodeTaintEndpoints,
		enum.DRAIN_NODE:                        utils.GetDrainNodeEndpoints,
		enum.VIEW_CLUSTER_ROLE:                 utils.GetViewClusterRoleEndpoints,
		enum.MANAGE_CLUSTER_ROLE:               utils.GetManageClusterRoleEndpoints,
		enum.VIEW_CLUSTER_ROLE_BINDING:         utils.GetViewClusterRoleBindingEndpoints,
		enum.MANAGE_CLUSTER_ROLE_BINDING:       utils.GetManageClusterRoleBindingEndpoints,
		enum.VIEW_STORAGE_CLASS:                utils.GetViewStorageClassEndpoints,
		enum.MANAGE_STORAGE_CLASS:              utils.GetManageStorageClassEndpoints,
		enum.VIEW_CUSTOM_RESOURCES:             utils.GetViewCustomResourceEndpoints,
		enum.MANAGE_CUSTOM_RESOURCES:           utils.GetManageCustomResourceEndpoints,
		enum.VIEW_CUSTOM_RESOURCE_DEFINITION:   utils.GetViewCustomResourceDefinitionEndpoints,
		enum.MANAGE_CUSTOM_RESOURCE_DEFINITION: utils.GetManageCustomResourceDefinitionEndpoints,
		enum.VIEW_LOGS:                         utils.GetViewLogsEndpoints,
		//enum.VIEW_ENDPOINTS:                    utils.GetViewEndpointSliceEndpoints,
		enum.MANAGE_ENDPOINTS:              utils.GetManageEndpointsEndpoints,
		enum.VIEW_ENDPOINT_SLICE:           utils.GetViewEndpointSliceEndpoints,
		enum.MANAGE_ENDPOINT_SLICE:         utils.GetManageEndpointSliceEndpoints,
		enum.VIEW_PDB:                      utils.GetViewPDBEndpoints,
		enum.MANAGE_PDB:                    utils.GetManagePDBEndpoints,
		enum.VIEW_CONTROLLER_REVISION:      utils.GetViewControllerRevisionEndpoints,
		enum.MANAGE_CONTROLLER_REVISION:    utils.GetManageControllerRevisionEndpoints,
		enum.VIEW_REPLICATION_CONTROLLER:   utils.GetViewReplicationControllerEndpoints,
		enum.MANAGE_REPLICATION_CONTROLLER: utils.GetManageReplicationControllerEndpoints,
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
			Status:       enum.VALID,
			CreatedBy:    string(enum.SYSTEM),
			UpdatedBy:    string(enum.SYSTEM),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
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
