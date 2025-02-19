package utils

import (
	"github.com/krack8/lighthouse/pkg/controller/auth/models"
)

// GetDefaultEndpoints returns endpoints for creating  namespaces
func GetDefaultEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/clusters"},
		{Method: "GET", Route: "/api/v1/clusters/@"},
		{Method: "GET", Route: "/api/v1/users/profile"},
		{Method: "GET", Route: "/api/v1/users/@/reset-password"},
		{Method: "GET", Route: "/api/v1/permissions"},
		{Method: "GET", Route: "/api/v1/permissions/@"},
		{Method: "GET", Route: "/api/v1/permissions/users"},
	}
}

// GetUserEndpoints returns endpoints for creating  namespaces
func GetUserEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/users"},
		{Method: "GET", Route: "/api/v1/users/@"},
	}
}

// GetManageUserEndpoints returns endpoints for creating  namespaces
func GetManageUserEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/users"},
		{Method: "DELETE", Route: "/api/v1/users/@"},
	}
}

// GetRolesEndpoints returns endpoints for creating  namespaces
func GetRolesEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/roles"},
		{Method: "GET", Route: "/api/v1/roles/@"},
	}
}

// GetManageRolesEndpoints returns endpoints for creating  namespaces
func GetManageRolesEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/roles"},
		{Method: "GET", Route: "/api/v1/roles"},
		{Method: "GET", Route: "/api/v1/roles/@"},
		{Method: "PUT", Route: "/api/v1/roles/@"},
		{Method: "DELETE", Route: "/api/v1/roles/@"},
		{Method: "GET", Route: "/api/v1/roles/@/users"},
		{Method: "POST", Route: "/api/v1/assign-roles"},
	}
}

// GetManageClustersEndpoints returns endpoints for creating  cluster
func GetAddClustersEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/clusters"},
		{Method: "GET", Route: "/api/v1/clusters/@/details"},
	}
}

// GetCreateNamespaceEndpoints returns endpoints for creating  namespaces
func GetCreateNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
	}
}

// GetViewNamespaceEndpoints returns endpoints for viewing  namespaces
func GetViewNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace/names"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
		{Method: "GET", Route: "/api/v1/event"},
	}
}

// GetUpdateNamespaceEndpoints returns endpoints for updating  namespaces
func GetUpdateNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
	}
}

// GetDeleteNamespaceEndpoints returns endpoints for deleting  namespaces
func GetDeleteNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "DELETE", Route: "/api/v1/namespace/@"},
	}
}

// GetViewDeploymentEndpoints returns endpoints for viewing  namespace deployments
func GetViewDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/deployment"},
		{Method: "GET", Route: "/api/v1/deployment/stats"},
		{Method: "GET", Route: "/api/v1/deployment/@"},
		{Method: "GET", Route: "/api/v1/deployment/@/pods"},
	}
}

// GetManageDeploymentEndpoints returns endpoints for managing  namespace deployments
func GetManageDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/deployment"},
		{Method: "DELETE", Route: "/api/v1/deployment/@"},
	}
}

// GetViewPodEndpoints returns endpoints for viewing  namespace pods
func GetViewPodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pod"},
		{Method: "GET", Route: "/api/v1/pod/@"},
		{Method: "GET", Route: "/api/v1/pod/stats"},
	}
}

// GetManagePodEndpoints returns endpoints for managing  namespace pods
func GetManagePodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/pod"},
		{Method: "DELETE", Route: "/api/v1/pod/@"},
	}
}

// GetViewReplicaSetEndpoints returns endpoints for viewing  namespace replica sets
func GetViewReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/replicaset"},
		{Method: "GET", Route: "/api/v1/replicaset/@"},
	}
}

// GetManageReplicaSetEndpoints returns endpoints for managing  namespace replica sets
func GetManageReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/replicaset"},
		{Method: "DELETE", Route: "/api/v1/replicaset/@"},
	}
}

// GetViewStatefulSetEndpoints returns endpoints for viewing  namespace stateful sets
func GetViewStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/statefulset"},
		{Method: "GET", Route: "/api/v1/statefulset/@"},
		{Method: "GET", Route: "/api/v1/statefulset/@/pods"},
		{Method: "GET", Route: "/api/v1/statefulset/stats"},
	}
}

