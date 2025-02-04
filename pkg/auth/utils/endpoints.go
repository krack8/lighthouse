package utils

import "github.com/krack8/lighthouse/pkg/auth/models"

// GetDefaultEndpoints returns endpoints for creating K8s namespaces
func GetDefaultEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/clusters"},
	}
}

// GetCreateNamespaceEndpoints returns endpoints for creating K8s namespaces
func GetCreateNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
	}
}

// GetViewNamespaceEndpoints returns endpoints for viewing K8s namespaces
func GetViewNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace/names"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
		{Method: "GET", Route: "/api/v1/event"},
	}
}

// GetK8sNamespaceCreateEndpoints returns endpoints for creating K8s namespaces
func GetK8sNamespaceCreateEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
	}
}

// GetViewK8sNamespaceEndpoints returns endpoints for viewing K8s namespaces
func GetViewK8sNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/namespace"},
		{Method: "GET", Route: "api/v1/namespace/names"},
		{Method: "GET", Route: "/api/v1/namespace/@"},
		{Method: "GET", Route: "/api/v1/event"},
	}
}

// GetK8sNamespaceUpdateEndpoints returns endpoints for updating K8s namespaces
func GetK8sNamespaceUpdateEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
	}
}

// GetK8sNamespaceDeleteEndpoints returns endpoints for deleting K8s namespaces
func GetK8sNamespaceDeleteEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "DELETE", Route: "/api/v1/namespace/@"},
	}
}

// GetViewK8sNamespaceDeploymentEndpoints returns endpoints for viewing K8s namespace deployments
func GetViewK8sNamespaceDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/deployment"},
		{Method: "GET", Route: "/api/v1/deployment/stats"},
		{Method: "GET", Route: "/api/v1/deployment/@"},
		{Method: "GET", Route: "/api/v1/deployment/@/pods"},
	}
}

// GetK8sNamespaceManageDeploymentEndpoints returns endpoints for managing K8s namespace deployments
func GetK8sNamespaceManageDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/deployment"},
		{Method: "DELETE", Route: "/api/v1/deployment/@"},
	}
}

// GetViewK8sNamespacePodEndpoints returns endpoints for viewing K8s namespace pods
func GetViewK8sNamespacePodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pod"},
		{Method: "GET", Route: "/api/v1/pod/@"},
		{Method: "GET", Route: "/api/v1/pod/stats"},
	}
}

// GetK8sNamespaceManagePodEndpoints returns endpoints for managing K8s namespace pods
func GetK8sNamespaceManagePodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/pod"},
		{Method: "DELETE", Route: "/api/v1/pod/@"},
	}
}

// GetViewK8sNamespaceReplicaSetEndpoints returns endpoints for viewing K8s namespace replica sets
func GetViewK8sNamespaceReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/replicaset"},
		{Method: "GET", Route: "/api/v1/replicaset/@"},
	}
}

// GetK8sNamespaceManageReplicaSetEndpoints returns endpoints for managing K8s namespace replica sets
func GetK8sNamespaceManageReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/replicaset"},
		{Method: "DELETE", Route: "/api/v1/replicaset/@"},
	}
}

// GetViewK8sNamespaceStatefulSetEndpoints returns endpoints for viewing K8s namespace stateful sets
func GetViewK8sNamespaceStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/statefulset"},
		{Method: "GET", Route: "/api/v1/statefulset/@"},
		{Method: "GET", Route: "/api/v1/statefulset/@/pods"},
		{Method: "GET", Route: "/api/v1/statefulset/stats"},
	}
}

// GetK8sNamespaceManageStatefulSetEndpoints returns endpoints for managing K8s namespace stateful sets
func GetK8sNamespaceManageStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/statefulset"},
		{Method: "DELETE", Route: "/api/v1/statefulset/@"},
	}
}

// GetViewK8sNamespaceSecretEndpoints returns endpoints for viewing K8s namespace secrets
func GetViewK8sNamespaceSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/secret"},
		{Method: "GET", Route: "/api/v1/secret/@"},
	}
}

