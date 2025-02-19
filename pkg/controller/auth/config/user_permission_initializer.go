package config

import (
	"context"
	"fmt"
	enum2 "github.com/krack8/lighthouse/pkg/controller/auth/enum"
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
	"github.com/krack8/lighthouse/pkg/controller/auth/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
	"time"
)

// PermissionInitializer manages permission initialization
type PermissionInitializer struct {
	permissionCollection *mongo.Collection
	endpointRegistry     map[enum2.PermissionName]func() []models.Endpoint
	mu                   sync.RWMutex
}

// NewPermissionInitializer creates a new initializer instance
func NewPermissionInitializer(collection *mongo.Collection) *PermissionInitializer {
	pi := &PermissionInitializer{
		permissionCollection: collection,
		endpointRegistry:     make(map[enum2.PermissionName]func() []models.Endpoint),
	}
	pi.registerEndpoints()
	return pi
}

// registerEndpoints maps permission names to their endpoint functions. register additional
func (pi *PermissionInitializer) registerEndpoints() {
	pi.endpointRegistry = map[enum2.PermissionName]func() []models.Endpoint{
		enum2.DEFAULT_PERMISSION: utils.GetDefaultEndpoints,
		enum2.VIEW_USER:          utils.GetUserEndpoints,
		enum2.MANAGE_USER:        utils.GetManageUserEndpoints,
		enum2.VIEW_ROLE:          utils.GetRolesEndpoints,
		enum2.MANAGE_ROLE:        utils.GetManageRolesEndpoints,
		enum2.ADD_CLUSTER:        utils.GetAddClustersEndpoints,
		//enum.MANAGE_NAMESPACE:                  utils.GetManageEndpointsEndpoints,
		enum2.CREATE_NAMESPACE:                  utils.GetCreateNamespaceEndpoints,
		enum2.VIEW_NAMESPACE:                    utils.GetViewNamespaceEndpoints,
		enum2.UPDATE_NAMESPACE:                  utils.GetUpdateNamespaceEndpoints,
		enum2.DELETE_NAMESPACE:                  utils.GetDeleteNamespaceEndpoints,
		enum2.VIEW_DEPLOYMENT:                   utils.GetViewDeploymentEndpoints,
		enum2.VIEW_REPLICA_SET:                  utils.GetViewReplicaSetEndpoints,
		enum2.MANAGE_POD:                        utils.GetManagePodEndpoints,
		enum2.VIEW_POD:                          utils.GetViewPodEndpoints,
		enum2.MANAGE_DEPLOYMENT:                 utils.GetManageDeploymentEndpoints,
		enum2.MANAGE_REPLICA_SET:                utils.GetManageReplicaSetEndpoints,
		enum2.VIEW_STATEFUL_SET:                 utils.GetViewStatefulSetEndpoints,
		enum2.MANAGE_STATEFUL_SET:               utils.GetManageStatefulSetEndpoints,
		enum2.VIEW_DAEMON_SET:                   utils.GetViewDaemonSetEndpoints,
		enum2.MANAGE_DAEMON_SET:                 utils.GetManageDaemonSetEndpoints,
		enum2.VIEW_SECRET:                       utils.GetViewSecretEndpoints,
		enum2.MANAGE_SECRET:                     utils.GetManageSecretEndpoints,
		enum2.VIEW_CONFIG_MAP:                   utils.GetViewConfigMapEndpoints,
		enum2.MANAGE_CONFIG_MAP:                 utils.GetManageConfigMapEndpoints,
		enum2.VIEW_SERVICE_ACCOUNT:              utils.GetViewServiceAccountEndpoints,
		enum2.MANAGE_SERVICE_ACCOUNT:            utils.GetManageServiceAccountEndpoints,
		enum2.VIEW_SERVICE:                      utils.GetViewServiceEndpoints,
		enum2.MANAGE_SERVICE:                    utils.GetManageServiceEndpoints,
		enum2.VIEW_INGRESS:                      utils.GetViewIngressEndpoints,
		enum2.MANAGE_INGRESS:                    utils.GetManageIngressEndpoints,
		enum2.VIEW_CERTIFICATE:                  utils.GetViewCertificateEndpoints,
		enum2.MANAGE_CERTIFICATE:                utils.GetManageCertificateEndpoints,
		enum2.VIEW_NAMESPACE_ROLE:               utils.GetViewNamespaceRoleEndpoints,
		enum2.MANAGE_NAMESPACE_ROLE:             utils.GetManageNamespaceRoleEndpoints,
		enum2.VIEW_NAMESPACE_ROLE_BINDING:       utils.GetViewNamespaceRoleBindingEndpoints,
		enum2.MANAGE_NAMESPACE_ROLE_BINDING:     utils.GetManageNamespaceRoleBindingEndpoints,
		enum2.VIEW_JOB:                          utils.GetViewJobEndpoints,
		enum2.MANAGE_JOB:                        utils.GetManageJobEndpoints,
		enum2.VIEW_CRON_JOB:                     utils.GetViewCronJobEndpoints,
		enum2.MANAGE_CRON_JOB:                   utils.GetManageCronJobEndpoints,
		enum2.VIEW_NAMESPACE_NETWORK_POLICY:     utils.GetViewNetworkPolicyEndpoints,
		enum2.MANAGE_NAMESPACE_NETWORK_POLICY:   utils.GetManageNetworkPolicyEndpoints,
		enum2.VIEW_NAMESPACE_RESOURCE_QUOTA:     utils.GetViewResourceQuotaEndpoints,
		enum2.MANAGE_RESOURCE_QUOTA:             utils.GetManageResourceQuotaEndpoints,
		enum2.VIEW_PERSISTENT_VOLUME:            utils.GetViewPersistentVolumeEndpoints,
		enum2.MANAGE_PERSISTENT_VOLUME:          utils.GetManagePersistentVolumeEndpoints,
		enum2.VIEW_PERSISTENT_VOLUME_CLAIM:      utils.GetViewPersistentVolumeClaimEndpoints,
		enum2.MANAGE_PERSISTENT_VOLUME_CLAIM:    utils.GetManagePersistentVolumeClaimEndpoints,
		enum2.VIEW_GATEWAY:                      utils.GetViewGatewayEndpoints,
		enum2.MANAGE_GATEWAY:                    utils.GetManageGatewayEndpoints,
		enum2.VIEW_VIRTUAL_SERVICE:              utils.GetViewVirtualServiceEndpoints,
		enum2.MANAGE_VIRTUAL_SERVICE:            utils.GetManageVirtualServiceEndpoints,
		enum2.VIEW_NODES:                        utils.GetViewNodeEndpoints,
		enum2.MANAGE_NODE_TAINT:                 utils.GetManageNodeTaintEndpoints,
		enum2.DRAIN_NODE:                        utils.GetDrainNodeEndpoints,
		enum2.VIEW_CLUSTER_ROLE:                 utils.GetViewClusterRoleEndpoints,
		enum2.MANAGE_CLUSTER_ROLE:               utils.GetManageClusterRoleEndpoints,
		enum2.VIEW_CLUSTER_ROLE_BINDING:         utils.GetViewClusterRoleBindingEndpoints,
		enum2.MANAGE_CLUSTER_ROLE_BINDING:       utils.GetManageClusterRoleBindingEndpoints,
		enum2.VIEW_STORAGE_CLASS:                utils.GetViewStorageClassEndpoints,
		enum2.MANAGE_STORAGE_CLASS:              utils.GetManageStorageClassEndpoints,
		enum2.VIEW_CUSTOM_RESOURCES:             utils.GetViewCustomResourceEndpoints,
		enum2.MANAGE_CUSTOM_RESOURCES:           utils.GetManageCustomResourceEndpoints,
		enum2.VIEW_CUSTOM_RESOURCE_DEFINITION:   utils.GetViewCustomResourceDefinitionEndpoints,
		enum2.MANAGE_CUSTOM_RESOURCE_DEFINITION: utils.GetManageCustomResourceDefinitionEndpoints,
		enum2.VIEW_LOGS:                         utils.GetViewLogsEndpoints,
		//enum.VIEW_ENDPOINTS:                    utils.GetViewEndpointSliceEndpoints,
		enum2.MANAGE_ENDPOINTS:              utils.GetManageEndpointsEndpoints,
		enum2.VIEW_ENDPOINT_SLICE:           utils.GetViewEndpointSliceEndpoints,
		enum2.MANAGE_ENDPOINT_SLICE:         utils.GetManageEndpointSliceEndpoints,
		enum2.VIEW_PDB:                      utils.GetViewPDBEndpoints,
		enum2.MANAGE_PDB:                    utils.GetManagePDBEndpoints,
		enum2.VIEW_CONTROLLER_REVISION:      utils.GetViewControllerRevisionEndpoints,
		enum2.MANAGE_CONTROLLER_REVISION:    utils.GetManageControllerRevisionEndpoints,
		enum2.VIEW_REPLICATION_CONTROLLER:   utils.GetViewReplicationControllerEndpoints,
		enum2.MANAGE_REPLICATION_CONTROLLER: utils.GetManageReplicationControllerEndpoints,
		// Add more mappings
	}
}

// InitializePermissions initializes all permissions
func (pi *PermissionInitializer) InitializePermissions(ctx context.Context) error {
	for permName, def := range enum2.PermissionDefinitions {
		endpointFunc, exists := pi.endpointRegistry[permName]
		if !exists {
			return fmt.Errorf("no endpoint function registered for permission: %s", permName)
		}

		err := pi.initializePermission(ctx, models.Permission{
			Name:         string(permName),
			Description:  string(def.Description),
			Category:     def.Category,
			EndpointList: endpointFunc(),
			Status:       enum2.VALID,
			CreatedBy:    string(enum2.SYSTEM),
			UpdatedBy:    string(enum2.SYSTEM),
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
