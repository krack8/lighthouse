package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/authApi"
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
	apiGroup.POST("/users", authApi.UserController.CreateUserHandler)
	apiGroup.GET("/users", authApi.UserController.GetAllUsersHandler)
	apiGroup.GET("/users/:id", authApi.UserController.GetUserHandler)
	apiGroup.PUT("/users/:id", authApi.UserController.UpdateUserHandler)
	apiGroup.DELETE("/users/:id", authApi.UserController.DeleteUserHandler)

	// RBAC routes
	apiGroup.POST("/rbac/permissions", authApi.RbacController.CreatePermissionHandler)
	apiGroup.POST("/rbac/roles", authApi.RbacController.CreateRoleHandler)
	apiGroup.POST("/rbac/assign-roles", authApi.RbacController.AssignRolesHandler)

	// Namespace
	apiGroup.GET("api/v1/namespace", api.NamespaceController().GetNamespaceList)
	apiGroup.GET("api/v1/namespace/names", api.NamespaceController().GetNamespaceNameList)
	apiGroup.GET("api/v1/namespace/:name", api.NamespaceController().GetNamespaceDetails)
	apiGroup.POST("api/v1/namespace", api.NamespaceController().DeployNamespace)
	apiGroup.DELETE("api/v1/namespace/:name", api.NamespaceController().DeleteNamespace)

	// Certificate
	apiGroup.GET("api/v1/certificate", api.CertificateController().GetCertificateList)
	apiGroup.GET("api/v1/certificate/:name", api.CertificateController().GetCertificateDetails)
	apiGroup.POST("api/v1/certificate", api.CertificateController().DeployCertificate)
	apiGroup.DELETE("api/v1/certificate/:name", api.CertificateController().DeleteCertificate)

	// Config Map
	apiGroup.GET("api/v1/config-map", api.ConfigMapController().GetConfigMapList)
	apiGroup.GET("api/v1/config-map/:name", api.ConfigMapController().GetConfigMapDetails)
	apiGroup.POST("api/v1/config-map", api.ConfigMapController().DeployConfigMap)
	apiGroup.DELETE("api/v1/config-map/:name", api.ConfigMapController().DeleteConfigMap)

	// ClusterRole
	apiGroup.GET("api/v1/cluster-role", api.ClusterRoleController().GetClusterRoleList)
	apiGroup.GET("api/v1/cluster-role/:name", api.ClusterRoleController().GetClusterRoleDetails)
	apiGroup.POST("api/v1/cluster-role", api.ClusterRoleController().DeployClusterRole)
	apiGroup.DELETE("api/v1/cluster-role/:name", api.ClusterRoleController().DeleteClusterRole)

	// ClusterRoleBinding
	apiGroup.GET("api/v1/cluster-role-binding", api.ClusterRoleBindingController().GetClusterRoleBindingList)
	apiGroup.GET("api/v1/cluster-role-binding/:name", api.ClusterRoleBindingController().GetClusterRoleBindingDetails)
	apiGroup.POST("api/v1/cluster-role-binding", api.ClusterRoleBindingController().DeployClusterRoleBinding)
	apiGroup.DELETE("api/v1/cluster-role-binding/:name", api.ClusterRoleBindingController().DeleteClusterRoleBinding)

	// ControllerRevision
	apiGroup.GET("api/v1/controller-revision", api.ControllerRevisionController().GetControllerRevisionList)
	apiGroup.GET("api/v1/controller-revision/:name", api.ControllerRevisionController().GetControllerRevisionDetails)
	apiGroup.POST("api/v1/controller-revision", api.ControllerRevisionController().DeployControllerRevision)
	apiGroup.DELETE("api/v1/controller-revision/:name", api.ControllerRevisionController().DeleteControllerRevision)

	// CRD
	apiGroup.GET("api/v1/crd", api.CrdController().GetCrdList)
	apiGroup.GET("api/v1/crd/:name", api.CrdController().GetCrdDetails)
	apiGroup.POST("api/v1/crd", api.CrdController().DeployCrd)
	apiGroup.DELETE("api/v1/crd/:name", api.CrdController().DeleteCrd)

	// Custom Resource
	apiGroup.GET("api/v1/custom-resource", api.CustomResourceController().GetCustomResourceList)
	apiGroup.GET("api/v1/custom-resource/:name", api.CustomResourceController().GetCustomResourceDetails)
	apiGroup.POST("api/v1/custom-resource", api.CustomResourceController().DeployCustomResource)
	apiGroup.DELETE("api/v1/custom-resource/:name", api.CustomResourceController().DeleteCustomResource)

	//Cronjob
	apiGroup.GET("api/v1/cronjob", api.CronJobController().GetCronJobList)
	apiGroup.GET("api/v1/cronjob/:name", api.CronJobController().GetCronJobDetails)
	apiGroup.POST("api/v1/cronjob", api.CronJobController().DeployCronJob)
	apiGroup.DELETE("api/v1/cronjob/:name", api.CronJobController().DeleteCronJob)

	// Daemonset
	apiGroup.GET("api/v1/daemonset", api.DaemonSetController().GetDaemonSetList)
	apiGroup.GET("api/v1/daemonset/:name", api.DaemonSetController().GetDaemonSetDetails)
	apiGroup.POST("api/v1/daemonset", api.DaemonSetController().DeployDaemonSet)
	apiGroup.DELETE("api/v1/daemonset/:name", api.DaemonSetController().DeleteDaemonSet)
	apiGroup.GET("api/v1/daemonset/stats", api.DaemonSetController().GetDaemonSetStats)

	// Deployment
	apiGroup.GET("api/v1/deployment", api.DeploymentController().GetDeploymentList)
	apiGroup.GET("api/v1/deployment/:name", api.DeploymentController().GetDeploymentDetails)
	apiGroup.POST("api/v1/deployment", api.DeploymentController().DeployDeployment)
	apiGroup.DELETE("api/v1/deployment/:name", api.DeploymentController().DeleteDeployment)
	apiGroup.GET("api/v1/deployment/stats", api.DeploymentController().GetDeploymentStats)
	apiGroup.GET("api/v1/deployment/:name/pods", api.DeploymentController().GetDeploymentPodList)

	// Endpoints
	apiGroup.GET("api/v1/endpoints", api.EndpointsController().GetEndpointsList)
	apiGroup.GET("api/v1/endpoints/:name", api.EndpointsController().GetEndpointsDetails)
	apiGroup.POST("api/v1/endpoints", api.EndpointsController().DeployEndpoints)
	apiGroup.DELETE("api/v1/endpoints/:name", api.EndpointsController().DeleteEndpoints)

	// EndpointSlice
	apiGroup.GET("api/v1/endpoint-slice", api.EndpointSliceController().GetEndpointSliceList)
	apiGroup.GET("api/v1/endpoint-slice/:name", api.EndpointSliceController().GetEndpointSliceDetails)
	apiGroup.POST("api/v1/endpoint-slice", api.EndpointSliceController().DeployEndpointSlice)
	apiGroup.DELETE("api/v1/endpoint-slice/:name", api.EndpointSliceController().DeleteEndpointSlice)

	// event
	apiGroup.GET("api/v1/event", api.EventController().GetEventList)
	apiGroup.GET("api/v1/event/:name", api.EventController().GetEventDetails)

	// HPA
	apiGroup.GET("api/v1/hpa", api.HpaController().GetHpaList)
	apiGroup.GET("api/v1/hpa/:name", api.HpaController().GetHpaDetails)

	// Ingress
	apiGroup.GET("api/v1/ingress", api.IngressController().GetIngressList)
	apiGroup.GET("api/v1/ingress/:name", api.IngressController().GetIngressDetails)
	apiGroup.POST("api/v1/ingress", api.IngressController().DeployIngress)
	apiGroup.DELETE("api/v1/ingress/:name", api.IngressController().DeleteIngress)

	// Istio Gateway
	apiGroup.GET("api/v1/gateway", api.IstioGatewayController().GetIstioGatewayList)
	apiGroup.GET("api/v1/gateway/:name", api.IstioGatewayController().GetIstioGatewayDetails)
	apiGroup.POST("api/v1/gateway", api.IstioGatewayController().DeployIstioGateway)
	apiGroup.DELETE("api/v1/gateway/:name", api.IstioGatewayController().DeleteIstioGateway)

	//Job
	apiGroup.GET("api/v1/job", api.JobController().GetJobList)
	apiGroup.GET("api/v1/job/:name", api.JobController().GetJobDetails)
	apiGroup.POST("api/v1/job", api.JobController().DeployJob)
	apiGroup.DELETE("api/v1/job/:name", api.JobController().DeleteJob)

	//Load Balancer
	apiGroup.GET("api/v1/load-balancer", api.LoadBalancerController().GetLoadBalancerList)
	apiGroup.GET("api/v1/load-balancer/:name", api.LoadBalancerController().GetLoadBalancerDetails)

	// Manifest
	apiGroup.POST("api/v1/manifest", api.ManifestController().DeployManifest)

	// Network Policy
	apiGroup.GET("api/v1/network-policy", api.NetworkPolicyController().GetNetworkPolicyList)
	apiGroup.GET("api/v1/network-policy/:name", api.NetworkPolicyController().GetNetworkPolicyDetails)

	//Node
	apiGroup.GET("api/v1/node", api.NodeController().GetNodeList)
	apiGroup.GET("api/v1/node/:name", api.NodeController().GetNodeDetails)
	apiGroup.GET("api/v1/node/cordon/:name", api.NodeController().NodeCordon)
	apiGroup.POST("api/v1/node/taint/:name", api.NodeController().NodeTaint)
	apiGroup.POST("api/v1/node/untaint/:name", api.NodeController().NodeUnTaint)

	// Pod
	apiGroup.GET("api/v1/pod", api.PodController().GetPodList)
	apiGroup.GET("api/v1/pod/:name", api.PodController().GetPodDetails)
	apiGroup.GET("api/v1/pod/logs/:name", api.PodController().GetPodLogs)
	apiGroup.POST("api/v1/pod", api.PodController().DeployPod)
	apiGroup.DELETE("api/v1/pod/:name", api.PodController().DeletePod)
	apiGroup.GET("api/v1/pod/stats", api.PodController().GetPodStats)

	// PodDisruptionBudgets
	apiGroup.GET("api/v1/PDB", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsList)
	apiGroup.GET("api/v1/PDB/:name", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsDetails)
	apiGroup.POST("api/v1/PDB", api.PodDisruptionBudgetsController().DeployPodDisruptionBudgets)
	apiGroup.DELETE("api/v1/PDB/:name", api.PodDisruptionBudgetsController().DeletePodDisruptionBudgets)

	// Pod Metrics
	apiGroup.GET("api/v1/pod-metrics", api.PodMetricsController().GetPodMetricsList)
	apiGroup.GET("api/v1/pod-metrics/:pod", api.PodMetricsController().GetPodMetricsDetails)

	// PV
	apiGroup.GET("api/v1/pv", api.PvController().GetPvList)
	apiGroup.GET("api/v1/pv/:name", api.PvController().GetPvDetails)
	apiGroup.POST("api/v1/pv", api.PvController().DeployPv)
	apiGroup.DELETE("api/v1/pv/:name", api.PvController().DeletePv)

	// Persistent Volume Claim
	apiGroup.GET("api/v1/pvc", api.PvcController().GetPvcList)
	apiGroup.GET("api/v1/pvc/:name", api.PvcController().GetPvcDetails)
	apiGroup.POST("api/v1/pvc", api.PvcController().DeployPvc)
	apiGroup.DELETE("api/v1/pvc/:name", api.PvcController().DeletePvc)

	// ReplicaSet
	apiGroup.GET("api/v1/replicaset", api.ReplicaSetController().GetReplicaSetList)
	apiGroup.GET("api/v1/replicaset/:name", api.ReplicaSetController().GetReplicaSetDetails)
	apiGroup.GET("api/v1/replicaset/stats", api.ReplicaSetController().GetReplicaSetStats)
	apiGroup.POST("api/v1/replicaset", api.ReplicaSetController().DeployReplicaSet)
	apiGroup.DELETE("api/v1/replicaset/:name", api.ReplicaSetController().DeleteReplicaSet)

	// ReplicationController
	apiGroup.GET("api/v1/replication-controller", api.ReplicationControllerController().GetReplicationControllerList)
	apiGroup.GET("api/v1/replication-controller/:name", api.ReplicationControllerController().GetReplicationControllerDetails)
	apiGroup.POST("api/v1/replication-controller", api.ReplicationControllerController().DeployReplicationController)
	apiGroup.DELETE("api/v1/replication-controller/:name", api.ReplicationControllerController().DeleteReplicationController)

	// Resource Quota
	apiGroup.GET("api/v1/resource-quota", api.ResourceQuotaController().GetResourceQuotaList)
	apiGroup.GET("api/v1/resource-quota/:name", api.ResourceQuotaController().GetResourceQuotaDetails)
	apiGroup.POST("api/v1/resource-quota", api.ResourceQuotaController().DeployResourceQuota)
	apiGroup.DELETE("api/v1/resource-quota/:name", api.ResourceQuotaController().DeleteResourceQuota)

	// Role
	apiGroup.GET("api/v1/role", api.RoleController().GetRoleList)
	apiGroup.GET("api/v1/role/:name", api.RoleController().GetRoleDetails)
	apiGroup.POST("api/v1/role", api.RoleController().DeployRole)
	apiGroup.DELETE("api/v1/role/:name", api.RoleController().DeleteRole)

	// RoleBinding
	apiGroup.GET("api/v1/role-binding", api.RoleBindingController().GetRoleBindingList)
	apiGroup.GET("api/v1/role-binding/:name", api.RoleBindingController().GetRoleBindingDetails)
	apiGroup.POST("api/v1/role-binding", api.RoleBindingController().DeployRoleBinding)
	apiGroup.DELETE("api/v1/role-binding/:name", api.RoleBindingController().DeleteRoleBinding)

	// Service Account
	apiGroup.GET("api/v1/service-account", api.ServiceAccountController().GetServiceAccountList)
	apiGroup.GET("api/v1/service-account/:name", api.ServiceAccountController().GetServiceAccountDetails)
	apiGroup.POST("api/v1/service-account", api.ServiceAccountController().DeployServiceAccount)
	apiGroup.DELETE("api/v1/service-account/:name", api.ServiceAccountController().DeleteServiceAccount)

	// Secret
	apiGroup.GET("api/v1/secret", api.SecretController().GetSecretList)
	apiGroup.GET("api/v1/secret/:name", api.SecretController().GetSecretDetails)
	apiGroup.POST("api/v1/secret", api.SecretController().DeploySecret)
	apiGroup.DELETE("api/v1/secret/:name", api.SecretController().DeleteSecret)

	// StatefulSet
	apiGroup.GET("api/v1/statefulset", api.StatefulSetController().GetStatefulSetList)
	apiGroup.GET("api/v1/statefulset/:name", api.StatefulSetController().GetStatefulSetDetails)
	apiGroup.POST("api/v1/statefulset", api.StatefulSetController().DeployStatefulSet)
	apiGroup.DELETE("api/v1/statefulset/:name", api.StatefulSetController().DeleteStatefulSet)
	apiGroup.GET("api/v1/statefulset/stats", api.StatefulSetController().GetStatefulSetStats)
	apiGroup.GET("api/v1/statefulset/:name/pods", api.StatefulSetController().GetStatefulSetPodList)

	// Storage class
	apiGroup.GET("api/v1/storage-class", api.StorageClassController().GetStorageClassList)
	apiGroup.GET("api/v1/storage-class/:name", api.StorageClassController().GetStorageClassDetails)
	apiGroup.POST("api/v1/storage-class", api.StorageClassController().DeployStorageClass)
	apiGroup.DELETE("api/v1/storage-class/:name", api.StorageClassController().DeleteStorageClass)

	// Service
	apiGroup.GET("api/v1/service", api.SvcController().GetSvcList)
	apiGroup.GET("api/v1/service/:name", api.SvcController().GetSvcDetails)
	apiGroup.POST("api/v1/service", api.SvcController().DeploySVC)
	apiGroup.DELETE("api/v1/service/:name", api.SvcController().DeleteSvc)

	// Virtual Service
	apiGroup.GET("api/v1/virtual-service", api.VirtualServiceController().GetVirtualServiceList)
	apiGroup.GET("api/v1/virtual-service/:name", api.VirtualServiceController().GetVirtualServiceDetails)
	apiGroup.POST("api/v1/virtual-service", api.VirtualServiceController().DeployVirtualService)
	apiGroup.DELETE("api/v1/virtual-service/:name", api.VirtualServiceController().DeleteVirtualService)

	// Volume Snapshot
	apiGroup.GET("api/v1/volume-snapshot", api.VolumeSnapshotController().GetVolumeSnapshotList)
	apiGroup.GET("api/v1/volume-snapshot/:name", api.VolumeSnapshotController().GetVolumeSnapshotDetails)
	apiGroup.POST("api/v1/volume-snapshot", api.VolumeSnapshotController().DeployVolumeSnapshot)
	apiGroup.DELETE("api/v1/volume-snapshot/:name", api.VolumeSnapshotController().DeleteVolumeSnapshot)

	// Volume Snapshot Content
	apiGroup.GET("api/v1/volume-snapshot-content", api.VolumeSnapshotContentController().GetVolumeSnapshotContentList)
	apiGroup.GET("api/v1/volume-snapshot-content/:name", api.VolumeSnapshotContentController().GetVolumeSnapshotContentDetails)

	// Volume Snapshot Class
	apiGroup.GET("api/v1/volume-snapshot-class", api.VolumeSnapshotClassController().GetVolumeSnapshotClassList)
	apiGroup.GET("api/v1/volume-snapshot-class/:name", api.VolumeSnapshotClassController().GetVolumeSnapshotClassDetails)
}