// GetManageStatefulSetEndpoints returns endpoints for managing  namespace stateful sets
func GetManageStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/statefulset"},
		{Method: "DELETE", Route: "/api/v1/statefulset/@"},
	}
}

// GetViewDaemonSetEndpoints returns endpoints for viewing daemon sets
func GetViewDaemonSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/daemonset"},
		{Method: "GET", Route: "/api/v1/daemonset/stats"},
		{Method: "GET", Route: "/api/v1/daemonset/@"},
	}
}

// GetManageDaemonSetEndpoints returns endpoints for managing daemon sets
func GetManageDaemonSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/daemonset"},
		{Method: "DELETE", Route: "/api/v1/daemonset/@"},
	}
}

// GetViewSecretEndpoints returns endpoints for viewing  namespace secrets
func GetViewSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/secret"},
		{Method: "GET", Route: "/api/v1/secret/@"},
	}
}

// GetManageSecretEndpoints returns endpoints for managing  namespace secrets
func GetManageSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/secret"},
		{Method: "DELETE", Route: "/api/v1/secret/@"},
	}
}

// GetViewConfigMapEndpoints returns endpoints for viewing  namespace config maps
func GetViewConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/config-map"},
		{Method: "GET", Route: "/api/v1/config-map/@"},
	}
}

// GetManageConfigMapEndpoints returns endpoints for managing  namespace config maps
func GetManageConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/config-map"},
		{Method: "DELETE", Route: "/api/v1/config-map/@"},
	}
}

// GetViewServiceEndpoints returns endpoints for viewing  namespace services
func GetViewServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service"},
		{Method: "GET", Route: "/api/v1/service/@"},
	}
}

// GetManageServiceEndpoints returns endpoints for managing  namespace services
func GetManageServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service"},
		{Method: "DELETE", Route: "/api/v1/service/@"},
	}
}

// GetViewServiceAccountEndpoints returns endpoints for viewing  namespace service accounts
func GetViewServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// GetManageServiceAccountEndpoints returns endpoints for managing  namespace service accounts
func GetManageServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// GetViewNodeEndpoints returns endpoints for viewing  nodes
func GetViewNodeEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/node"},
		{Method: "GET", Route: "/api/v1/node/@"},
	}
}

// GetManageNodeTaintEndpoints returns endpoints for managing  node taints
func GetManageNodeTaintEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "api/v1/node/taint/@"},
		{Method: "POST", Route: "api/v1/node/untaint/@"},
	}
}

// GetDrainNodeEndpoints returns endpoints for draining nodes
func GetDrainNodeEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/node/cordon/@"},
	}
}

// GetViewCustomResourceEndpoints returns endpoints for viewing  custom resources
func GetViewCustomResourceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/custom-resource"},
		{Method: "GET", Route: "/api/v1/custom-resource/@"},
	}
}

// GetManageCustomResourceEndpoints returns endpoints for managing  custom resources
func GetManageCustomResourceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/custom-resource"},
		{Method: "DELETE", Route: "/api/v1/custom-resource/@"},
	}
}

// GetViewLogsEndpoints returns endpoints for viewing  namespace logs
func GetViewLogsEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pod/logs/"},
	}
}

// GetManageEndpointsEndpoints returns endpoints for managing endpoints
func GetManageEndpointsEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/endpoints"},
		{Method: "DELETE", Route: "/api/v1/endpoints/@"},
	}
}

// GetViewEndpointSliceEndpoints returns endpoints for viewing endpoint slices
func GetViewEndpointSliceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/endpoint-slice"},
		{Method: "GET", Route: "/api/v1/endpoint-slice/@"},
	}
}

// GetManageEndpointSliceEndpoints returns endpoints for viewing endpoint slices
func GetManageEndpointSliceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/endpoint-slice"},
		{Method: "DELETE", Route: "/api/v1/endpoint-slice/@"},
	}
}

