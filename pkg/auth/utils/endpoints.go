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

// GetUpdateNamespaceEndpoint returns endpoints for updating K8s namespace operations
func GetUpdateNamespaceEndpoint() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/namespace"},
	}
}

// GetDeleteNamespaceEndpoint returns endpoints for deleting K8s namespace operations
func GetDeleteNamespaceEndpoint() []models.Endpoint {
	return []models.Endpoint{
		{Method: "DELETE", Route: "/api/v1/namespace/@"},
	}
}

// ViewDeploymentEndpoints returns a slice of endpoints for viewing deployments
func ViewDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/deployment"},
		{Method: "GET", Route: "/api/v1/deployment/stats"},
		{Method: "GET", Route: "/api/v1/deployment/@"},
		{Method: "GET", Route: "/api/v1/deployment/@/pods"},
	}
}

// ManageDeploymentEndpoints returns a slice of endpoints for managing deployments
func ManageDeploymentEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/deployment"},
		{Method: "DELETE", Route: "/api/v1/deployment/@"},
	}
}

// ViewPodEndpoints returns a slice of endpoints for viewing pods
func ViewPodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/pod"},
		{Method: "GET", Route: "/api/v1/pod/@"},
		{Method: "GET", Route: "/api/v1/pod/stats"},
	}
}

// ManagePodEndpoints returns a slice of endpoints for managing pods
func ManagePodEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/pod"},
		{Method: "DELETE", Route: "/api/v1/pod/@"},
	}
}

// ViewReplicaSetEndpoints returns a slice of endpoints for viewing replica sets
func ViewReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/replicaset"},
		{Method: "GET", Route: "/api/v1/replicaset/@"},
	}
}

// ManageReplicaSetEndpoints returns a slice of endpoints for managing replica sets
func ManageReplicaSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/replicaset"},
		{Method: "DELETE", Route: "/api/v1/replicaset/@"},
	}
}

// ViewStatefulSetEndpoints returns a slice of endpoints for viewing stateful sets
func ViewStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/statefulset"},
		{Method: "GET", Route: "/api/v1/statefulset/@"},
		{Method: "GET", Route: "/api/v1/statefulset/@/pods"},
		{Method: "GET", Route: "/api/v1/statefulset/stats"},
	}
}

// ManageStatefulSetEndpoints returns a slice of endpoints for managing stateful sets
func ManageStatefulSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/statefulset"},
		{Method: "DELETE", Route: "/api/v1/statefulset/@"},
	}
}

// ViewDaemonSetEndpoints returns a slice of endpoints for viewing daemon sets
func ViewDaemonSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/daemonset"},
		{Method: "GET", Route: "/api/v1/daemonset/stats"},
		{Method: "GET", Route: "/api/v1/daemonset/@"},
	}
}

// ManageDaemonSetEndpoints returns a slice of endpoints for managing daemon sets
func ManageDaemonSetEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/daemonset"},
		{Method: "DELETE", Route: "/api/v1/daemonset/@"},
	}
}

// ViewSecretEndpoints returns a slice of endpoints for viewing secrets
func ViewSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/secret"},
		{Method: "GET", Route: "/api/v1/secret/@"},
	}
}

// ManageSecretEndpoints returns a slice of endpoints for managing secrets
func ManageSecretEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/secret"},
		{Method: "DELETE", Route: "/api/v1/secret/@"},
	}
}

// ViewConfigMapEndpoints returns a slice of endpoints for viewing config maps
func ViewConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/config-map"},
		{Method: "GET", Route: "/api/v1/config-map/@"},
	}
}

// ManageConfigMapEndpoints returns a slice of endpoints for managing config maps
func ManageConfigMapEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/config-map"},
		{Method: "DELETE", Route: "/api/v1/config-map/@"},
	}
}

// ViewServiceAccountEndpoints returns a slice of endpoints for viewing service accounts
func ViewServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// ManageServiceAccountEndpoints returns a slice of endpoints for managing service accounts
func ManageServiceAccountEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service-account"},
		{Method: "GET", Route: "/api/v1/service-account/@"},
	}
}

// ViewServiceEndpoints returns a slice of endpoints for viewing services
func ViewServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "GET", Route: "/api/v1/service"},
		{Method: "GET", Route: "/api/v1/service/@"},
	}
}

// ManageServiceEndpoints returns a slice of endpoints for managing services
func ManageServiceEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/service"},
		{Method: "DELETE", Route: "/api/v1/service/@"},
	}
}

// ManageIngressEndpoints returns a slice of endpoints for managing ingress
func ManageIngressEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/ingress"},
		{Method: "DELETE", Route: "/api/v1/ingress/@"},
	}
}

// ManageCertificateEndpoints returns a slice of endpoints for managing certificates
func ManageCertificateEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/certificate"},
		{Method: "DELETE", Route: "/api/v1/certificate/@"},
	}
}

// ManageRoleEndpoints returns a slice of endpoints for managing roles
func ManageRoleEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/role"},
		{Method: "DELETE", Route: "/api/v1/role/@"},
	}
}

// ManageRoleBindingEndpoints returns a slice of endpoints for managing role bindings
func ManageRoleBindingEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/role-binding"},
		{Method: "DELETE", Route: "/api/v1/role-binding/@"},
	}
}

// ManageJobEndpoints returns a slice of endpoints for managing jobs
func ManageJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/job"},
		{Method: "DELETE", Route: "/api/v1/job/@"},
	}
}

// ManageCronJobEndpoints returns a slice of endpoints for managing cron jobs
func ManageCronJobEndpoints() []models.Endpoint {
	return []models.Endpoint{
		{Method: "POST", Route: "/api/v1/cronjob"},
		{Method: "DELETE", Route: "/api/v1/cronjob/@"},
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