// GetK8sNamespaceManageSecretEndpoints returns endpoints for managing K8s namespace secrets
func GetK8sNamespaceManageSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/secret"},
		{Method: "DELETE", Route: "/api/v1/secret/@"},
	}
}

// GetViewK8sNamespaceConfigMapEndpoints returns endpoints for viewing K8s namespace config maps
func GetViewK8sNamespaceConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/config-map"},
		{Method: "GET", Route: "/api/v1/config-map/@"},
	}
}

// GetK8sNamespaceManageConfigMapEndpoints returns endpoints for managing K8s namespace config maps
func GetK8sNamespaceManageConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/config-map"},
		{Method: "DELETE", Route: "/api/v1/config-map/@"},
	}
}

// GetViewK8sNamespaceServiceEndpoints returns endpoints for viewing K8s namespace services
func GetViewK8sNamespaceServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service"},
		{Method: "GET", Route: "/api/v1/service/@"},
	}
}

// GetK8sNamespaceManageServiceEndpoints returns endpoints for managing K8s namespace services
func GetK8sNamespaceManageServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service"},
		{Method: "DELETE", Route: "/api/v1/service/@"},
	}
}

// GetViewK8sNamespaceServiceAccountEndpoints returns endpoints for viewing K8s namespace service accounts
func GetViewK8sNamespaceServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// GetK8sNamespaceManageServiceAccountEndpoints returns endpoints for managing K8s namespace service accounts
func GetK8sNamespaceManageServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// GetViewK8sNodeEndpoints returns endpoints for viewing K8s nodes
func GetViewK8sNodeEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/node"},
		{Method: "GET", Route: "/api/v1/node/@"},
	}
}

// GetManageK8sNodeTaintEndpoints returns endpoints for managing K8s node taints
func GetManageK8sNodeTaintEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "api/v1/node/taint/@"},
		{Method: "POST", Route: "api/v1/node/untaint/@"},
	}
}

// GetViewK8sCustomResourceEndpoints returns endpoints for viewing K8s custom resources
func GetViewK8sCustomResourceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/custom-resource"},
		{Method: "GET", Route: "/api/v1/custom-resource/@"},
	}
}

// GetManageK8sCustomResourceEndpoints returns endpoints for managing K8s custom resources
func GetManageK8sCustomResourceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/custom-resource"},
		{Method: "DELETE", Route: "/api/v1/custom-resource/@"},
	}
}

// GetViewK8sNamespaceLogsEndpoints returns endpoints for viewing K8s namespace logs
func GetViewK8sNamespaceLogsEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pod/logs/"},
	}
}

// GetManageNamespaceEndpoints returns endpoints for managing K8s namespace endpoints
func GetManageNamespaceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "api/v1/endpoints"},
		{Method: "DELETE", Route: "api/v1/endpoints/@"},
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

// ViewNamespacePDB returns endpoints for viewing PDBs
func ViewNamespacePDB() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/PDB"},
		{Method: "GET", Route: "/api/v1/PDB/@"},
	}
}

// ManageNamespacePDB returns endpoints for managing PDBs
func ManageNamespacePDB() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/PDB"},
		{Method: "DELETE", Route: "/api/v1/PDB/@"},
	}
}

// ViewNamespaceControllerRevision returns endpoints for viewing controller revisions
func ViewNamespaceControllerRevision() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/controller-revision"},
		{Method: "GET", Route: "/api/v1/controller-revision/@"},
	}
}

// ManageNamespaceControllerRevision returns endpoints for managing controller revisions
func ManageNamespaceControllerRevision() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/controller-revision"},
		{Method: "DELETE", Route: "/api/v1/controller-revision/@"},
	}
}

// ViewNamespaceReplicationController returns endpoints for viewing replication controllers
func ViewNamespaceReplicationController() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/replication-controller"},
		{Method: "GET", Route: "/api/v1/replication-controller/@"},
	}
}

// ManageNamespaceReplicationController returns endpoints for managing replication controllers
func ManageNamespaceReplicationController() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/replication-controller"},
		{Method: "DELETE", Route: "/api/v1/replication-controller/@"},
	}
}
