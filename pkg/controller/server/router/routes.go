package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/auth/authApi"
	controllers2 "github.com/krack8/lighthouse/pkg/controller/auth/controllers"
	api2 "github.com/krack8/lighthouse/pkg/controller/rest/api"
)

var userController *controllers2.UserController

// Declare the userService as a global variable
var rbacController *controllers2.RbacController

func AddApiRoutes(httpRg *gin.RouterGroup) {

	// User routes
	httpRg.POST("/users", authApi.UserController.CreateUserHandler)
	httpRg.GET("/users", authApi.UserController.GetAllUsersHandler)
	httpRg.GET("/users/:id", authApi.UserController.GetUserHandler)
	httpRg.PUT("/users/:id", authApi.UserController.UpdateUserHandler)
	httpRg.DELETE("/users/:id", authApi.UserController.DeleteUserHandler)
	httpRg.GET("/users/profile", authApi.UserController.GetUserProfileInfoHandler)
	httpRg.POST("/users/:id/reset-password", authApi.UserController.ResetPasswordHandler)
	//httpRg.POST("/forgot-password", authApi.UserController.ForgotPasswordHandler) 	//TODO: need to integrate mail server

	// Cluster routes
	httpRg.GET("/clusters", authApi.ClusterController.GetAllClustersHandler)
	httpRg.GET("/clusters/:id", authApi.ClusterController.GetClusterHandler)
	httpRg.POST("/clusters", authApi.ClusterController.CreateAgentClusterHandler)
	httpRg.GET("/clusters/:id/details", authApi.ClusterController.GetClusterHelmDetailsHandler)
	httpRg.DELETE("/clusters/:id", authApi.ClusterController.DeleteClusterHandler)

	// RBAC routes
	httpRg.POST("/roles", authApi.RbacController.CreateRoleHandler)
	httpRg.POST("/assign-roles", authApi.RbacController.AssignRolesHandler)
	httpRg.GET("/permissions", authApi.RbacController.GetAllPermissionsHandler)
	httpRg.GET("/permissions/:id", authApi.RbacController.GetPermissionByIDHandler)
	httpRg.GET("/permissions/users", authApi.RbacController.GetUserPermissionsHandler)
	httpRg.GET("/roles", authApi.RbacController.GetAllRolesHandler)
	httpRg.GET("/roles/:id", authApi.RbacController.GetRoleByIDHandler)
	httpRg.PUT("/roles/:id", authApi.RbacController.UpdateRoleHandler)
	httpRg.DELETE("/roles/:id", authApi.RbacController.DeleteRoleHandler)
	httpRg.GET("/roles/:id/users", authApi.RbacController.GetUsersByRoleIDHandler)

	// Namespace
	httpRg.GET("/namespace", api2.NamespaceController().GetNamespaceList)
	httpRg.GET("/namespace/names", api2.NamespaceController().GetNamespaceNameList)
	httpRg.GET("/namespace/:name", api2.NamespaceController().GetNamespaceDetails)
	httpRg.POST("/namespace", api2.NamespaceController().DeployNamespace)
	httpRg.DELETE("/namespace/:name", api2.NamespaceController().DeleteNamespace)

	// Certificate
	httpRg.GET("/certificate", api2.CertificateController().GetCertificateList)
	httpRg.GET("/certificate/:name", api2.CertificateController().GetCertificateDetails)
	httpRg.POST("/certificate", api2.CertificateController().DeployCertificate)
	httpRg.DELETE("/certificate/:name", api2.CertificateController().DeleteCertificate)

	// Config Map
	httpRg.GET("/config-map", api2.ConfigMapController().GetConfigMapList)
	httpRg.GET("/config-map/:name", api2.ConfigMapController().GetConfigMapDetails)
	httpRg.POST("/config-map", api2.ConfigMapController().DeployConfigMap)
	httpRg.DELETE("/config-map/:name", api2.ConfigMapController().DeleteConfigMap)

	// ClusterRole
	httpRg.GET("/cluster-role", api2.ClusterRoleController().GetClusterRoleList)
	httpRg.GET("/cluster-role/:name", api2.ClusterRoleController().GetClusterRoleDetails)
	httpRg.POST("/cluster-role", api2.ClusterRoleController().DeployClusterRole)
	httpRg.DELETE("/cluster-role/:name", api2.ClusterRoleController().DeleteClusterRole)

	// ClusterRoleBinding
	httpRg.GET("/cluster-role-binding", api2.ClusterRoleBindingController().GetClusterRoleBindingList)
	httpRg.GET("/cluster-role-binding/:name", api2.ClusterRoleBindingController().GetClusterRoleBindingDetails)
	httpRg.POST("/cluster-role-binding", api2.ClusterRoleBindingController().DeployClusterRoleBinding)
	httpRg.DELETE("/cluster-role-binding/:name", api2.ClusterRoleBindingController().DeleteClusterRoleBinding)

	// ControllerRevision
	httpRg.GET("/controller-revision", api2.ControllerRevisionController().GetControllerRevisionList)
	httpRg.GET("/controller-revision/:name", api2.ControllerRevisionController().GetControllerRevisionDetails)
	httpRg.POST("/controller-revision", api2.ControllerRevisionController().DeployControllerRevision)
	httpRg.DELETE("/controller-revision/:name", api2.ControllerRevisionController().DeleteControllerRevision)

	// CRD
	httpRg.GET("/crd", api2.CrdController().GetCrdList)
	httpRg.GET("/crd/:name", api2.CrdController().GetCrdDetails)
	httpRg.POST("/crd", api2.CrdController().DeployCrd)
	httpRg.DELETE("/crd/:name", api2.CrdController().DeleteCrd)

	// Custom Resource
	httpRg.GET("/custom-resource", api2.CustomResourceController().GetCustomResourceList)
	httpRg.GET("/custom-resource/:name", api2.CustomResourceController().GetCustomResourceDetails)
	httpRg.POST("/custom-resource", api2.CustomResourceController().DeployCustomResource)
	httpRg.DELETE("/custom-resource/:name", api2.CustomResourceController().DeleteCustomResource)

	//Cronjob
	httpRg.GET("/cronjob", api2.CronJobController().GetCronJobList)
	httpRg.GET("/cronjob/:name", api2.CronJobController().GetCronJobDetails)
	httpRg.POST("/cronjob", api2.CronJobController().DeployCronJob)
	httpRg.DELETE("/cronjob/:name", api2.CronJobController().DeleteCronJob)

	// Daemonset
	httpRg.GET("/daemonset", api2.DaemonSetController().GetDaemonSetList)
	httpRg.GET("/daemonset/:name", api2.DaemonSetController().GetDaemonSetDetails)
	httpRg.POST("/daemonset", api2.DaemonSetController().DeployDaemonSet)
	httpRg.DELETE("/daemonset/:name", api2.DaemonSetController().DeleteDaemonSet)
	httpRg.GET("/daemonset/stats", api2.DaemonSetController().GetDaemonSetStats)

	// Deployment
	httpRg.GET("/deployment", api2.DeploymentController().GetDeploymentList)
	httpRg.GET("/deployment/:name", api2.DeploymentController().GetDeploymentDetails)
	httpRg.POST("/deployment", api2.DeploymentController().DeployDeployment)
	httpRg.DELETE("/deployment/:name", api2.DeploymentController().DeleteDeployment)
	httpRg.GET("/deployment/stats", api2.DeploymentController().GetDeploymentStats)
	httpRg.GET("/deployment/:name/pods", api2.DeploymentController().GetDeploymentPodList)

	// Endpoints
	httpRg.GET("/endpoints", api2.EndpointsController().GetEndpointsList)
	httpRg.GET("/endpoints/:name", api2.EndpointsController().GetEndpointsDetails)
	httpRg.POST("/endpoints", api2.EndpointsController().DeployEndpoints)
	httpRg.DELETE("/endpoints/:name", api2.EndpointsController().DeleteEndpoints)

	// EndpointSlice
	httpRg.GET("/endpoint-slice", api2.EndpointSliceController().GetEndpointSliceList)
	httpRg.GET("/endpoint-slice/:name", api2.EndpointSliceController().GetEndpointSliceDetails)
	httpRg.POST("/endpoint-slice", api2.EndpointSliceController().DeployEndpointSlice)
	httpRg.DELETE("/endpoint-slice/:name", api2.EndpointSliceController().DeleteEndpointSlice)

	// event
	httpRg.GET("/event", api2.EventController().GetEventList)
	httpRg.GET("/event/:name", api2.EventController().GetEventDetails)

	// HPA
	httpRg.GET("/hpa", api2.HpaController().GetHpaList)
	httpRg.GET("/hpa/:name", api2.HpaController().GetHpaDetails)

	// Ingress
	httpRg.GET("/ingress", api2.IngressController().GetIngressList)
	httpRg.GET("/ingress/:name", api2.IngressController().GetIngressDetails)
	httpRg.POST("/ingress", api2.IngressController().DeployIngress)
	httpRg.DELETE("/ingress/:name", api2.IngressController().DeleteIngress)

	// Istio Gateway
	httpRg.GET("/gateway", api2.IstioGatewayController().GetIstioGatewayList)
	httpRg.GET("/gateway/:name", api2.IstioGatewayController().GetIstioGatewayDetails)
	httpRg.POST("/gateway", api2.IstioGatewayController().DeployIstioGateway)
	httpRg.DELETE("/gateway/:name", api2.IstioGatewayController().DeleteIstioGateway)

	//Job
	httpRg.GET("/job", api2.JobController().GetJobList)
	httpRg.GET("/job/:name", api2.JobController().GetJobDetails)
	httpRg.POST("/job", api2.JobController().DeployJob)
	httpRg.DELETE("/job/:name", api2.JobController().DeleteJob)

	//Load Balancer
	httpRg.GET("/load-balancer", api2.LoadBalancerController().GetLoadBalancerList)
	httpRg.GET("/load-balancer/:name", api2.LoadBalancerController().GetLoadBalancerDetails)

	// Manifest
	httpRg.POST("/manifest", api2.ManifestController().DeployManifest)

	// Network Policy
	httpRg.GET("/network-policy", api2.NetworkPolicyController().GetNetworkPolicyList)
	httpRg.GET("/network-policy/:name", api2.NetworkPolicyController().GetNetworkPolicyDetails)

	//Node
	httpRg.GET("/node", api2.NodeController().GetNodeList)
	httpRg.GET("/node/:name", api2.NodeController().GetNodeDetails)
	httpRg.GET("/node/cordon/:name", api2.NodeController().NodeCordon)
	httpRg.POST("/node/taint/:name", api2.NodeController().NodeTaint)
	httpRg.POST("/node/untaint/:name", api2.NodeController().NodeUnTaint)

	// Pod
	httpRg.GET("/pod", api2.PodController().GetPodList)
	httpRg.GET("/pod/:name", api2.PodController().GetPodDetails)
	httpRg.GET("/pod/logs/:name", api2.PodController().GetPodLogs)
	httpRg.POST("/pod", api2.PodController().DeployPod)
	httpRg.DELETE("/pod/:name", api2.PodController().DeletePod)
	httpRg.GET("/pod/stats", api2.PodController().GetPodStats)

	// PodDisruptionBudgets
	httpRg.GET("/PDB", api2.PodDisruptionBudgetsController().GetPodDisruptionBudgetsList)
	httpRg.GET("/PDB/:name", api2.PodDisruptionBudgetsController().GetPodDisruptionBudgetsDetails)
	httpRg.POST("/PDB", api2.PodDisruptionBudgetsController().DeployPodDisruptionBudgets)
	httpRg.DELETE("/PDB/:name", api2.PodDisruptionBudgetsController().DeletePodDisruptionBudgets)

	// Pod Metrics
	httpRg.GET("/pod-metrics", api2.PodMetricsController().GetPodMetricsList)
	httpRg.GET("/pod-metrics/:pod", api2.PodMetricsController().GetPodMetricsDetails)

	// PV
	httpRg.GET("/pv", api2.PvController().GetPvList)
	httpRg.GET("/pv/:name", api2.PvController().GetPvDetails)
	httpRg.POST("/pv", api2.PvController().DeployPv)
	httpRg.DELETE("/pv/:name", api2.PvController().DeletePv)

	// Persistent Volume Claim
	httpRg.GET("/pvc", api2.PvcController().GetPvcList)
	httpRg.GET("/pvc/:name", api2.PvcController().GetPvcDetails)
	httpRg.POST("/pvc", api2.PvcController().DeployPvc)
	httpRg.DELETE("/pvc/:name", api2.PvcController().DeletePvc)

	// ReplicaSet
	httpRg.GET("/replicaset", api2.ReplicaSetController().GetReplicaSetList)
	httpRg.GET("/replicaset/:name", api2.ReplicaSetController().GetReplicaSetDetails)
	httpRg.POST("/replicaset", api2.ReplicaSetController().DeployReplicaSet)
	httpRg.DELETE("/replicaset/:name", api2.ReplicaSetController().DeleteReplicaSet)

	// ReplicationController
	httpRg.GET("/replication-controller", api2.ReplicationControllerController().GetReplicationControllerList)
	httpRg.GET("/replication-controller/:name", api2.ReplicationControllerController().GetReplicationControllerDetails)
	httpRg.POST("/replication-controller", api2.ReplicationControllerController().DeployReplicationController)
	httpRg.DELETE("/replication-controller/:name", api2.ReplicationControllerController().DeleteReplicationController)

	// Resource Quota
	httpRg.GET("/resource-quota", api2.ResourceQuotaController().GetResourceQuotaList)
	httpRg.GET("/resource-quota/:name", api2.ResourceQuotaController().GetResourceQuotaDetails)
	httpRg.POST("/resource-quota", api2.ResourceQuotaController().DeployResourceQuota)
	httpRg.DELETE("/resource-quota/:name", api2.ResourceQuotaController().DeleteResourceQuota)

	// Role
	httpRg.GET("/role", api2.RoleController().GetRoleList)
	httpRg.GET("/role/:name", api2.RoleController().GetRoleDetails)
	httpRg.POST("/role", api2.RoleController().DeployRole)
	httpRg.DELETE("/role/:name", api2.RoleController().DeleteRole)

	// RoleBinding
	httpRg.GET("/role-binding", api2.RoleBindingController().GetRoleBindingList)
	httpRg.GET("/role-binding/:name", api2.RoleBindingController().GetRoleBindingDetails)
	httpRg.POST("/role-binding", api2.RoleBindingController().DeployRoleBinding)
	httpRg.DELETE("/role-binding/:name", api2.RoleBindingController().DeleteRoleBinding)

	// Service Account
	httpRg.GET("/service-account", api2.ServiceAccountController().GetServiceAccountList)
	httpRg.GET("/service-account/:name", api2.ServiceAccountController().GetServiceAccountDetails)
	httpRg.POST("/service-account", api2.ServiceAccountController().DeployServiceAccount)
	httpRg.DELETE("/service-account/:name", api2.ServiceAccountController().DeleteServiceAccount)

	// Secret
	httpRg.GET("/secret", api2.SecretController().GetSecretList)
	httpRg.GET("/secret/:name", api2.SecretController().GetSecretDetails)
	httpRg.POST("/secret", api2.SecretController().DeploySecret)
	httpRg.DELETE("/secret/:name", api2.SecretController().DeleteSecret)

	// StatefulSet
	httpRg.GET("/statefulset", api2.StatefulSetController().GetStatefulSetList)
	httpRg.GET("/statefulset/:name", api2.StatefulSetController().GetStatefulSetDetails)
	httpRg.POST("/statefulset", api2.StatefulSetController().DeployStatefulSet)
	httpRg.DELETE("/statefulset/:name", api2.StatefulSetController().DeleteStatefulSet)
	httpRg.GET("/statefulset/stats", api2.StatefulSetController().GetStatefulSetStats)
	httpRg.GET("/statefulset/:name/pods", api2.StatefulSetController().GetStatefulSetPodList)

	// Storage class
	httpRg.GET("/storage-class", api2.StorageClassController().GetStorageClassList)
	httpRg.GET("/storage-class/:name", api2.StorageClassController().GetStorageClassDetails)
	httpRg.POST("/storage-class", api2.StorageClassController().DeployStorageClass)
	httpRg.DELETE("/storage-class/:name", api2.StorageClassController().DeleteStorageClass)

	// Service
	httpRg.GET("/service", api2.SvcController().GetSvcList)
	httpRg.GET("/service/:name", api2.SvcController().GetSvcDetails)
	httpRg.POST("/service", api2.SvcController().DeploySVC)
	httpRg.DELETE("/service/:name", api2.SvcController().DeleteSvc)

	// Virtual Service
	httpRg.GET("/virtual-service", api2.VirtualServiceController().GetVirtualServiceList)
	httpRg.GET("/virtual-service/:name", api2.VirtualServiceController().GetVirtualServiceDetails)
	httpRg.POST("/virtual-service", api2.VirtualServiceController().DeployVirtualService)
	httpRg.DELETE("/virtual-service/:name", api2.VirtualServiceController().DeleteVirtualService)

	// Volume Snapshot
	httpRg.GET("/volume-snapshot", api2.VolumeSnapshotController().GetVolumeSnapshotList)
	httpRg.GET("/volume-snapshot/:name", api2.VolumeSnapshotController().GetVolumeSnapshotDetails)
	httpRg.POST("/volume-snapshot", api2.VolumeSnapshotController().DeployVolumeSnapshot)
	httpRg.DELETE("/volume-snapshot/:name", api2.VolumeSnapshotController().DeleteVolumeSnapshot)

	// Volume Snapshot Content
	httpRg.GET("/volume-snapshot-content", api2.VolumeSnapshotContentController().GetVolumeSnapshotContentList)
	httpRg.GET("/volume-snapshot-content/:name", api2.VolumeSnapshotContentController().GetVolumeSnapshotContentDetails)

	// Volume Snapshot Class
	httpRg.GET("/volume-snapshot-class", api2.VolumeSnapshotClassController().GetVolumeSnapshotClassList)
	httpRg.GET("/volume-snapshot-class/:name", api2.VolumeSnapshotClassController().GetVolumeSnapshotClassDetails)
}

func AddWsApiRoutes(wsRg *gin.RouterGroup) {
	wsRg.GET("/pod/:name/exec", api2.PodController().ExecPod)
}