// GetViewPDBEndpoints returns endpoints for viewing pod disruption budgets
func GetViewPDBEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/PDB"},
		{Method: "GET", Route: "/api/v1/PDB/@"},
	}
}

// GetManagePDBEndpoints returns endpoints for managing pod disruption budgets
func GetManagePDBEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/PDB"},
		{Method: "DELETE", Route: "/api/v1/PDB/@"},
	}
}

// GetViewControllerRevisionEndpoints returns endpoints for viewing controller revisions
func GetViewControllerRevisionEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/controller-revision"},
		{Method: "GET", Route: "/api/v1/controller-revision/@"},
	}
}

// GetManageControllerRevisionEndpoints returns endpoints for managing controller revisions
func GetManageControllerRevisionEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/controller-revision"},
		{Method: "DELETE", Route: "/api/v1/controller-revision/@"},
	}
}

// GetViewReplicationControllerEndpoints returns endpoints for viewing replication controllers
func GetViewReplicationControllerEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/replication-controller"},
		{Method: "GET", Route: "/api/v1/replication-controller/@"},
	}
}

// GetManageReplicationControllerEndpoints returns endpoints for managing replication controllers
func GetManageReplicationControllerEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/replication-controller"},
		{Method: "DELETE", Route: "/api/v1/replication-controller/@"},
	}
}

// GetManageCustomResourceDefinitionEndpoints returns endpoints for managing custom resource definitions
func GetManageCustomResourceDefinitionEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/crd"},
		{Method: "DELETE", Route: "/api/v1/crd/@"},
	}
}

// GetViewCustomResourceDefinitionEndpoints returns endpoints for viewing custom resource definitions
func GetViewCustomResourceDefinitionEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/crd"},
		{Method: "GET", Route: "/api/v1/crd/@"},
	}
}

// GetManageStorageClassEndpoints returns endpoints for managing storage classes
func GetManageStorageClassEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/storage-class"},
		{Method: "DELETE", Route: "/api/v1/storage-class/@"},
	}
}

// GetViewStorageClassEndpoints returns endpoints for viewing storage classes
func GetViewStorageClassEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/storage-class"},
		{Method: "GET", Route: "/api/v1/storage-class/@"},
	}
}

// GetManageClusterRoleBindingEndpoints returns endpoints for managing cluster role bindings
func GetManageClusterRoleBindingEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/cluster-role-binding"},
		{Method: "DELETE", Route: "/api/v1/cluster-role-binding/@"},
	}
}

// GetViewClusterRoleBindingEndpoints returns endpoints for viewing cluster role bindings
func GetViewClusterRoleBindingEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/cluster-role-binding"},
		{Method: "GET", Route: "/api/v1/cluster-role-binding/@"},
	}
}

// GetManageClusterRoleEndpoints returns endpoints for managing cluster roles
func GetManageClusterRoleEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/cluster-role"},
		{Method: "DELETE", Route: "/api/v1/cluster-role/@"},
	}
}

// GetViewClusterRoleEndpoints returns endpoints for viewing cluster roles
func GetViewClusterRoleEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/cluster-role"},
		{Method: "GET", Route: "/api/v1/cluster-role/@"},
	}
}

// GetManagePersistentVolumeEndpoints returns endpoints for managing persistent volumes
func GetManagePersistentVolumeEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/pv"},
		{Method: "DELETE", Route: "/api/v1/pv/@"},
	}
}

// GetViewPersistentVolumeEndpoints returns endpoints for viewing persistent volumes
func GetViewPersistentVolumeEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pv"},
		{Method: "GET", Route: "/api/v1/pv/@"},
	}
}

// GetViewIngressEndpoints returns endpoints for viewing ingresses
func GetViewIngressEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/ingress"},
		{Method: "GET", Route: "/api/v1/ingress/@"},
	}
}

// GetManageIngressEndpoints returns endpoints for managing ingresses
func GetManageIngressEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/ingress"},
		{Method: "DELETE", Route: "/api/v1/ingress/@"},
	}
}

// GetViewCertificateEndpoints returns endpoints for viewing certificates
func GetViewCertificateEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/certificate"},
		{Method: "GET", Route: "/api/v1/certificate/@"},
	}
}

