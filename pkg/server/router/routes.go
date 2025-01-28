package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
	"github.com/krack8/lighthouse/pkg/controller/api"
)

var userController *controllers.UserController

// Declare the userService as a global variable
var rbacController *controllers.RbacController

func AddApiRoutes(httpRg *gin.RouterGroup) {

	// Define the login route separately without middleware
	// Login route
	httpRg.POST("/auth/login", controllers.LoginHandler)

	// Refresh token route
	httpRg.POST("/auth/refresh-token", controllers.RefreshTokenHandler)

	// Apply the AuthMiddleware to the /api/v1 routes
	apiGroup := httpRg.Group("/api/v1", middleware.AuthMiddleware())

	// User routes
	apiGroup.POST("/users", userController.CreateUserHandler)
	apiGroup.GET("/users", userController.GetAllUsersHandler)
	apiGroup.GET("/users/:id", userController.GetUserHandler)
	apiGroup.PUT("/users/:id", userController.UpdateUserHandler)
	apiGroup.DELETE("/users/:id", userController.DeleteUserHandler)

	// RBAC routes
	apiGroup.POST("/rbac/permissions", rbacController.CreatePermissionHandler)
	apiGroup.POST("/rbac/roles", rbacController.CreateRoleHandler)
	apiGroup.POST("/rbac/assign-roles", rbacController.AssignRolesHandler)

	// Namespace
	httpRg.GET("api/v1/namespace", api.NamespaceController().GetNamespaceList)
	httpRg.GET("api/v1/namespace/names", api.NamespaceController().GetNamespaceNameList)
	httpRg.GET("api/v1/namespace/:name", api.NamespaceController().GetNamespaceDetails)
	httpRg.POST("api/v1/namespace", api.NamespaceController().DeployNamespace)
	httpRg.DELETE("api/v1/namespace/:name", api.NamespaceController().DeleteNamespace)
	// Certificate
	httpRg.GET("api/v1/certificate", api.CertificateController().GetCertificateList)
	httpRg.GET("api/v1/certificate/:name", api.CertificateController().GetCertificateDetails)
	httpRg.POST("api/v1/certificate", api.CertificateController().DeployCertificate)
	httpRg.DELETE("api/v1/certificate/:name", api.CertificateController().DeleteCertificate)
	// Config Map
	httpRg.GET("api/v1/config-map", api.ConfigMapController().GetConfigMapList)
	httpRg.GET("api/v1/config-map/:name", api.ConfigMapController().GetConfigMapDetails)
	httpRg.POST("api/v1/config-map", api.ConfigMapController().DeployConfigMap)
	httpRg.DELETE("api/v1/config-map/:name", api.ConfigMapController().DeleteConfigMap)
	// ClusterRole
	httpRg.GET("api/v1/cluster-role", api.ClusterRoleController().GetClusterRoleList)
	httpRg.GET("api/v1/cluster-role/:name", api.ClusterRoleController().GetClusterRoleDetails)
	httpRg.POST("api/v1/cluster-role", api.ClusterRoleController().DeployClusterRole)
	httpRg.DELETE("api/v1/cluster-role/:name", api.ClusterRoleController().DeleteClusterRole)
	// ClusterRoleBinding
	httpRg.GET("api/v1/cluster-role-binding", api.ClusterRoleBindingController().GetClusterRoleBindingList)
	httpRg.GET("api/v1/cluster-role-binding/:name", api.ClusterRoleBindingController().GetClusterRoleBindingDetails)
	httpRg.POST("api/v1/cluster-role-binding", api.ClusterRoleBindingController().DeployClusterRoleBinding)
	httpRg.DELETE("api/v1/cluster-role-binding/:name", api.ClusterRoleBindingController().DeleteClusterRoleBinding)
	// ControllerRevision
	httpRg.GET("api/v1/controller-revision", api.ControllerRevisionController().GetControllerRevisionList)
	httpRg.GET("api/v1/controller-revision/:name", api.ControllerRevisionController().GetControllerRevisionDetails)
	httpRg.POST("api/v1/controller-revision", api.ControllerRevisionController().DeployControllerRevision)
	httpRg.DELETE("api/v1/controller-revision/:name", api.ControllerRevisionController().DeleteControllerRevision)
	// CRD
	httpRg.GET("api/v1/crd", api.CrdController().GetCrdList)
	httpRg.GET("api/v1/crd/:name", api.CrdController().GetCrdDetails)
	httpRg.POST("api/v1/crd", api.CrdController().DeployCrd)
	httpRg.DELETE("api/v1/crd/:name", api.CrdController().DeleteCrd)
	// Custom Resource
	httpRg.GET("api/v1/custom-resource", api.CustomResourceController().GetCustomResourceList)
	httpRg.GET("api/v1/custom-resource/:name", api.CustomResourceController().GetCustomResourceDetails)
	httpRg.POST("api/v1/custom-resource", api.CustomResourceController().DeployCustomResource)
	httpRg.DELETE("api/v1/custom-resource/:name", api.CustomResourceController().DeleteCustomResource)
	//Cronjob
	httpRg.GET("api/v1/cronjob", api.CronJobController().GetCronJobList)
	httpRg.GET("api/v1/cronjob/:name", api.CronJobController().GetCronJobDetails)
	httpRg.POST("api/v1/cronjob", api.CronJobController().DeployCronJob)
	httpRg.DELETE("api/v1/cronjob/:name", api.CronJobController().DeleteCronJob)
	// Daemonset
	httpRg.GET("api/v1/daemonset", api.DaemonSetController().GetDaemonSetList)
	httpRg.GET("api/v1/daemonset/:name", api.DaemonSetController().GetDaemonSetDetails)
	httpRg.POST("api/v1/daemonset", api.DaemonSetController().DeployDaemonSet)
	httpRg.DELETE("api/v1/daemonset/:name", api.DaemonSetController().DeleteDaemonSet)
	httpRg.GET("api/v1/daemonset/stats", api.DaemonSetController().GetDaemonSetStats)
	// Deployment
	httpRg.GET("api/v1/deployment", api.DeploymentController().GetDeploymentList)
	httpRg.GET("api/v1/deployment/:name", api.DeploymentController().GetDeploymentDetails)
	httpRg.POST("api/v1/deployment", api.DeploymentController().DeployDeployment)
	httpRg.DELETE("api/v1/deployment/:name", api.DeploymentController().DeleteDeployment)
	httpRg.GET("api/v1/deployment/stats", api.DeploymentController().GetDeploymentStats)
	httpRg.GET("api/v1/deployment/:name/pods", api.DeploymentController().GetDeploymentPodList)
	// Endpoints
	httpRg.GET("api/v1/endpoints", api.EndpointsController().GetEndpointsList)
	httpRg.GET("api/v1/endpoints/:name", api.EndpointsController().GetEndpointsDetails)
	httpRg.POST("api/v1/endpoints", api.EndpointsController().DeployEndpoints)
	httpRg.DELETE("api/v1/endpoints/:name", api.EndpointsController().DeleteEndpoints)
	// EndpointSlice
	httpRg.GET("api/v1/endpoint-slice", api.EndpointSliceController().GetEndpointSliceList)
	httpRg.GET("api/v1/endpoint-slice/:name", api.EndpointSliceController().GetEndpointSliceDetails)
	httpRg.POST("api/v1/endpoint-slice", api.EndpointSliceController().DeployEndpointSlice)
	httpRg.DELETE("api/v1/endpoint-slice/:name", api.EndpointSliceController().DeleteEndpointSlice)
	// event
	httpRg.GET("api/v1/event", api.EventController().GetEventList)
	httpRg.GET("api/v1/event/:name", api.EventController().GetEventDetails)
	// HPA
	httpRg.GET("api/v1/hpa", api.HpaController().GetHpaList)
	httpRg.GET("api/v1/hpa/:name", api.HpaController().GetHpaDetails)
	// Ingress
	httpRg.GET("api/v1/ingress", api.IngressController().GetIngressList)
	httpRg.GET("api/v1/ingress/:name", api.IngressController().GetIngressDetails)
	httpRg.POST("api/v1/ingress", api.IngressController().DeployIngress)
	httpRg.DELETE("api/v1/ingress/:name", api.IngressController().DeleteIngress)
	// Istio Gateway
	httpRg.GET("api/v1/gateway", api.IstioGatewayController().GetIstioGatewayList)
	httpRg.GET("api/v1/gateway/:name", api.IstioGatewayController().GetIstioGatewayDetails)
	httpRg.POST("api/v1/gateway", api.IstioGatewayController().DeployIstioGateway)
	httpRg.DELETE("api/v1/gateway/:name", api.IstioGatewayController().DeleteIstioGateway)
	//Job
	httpRg.GET("api/v1/job", api.JobController().GetJobList)
	httpRg.GET("api/v1/job/:name", api.JobController().GetJobDetails)
	httpRg.POST("api/v1/job", api.JobController().DeployJob)
	httpRg.DELETE("api/v1/job/:name", api.JobController().DeleteJob)
	//Load Balancer
	httpRg.GET("api/v1/load-balancer", api.LoadBalancerController().GetLoadBalancerList)
	httpRg.GET("api/v1/load-balancer/:name", api.LoadBalancerController().GetLoadBalancerDetails)
	// Manifest
	httpRg.POST("api/v1/manifest", api.ManifestController().DeployManifest)
	// Network Policy
	httpRg.GET("api/v1/network-policy", api.NetworkPolicyController().GetNetworkPolicyList)
	httpRg.GET("api/v1/network-policy/:name", api.NetworkPolicyController().GetNetworkPolicyDetails)
	//Node
	httpRg.GET("api/v1/node", api.NodeController().GetNodeList)
	httpRg.GET("api/v1/node/:name", api.NodeController().GetNodeDetails)
	httpRg.GET("api/v1/node/cordon/:name", api.NodeController().NodeCordon)
	httpRg.POST("api/v1/node/taint/:name", api.NodeController().NodeTaint)
	httpRg.POST("api/v1/node/untaint/:name", api.NodeController().NodeUnTaint)
	// Pod
	httpRg.GET("api/v1/pod", api.PodController().GetPodList)
	httpRg.GET("api/v1/pod/:name", api.PodController().GetPodDetails)
	httpRg.GET("api/v1/pod/logs/:name", api.PodController().GetPodLogs)
	httpRg.POST("api/v1/pod", api.PodController().DeployPod)
	httpRg.DELETE("api/v1/pod/:name", api.PodController().DeletePod)
	httpRg.GET("api/v1/pod/stats", api.PodController().GetPodStats)
	// PodDisruptionBudgets
	httpRg.GET("api/v1/PDB", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsList)
	httpRg.GET("api/v1/PDB/:name", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsDetails)
	httpRg.POST("api/v1/PDB", api.PodDisruptionBudgetsController().DeployPodDisruptionBudgets)
	httpRg.DELETE("api/v1/PDB/:name", api.PodDisruptionBudgetsController().DeletePodDisruptionBudgets)
	// Pod Metrics
	httpRg.GET("api/v1/pod-metrics", api.PodMetricsController().GetPodMetricsList)
	httpRg.GET("api/v1/pod-metrics/:pod", api.PodMetricsController().GetPodMetricsDetails)
	// PV
	httpRg.GET("api/v1/pv", api.PvController().GetPvList)
	httpRg.GET("api/v1/pv/:name", api.PvController().GetPvDetails)
	httpRg.POST("api/v1/pv", api.PvController().DeployPv)
	httpRg.DELETE("api/v1/pv/:name", api.PvController().DeletePv)
	// Persistent Volume Claim
	httpRg.GET("api/v1/pvc", api.PvcController().GetPvcList)
	httpRg.GET("api/v1/pvc/:name", api.PvcController().GetPvcDetails)
	httpRg.POST("api/v1/pvc", api.PvcController().DeployPvc)
	httpRg.DELETE("api/v1/pvc/:name", api.PvcController().DeletePvc)
	// ReplicaSet
	httpRg.GET("api/v1/replicaset", api.ReplicaSetController().GetReplicaSetList)
	httpRg.GET("api/v1/replicaset/:name", api.ReplicaSetController().GetReplicaSetDetails)
	httpRg.GET("api/v1/replicaset/stats", api.ReplicaSetController().GetReplicaSetStats)
	httpRg.POST("api/v1/replicaset", api.ReplicaSetController().DeployReplicaSet)
	httpRg.DELETE("api/v1/replicaset/:name", api.ReplicaSetController().DeleteReplicaSet)
	// ReplicationController
	httpRg.GET("api/v1/replication-controller", api.ReplicationControllerController().GetReplicationControllerList)
	httpRg.GET("api/v1/replication-controller/:name", api.ReplicationControllerController().GetReplicationControllerDetails)
	httpRg.POST("api/v1/replication-controller", api.ReplicationControllerController().DeployReplicationController)
	httpRg.DELETE("api/v1/replication-controller/:name", api.ReplicationControllerController().DeleteReplicationController)
	// Resource Quota
	httpRg.GET("api/v1/resource-quota", api.ResourceQuotaController().GetResourceQuotaList)
	httpRg.GET("api/v1/resource-quota/:name", api.ResourceQuotaController().GetResourceQuotaDetails)
	httpRg.POST("api/v1/resource-quota", api.ResourceQuotaController().DeployResourceQuota)
	httpRg.DELETE("api/v1/resource-quota/:name", api.ResourceQuotaController().DeleteResourceQuota)
	// Role
	httpRg.GET("api/v1/role", api.RoleController().GetRoleList)
	httpRg.GET("api/v1/role/:name", api.RoleController().GetRoleDetails)
	httpRg.POST("api/v1/role", api.RoleController().DeployRole)
	httpRg.DELETE("api/v1/role/:name", api.RoleController().DeleteRole)
	// RoleBinding
	httpRg.GET("api/v1/role-binding", api.RoleBindingController().GetRoleBindingList)
	httpRg.GET("api/v1/role-binding/:name", api.RoleBindingController().GetRoleBindingDetails)
	httpRg.POST("api/v1/role-binding", api.RoleBindingController().DeployRoleBinding)
	httpRg.DELETE("api/v1/role-binding/:name", api.RoleBindingController().DeleteRoleBinding)
	// Service Account
	httpRg.GET("api/v1/service-account", api.ServiceAccountController().GetServiceAccountList)
	httpRg.GET("api/v1/service-account/:name", api.ServiceAccountController().GetServiceAccountDetails)
	httpRg.POST("api/v1/service-account", api.ServiceAccountController().DeployServiceAccount)
	httpRg.DELETE("api/v1/service-account/:name", api.ServiceAccountController().DeleteServiceAccount)
	// Secret
	httpRg.GET("api/v1/secret", api.SecretController().GetSecretList)
	httpRg.GET("api/v1/secret/:name", api.SecretController().GetSecretDetails)
	httpRg.POST("api/v1/secret", api.SecretController().DeploySecret)
	httpRg.DELETE("api/v1/secret/:name", api.SecretController().DeleteSecret)
	// StatefulSet
	httpRg.GET("api/v1/statefulset", api.StatefulSetController().GetStatefulSetList)
	httpRg.GET("api/v1/statefulset/:name", api.StatefulSetController().GetStatefulSetDetails)
	httpRg.POST("api/v1/statefulset", api.StatefulSetController().DeployStatefulSet)
	httpRg.DELETE("api/v1/statefulset/:name", api.StatefulSetController().DeleteStatefulSet)
	httpRg.GET("api/v1/statefulset/stats", api.StatefulSetController().GetStatefulSetStats)
	httpRg.GET("api/v1/statefulset/:name/pods", api.StatefulSetController().GetStatefulSetPodList)
	// Storage class
	httpRg.GET("api/v1/storage-class", api.StorageClassController().GetStorageClassList)
	httpRg.GET("api/v1/storage-class/:name", api.StorageClassController().GetStorageClassDetails)
	httpRg.POST("api/v1/storage-class", api.StorageClassController().DeployStorageClass)
	httpRg.DELETE("api/v1/storage-class/:name", api.StorageClassController().DeleteStorageClass)
	// Service
	httpRg.GET("api/v1/service", api.SvcController().GetSvcList)
	httpRg.GET("api/v1/service/:name", api.SvcController().GetSvcDetails)
	httpRg.POST("api/v1/service", api.SvcController().DeploySVC)
	httpRg.DELETE("api/v1/service/:name", api.SvcController().DeleteSvc)
	// Virtual Service
	httpRg.GET("api/v1/virtual-service", api.VirtualServiceController().GetVirtualServiceList)
	httpRg.GET("api/v1/virtual-service/:name", api.VirtualServiceController().GetVirtualServiceDetails)
	httpRg.POST("api/v1/virtual-service", api.VirtualServiceController().DeployVirtualService)
	httpRg.DELETE("api/v1/virtual-service/:name", api.VirtualServiceController().DeleteVirtualService)
	// Volume Snapshot
	httpRg.GET("api/v1/volume-snapshot", api.VolumeSnapshotController().GetVolumeSnapshotList)
	httpRg.GET("api/v1/volume-snapshot/:name", api.VolumeSnapshotController().GetVolumeSnapshotDetails)
	httpRg.POST("api/v1/volume-snapshot", api.VolumeSnapshotController().DeployVolumeSnapshot)
	httpRg.DELETE("api/v1/volume-snapshot/:name", api.VolumeSnapshotController().DeleteVolumeSnapshot)
	// Volume Snapshot Content
	httpRg.GET("api/v1/volume-snapshot-content", api.VolumeSnapshotContentController().GetVolumeSnapshotContentList)
	httpRg.GET("api/v1/volume-snapshot-content/:name", api.VolumeSnapshotContentController().GetVolumeSnapshotContentDetails)
	// Volume Snapshot Class
	httpRg.GET("api/v1/volume-snapshot-class", api.VolumeSnapshotClassController().GetVolumeSnapshotClassList)
	httpRg.GET("api/v1/volume-snapshot-class/:name", api.VolumeSnapshotClassController().GetVolumeSnapshotClassDetails)
}