// GetManageCertificateEndpoints returns endpoints for managing certificates
func GetManageCertificateEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/certificate"},
		{Method: "DELETE", Route: "/api/v1/certificate/@"},
	}
}

// GetViewNamespaceRoleEndpoints returns endpoints for viewing roles
func GetViewNamespaceRoleEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/role"},
		{Method: "GET", Route: "/api/v1/role/@"},
	}
}

// GetManageNamespaceRoleEndpoints returns endpoints for managing roles
func GetManageNamespaceRoleEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/role"},
		{Method: "DELETE", Route: "/api/v1/role/@"},
	}
}

// GetViewNamespaceRoleBindingEndpoints returns endpoints for viewing role bindings
func GetViewNamespaceRoleBindingEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/role-binding"},
		{Method: "GET", Route: "/api/v1/role-binding/@"},
	}
}

// GetManageNamespaceRoleBindingEndpoints returns endpoints for managing role bindings
func GetManageNamespaceRoleBindingEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/role-binding"},
		{Method: "DELETE", Route: "/api/v1/role-binding/@"},
	}
}

// GetViewJobEndpoints returns endpoints for viewing jobs
func GetViewJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/job"},
		{Method: "GET", Route: "/api/v1/job/@"},
	}
}

// GetManageJobEndpoints returns endpoints for managing jobs
func GetManageJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/job"},
		{Method: "DELETE", Route: "/api/v1/job/@"},
	}
}

// GetViewCronJobEndpoints returns endpoints for viewing cron jobs
func GetViewCronJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/cronjob"},
		{Method: "GET", Route: "/api/v1/cronjob/@"},
	}
}

// GetManageCronJobEndpoints returns endpoints for managing cron jobs
func GetManageCronJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/cronjob"},
		{Method: "DELETE", Route: "/api/v1/cronjob/@"},
	}
}

// GetViewNetworkPolicyEndpoints returns endpoints for viewing network policies
func GetViewNetworkPolicyEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/network-policy"},
		{Method: "GET", Route: "/api/v1/network-policy/@"},
	}
}

// GetManageNetworkPolicyEndpoints returns endpoints for managing network policies
func GetManageNetworkPolicyEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/network-policy"},
		{Method: "DELETE", Route: "/api/v1/network-policy/@"},
	}
}

// GetViewResourceQuotaEndpoints returns endpoints for viewing resource quotas
func GetViewResourceQuotaEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/resource-quota"},
		{Method: "GET", Route: "/api/v1/resource-quota/@"},
	}
}

// GetManageResourceQuotaEndpoints returns endpoints for managing resource quotas
func GetManageResourceQuotaEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/resource-quota"},
		{Method: "DELETE", Route: "/api/v1/resource-quota/@"},
	}
}

// GetViewGatewayEndpoints returns endpoints for viewing gateways
func GetViewGatewayEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/gateway"},
		{Method: "GET", Route: "/api/v1/gateway/@"},
	}
}

// GetManageGatewayEndpoints returns endpoints for managing gateways
func GetManageGatewayEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/gateway"},
		{Method: "DELETE", Route: "/api/v1/gateway/@"},
	}
}

// GetViewVirtualServiceEndpoints returns endpoints for viewing virtual services
func GetViewVirtualServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/virtual-service"},
		{Method: "GET", Route: "/api/v1/virtual-service/@"},
	}
}

// GetManageVirtualServiceEndpoints returns endpoints for managing virtual services
func GetManageVirtualServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/virtual-service"},
		{Method: "DELETE", Route: "/api/v1/virtual-service/@"},
	}
}

// GetViewPersistentVolumeClaimEndpoints returns endpoints for viewing PVCs
func GetViewPersistentVolumeClaimEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pvc"},
		{Method: "GET", Route: "/api/v1/pvc/@"},
	}
}

// GetManagePersistentVolumeClaimEndpoints returns endpoints for managing PVCs
func GetManagePersistentVolumeClaimEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/pvc"},
		{Method: "DELETE", Route: "/api/v1/pvc/@"},
	}
}
