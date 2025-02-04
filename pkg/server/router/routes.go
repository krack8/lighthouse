package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/authApi"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	"github.com/krack8/lighthouse/pkg/controller/api"
)

var userController *controllers.UserController

// Declare the userService as a global variable
var rbacController *controllers.RbacController

func AddApiRoutes(httpRg *gin.RouterGroup) {

	// User routes
	httpRg.POST("/users", authApi.UserController.CreateUserHandler)
	httpRg.GET("/users", authApi.UserController.GetAllUsersHandler)
	httpRg.GET("/users/:id", authApi.UserController.GetUserHandler)
	httpRg.PUT("/users/:id", authApi.UserController.UpdateUserHandler)
	httpRg.DELETE("/users/:id", authApi.UserController.DeleteUserHandler)
	httpRg.GET("/users/profile", authApi.UserController.GetUserProfileInfoHandler)

	// Cluster routes
	httpRg.GET("/clusters", authApi.ClusterController.GetAllClustersHandler)
	httpRg.GET("/clusters/:id", authApi.ClusterController.GetClusterHandler)

	// RBAC routes
	//httpRg.POST("/permissions", authApi.RbacController.CreatePermissionHandler)   /test-api to create permission for dev
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
	httpRg.GET("/namespace", api.NamespaceController().GetNamespaceList)
	httpRg.GET("/namespace/names", api.NamespaceController().GetNamespaceNameList)
	httpRg.GET("/namespace/:name", api.NamespaceController().GetNamespaceDetails)
	httpRg.POST("/namespace", api.NamespaceController().DeployNamespace)
	httpRg.DELETE("/namespace/:name", api.NamespaceController().DeleteNamespace)

	// Certificate
	httpRg.GET("/certificate", api.CertificateController().GetCertificateList)
	httpRg.GET("/certificate/:name", api.CertificateController().GetCertificateDetails)
	httpRg.POST("/certificate", api.CertificateController().DeployCertificate)
	httpRg.DELETE("/certificate/:name", api.CertificateController().DeleteCertificate)

	// Config Map
	httpRg.GET("/config-map", api.ConfigMapController().GetConfigMapList)
	httpRg.GET("/config-map/:name", api.ConfigMapController().GetConfigMapDetails)
	httpRg.POST("/config-map", api.ConfigMapController().DeployConfigMap)
	httpRg.DELETE("/config-map/:name", api.ConfigMapController().DeleteConfigMap)

	// ClusterRole
	httpRg.GET("/cluster-role", api.ClusterRoleController().GetClusterRoleList)
	httpRg.GET("/cluster-role/:name", api.ClusterRoleController().GetClusterRoleDetails)
	httpRg.POST("/cluster-role", api.ClusterRoleController().DeployClusterRole)
	httpRg.DELETE("/cluster-role/:name", api.ClusterRoleController().DeleteClusterRole)

	// ClusterRoleBinding
	httpRg.GET("/cluster-role-binding", api.ClusterRoleBindingController().GetClusterRoleBindingList)
	httpRg.GET("/cluster-role-binding/:name", api.ClusterRoleBindingController().GetClusterRoleBindingDetails)
	httpRg.POST("/cluster-role-binding", api.ClusterRoleBindingController().DeployClusterRoleBinding)
	httpRg.DELETE("/cluster-role-binding/:name", api.ClusterRoleBindingController().DeleteClusterRoleBinding)

	// ControllerRevision
	httpRg.GET("/controller-revision", api.ControllerRevisionController().GetControllerRevisionList)
	httpRg.GET("/controller-revision/:name", api.ControllerRevisionController().GetControllerRevisionDetails)
	httpRg.POST("/controller-revision", api.ControllerRevisionController().DeployControllerRevision)
	httpRg.DELETE("/controller-revision/:name", api.ControllerRevisionController().DeleteControllerRevision)

	// CRD
	httpRg.GET("/crd", api.CrdController().GetCrdList)
	httpRg.GET("/crd/:name", api.CrdController().GetCrdDetails)
	httpRg.POST("/crd", api.CrdController().DeployCrd)
	httpRg.DELETE("/crd/:name", api.CrdController().DeleteCrd)

	// Custom Resource
	httpRg.GET("/custom-resource", api.CustomResourceController().GetCustomResourceList)
	httpRg.GET("/custom-resource/:name", api.CustomResourceController().GetCustomResourceDetails)
	httpRg.POST("/custom-resource", api.CustomResourceController().DeployCustomResource)
	httpRg.DELETE("/custom-resource/:name", api.CustomResourceController().DeleteCustomResource)

	//Cronjob
	httpRg.GET("/cronjob", api.CronJobController().GetCronJobList)
	httpRg.GET("/cronjob/:name", api.CronJobController().GetCronJobDetails)
	httpRg.POST("/cronjob", api.CronJobController().DeployCronJob)
	httpRg.DELETE("/cronjob/:name", api.CronJobController().DeleteCronJob)

	// Daemonset
	httpRg.GET("/daemonset", api.DaemonSetController().GetDaemonSetList)
	httpRg.GET("/daemonset/:name", api.DaemonSetController().GetDaemonSetDetails)
	httpRg.POST("/daemonset", api.DaemonSetController().DeployDaemonSet)
	httpRg.DELETE("/daemonset/:name", api.DaemonSetController().DeleteDaemonSet)
	httpRg.GET("/daemonset/stats", api.DaemonSetController().GetDaemonSetStats)

	// Deployment
	httpRg.GET("/deployment", api.DeploymentController().GetDeploymentList)
	httpRg.GET("/deployment/:name", api.DeploymentController().GetDeploymentDetails)
	httpRg.POST("/deployment", api.DeploymentController().DeployDeployment)
	httpRg.DELETE("/deployment/:name", api.DeploymentController().DeleteDeployment)
	httpRg.GET("/deployment/stats", api.DeploymentController().GetDeploymentStats)
	httpRg.GET("/deployment/:name/pods", api.DeploymentController().GetDeploymentPodList)

	// Endpoints
	httpRg.GET("/endpoints", api.EndpointsController().GetEndpointsList)
	httpRg.GET("/endpoints/:name", api.EndpointsController().GetEndpointsDetails)
	httpRg.POST("/endpoints", api.EndpointsController().DeployEndpoints)
	httpRg.DELETE("/endpoints/:name", api.EndpointsController().DeleteEndpoints)

	// EndpointSlice
	httpRg.GET("/endpoint-slice", api.EndpointSliceController().GetEndpointSliceList)
	httpRg.GET("/endpoint-slice/:name", api.EndpointSliceController().GetEndpointSliceDetails)
	httpRg.POST("/endpoint-slice", api.EndpointSliceController().DeployEndpointSlice)
	httpRg.DELETE("/endpoint-slice/:name", api.EndpointSliceController().DeleteEndpointSlice)

	// event
	httpRg.GET("/event", api.EventController().GetEventList)
	httpRg.GET("/event/:name", api.EventController().GetEventDetails)

	// HPA
	httpRg.GET("/hpa", api.HpaController().GetHpaList)
	httpRg.GET("/hpa/:name", api.HpaController().GetHpaDetails)

	// Ingress
	httpRg.GET("/ingress", api.IngressController().GetIngressList)
	httpRg.GET("/ingress/:name", api.IngressController().GetIngressDetails)
	httpRg.POST("/ingress", api.IngressController().DeployIngress)
	httpRg.DELETE("/ingress/:name", api.IngressController().DeleteIngress)

	// Istio Gateway
	httpRg.GET("/gateway", api.IstioGatewayController().GetIstioGatewayList)
	httpRg.GET("/gateway/:name", api.IstioGatewayController().GetIstioGatewayDetails)
	httpRg.POST("/gateway", api.IstioGatewayController().DeployIstioGateway)
	httpRg.DELETE("/gateway/:name", api.IstioGatewayController().DeleteIstioGateway)

	//Job
	httpRg.GET("/job", api.JobController().GetJobList)
	httpRg.GET("/job/:name", api.JobController().GetJobDetails)
	httpRg.POST("/job", api.JobController().DeployJob)
	httpRg.DELETE("/job/:name", api.JobController().DeleteJob)

	//Load Balancer
	httpRg.GET("/load-balancer", api.LoadBalancerController().GetLoadBalancerList)
	httpRg.GET("/load-balancer/:name", api.LoadBalancerController().GetLoadBalancerDetails)

	// Manifest
	httpRg.POST("/manifest", api.ManifestController().DeployManifest)

	// Network Policy
	httpRg.GET("/network-policy", api.NetworkPolicyController().GetNetworkPolicyList)
	httpRg.GET("/network-policy/:name", api.NetworkPolicyController().GetNetworkPolicyDetails)

	//Node
	httpRg.GET("/node", api.NodeController().GetNodeList)
	httpRg.GET("/node/:name", api.NodeController().GetNodeDetails)
	httpRg.GET("/node/cordon/:name", api.NodeController().NodeCordon)
	httpRg.POST("/node/taint/:name", api.NodeController().NodeTaint)
	httpRg.POST("/node/untaint/:name", api.NodeController().NodeUnTaint)

	// Pod
	httpRg.GET("/pod", api.PodController().GetPodList)
	httpRg.GET("/pod/:name", api.PodController().GetPodDetails)
	httpRg.GET("/pod/logs/:name", api.PodController().GetPodLogs)
	httpRg.POST("/pod", api.PodController().DeployPod)
	httpRg.DELETE("/pod/:name", api.PodController().DeletePod)
	httpRg.GET("/pod/stats", api.PodController().GetPodStats)

	// PodDisruptionBudgets
	httpRg.GET("/PDB", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsList)
	httpRg.GET("/PDB/:name", api.PodDisruptionBudgetsController().GetPodDisruptionBudgetsDetails)
	httpRg.POST("/PDB", api.PodDisruptionBudgetsController().DeployPodDisruptionBudgets)
	httpRg.DELETE("/PDB/:name", api.PodDisruptionBudgetsController().DeletePodDisruptionBudgets)

	// Pod Metrics
	httpRg.GET("/pod-metrics", api.PodMetricsController().GetPodMetricsList)
	httpRg.GET("/pod-metrics/:pod", api.PodMetricsController().GetPodMetricsDetails)

	// PV
	httpRg.GET("/pv", api.PvController().GetPvList)
	httpRg.GET("/pv/:name", api.PvController().GetPvDetails)
	httpRg.POST("/pv", api.PvController().DeployPv)
	httpRg.DELETE("/pv/:name", api.PvController().DeletePv)

	// Persistent Volume Claim
	httpRg.GET("/pvc", api.PvcController().GetPvcList)
	httpRg.GET("/pvc/:name", api.PvcController().GetPvcDetails)
	httpRg.POST("/pvc", api.PvcController().DeployPvc)
	httpRg.DELETE("/pvc/:name", api.PvcController().DeletePvc)

	// ReplicaSet
	httpRg.GET("/replicaset", api.ReplicaSetController().GetReplicaSetList)
	httpRg.GET("/replicaset/:name", api.ReplicaSetController().GetReplicaSetDetails)
	httpRg.GET("/replicaset/stats", api.ReplicaSetController().GetReplicaSetStats)
	httpRg.POST("/replicaset", api.ReplicaSetController().DeployReplicaSet)
	httpRg.DELETE("/replicaset/:name", api.ReplicaSetController().DeleteReplicaSet)

	// ReplicationController
	httpRg.GET("/replication-controller", api.ReplicationControllerController().GetReplicationControllerList)
	httpRg.GET("/replication-controller/:name", api.ReplicationControllerController().GetReplicationControllerDetails)
	httpRg.POST("/replication-controller", api.ReplicationControllerController().DeployReplicationController)
	httpRg.DELETE("/replication-controller/:name", api.ReplicationControllerController().DeleteReplicationController)

	// Resource Quota
	httpRg.GET("/resource-quota", api.ResourceQuotaController().GetResourceQuotaList)
	httpRg.GET("/resource-quota/:name", api.ResourceQuotaController().GetResourceQuotaDetails)
	httpRg.POST("/resource-quota", api.ResourceQuotaController().DeployResourceQuota)
	httpRg.DELETE("/resource-quota/:name", api.ResourceQuotaController().DeleteResourceQuota)

	// Role
	httpRg.GET("/role", api.RoleController().GetRoleList)
	httpRg.GET("/role/:name", api.RoleController().GetRoleDetails)
	httpRg.POST("/role", api.RoleController().DeployRole)
	httpRg.DELETE("/role/:name", api.RoleController().DeleteRole)

	// RoleBinding
	httpRg.GET("/role-binding", api.RoleBindingController().GetRoleBindingList)
	httpRg.GET("/role-binding/:name", api.RoleBindingController().GetRoleBindingDetails)
	httpRg.POST("/role-binding", api.RoleBindingController().DeployRoleBinding)
	httpRg.DELETE("/role-binding/:name", api.RoleBindingController().DeleteRoleBinding)

	// Service Account
	httpRg.GET("/service-account", api.ServiceAccountController().GetServiceAccountList)
	httpRg.GET("/service-account/:name", api.ServiceAccountController().GetServiceAccountDetails)
	httpRg.POST("/service-account", api.ServiceAccountController().DeployServiceAccount)
	httpRg.DELETE("/service-account/:name", api.ServiceAccountController().DeleteServiceAccount)

	// Secret
	httpRg.GET("/secret", api.SecretController().GetSecretList)
	httpRg.GET("/secret/:name", api.SecretController().GetSecretDetails)
	httpRg.POST("/secret", api.SecretController().DeploySecret)
	httpRg.DELETE("/secret/:name", api.SecretController().DeleteSecret)

	// StatefulSet
	httpRg.GET("/statefulset", api.StatefulSetController().GetStatefulSetList)
	httpRg.GET("/statefulset/:name", api.StatefulSetController().GetStatefulSetDetails)
	httpRg.POST("/statefulset", api.StatefulSetController().DeployStatefulSet)
	httpRg.DELETE("/statefulset/:name", api.StatefulSetController().DeleteStatefulSet)
	httpRg.GET("/statefulset/stats", api.StatefulSetController().GetStatefulSetStats)
	httpRg.GET("/statefulset/:name/pods", api.StatefulSetController().GetStatefulSetPodList)

	// Storage class
	httpRg.GET("/storage-class", api.StorageClassController().GetStorageClassList)
	httpRg.GET("/storage-class/:name", api.StorageClassController().GetStorageClassDetails)
	httpRg.POST("/storage-class", api.StorageClassController().DeployStorageClass)
	httpRg.DELETE("/storage-class/:name", api.StorageClassController().DeleteStorageClass)

	// Service
	httpRg.GET("/service", api.SvcController().GetSvcList)
	httpRg.GET("/service/:name", api.SvcController().GetSvcDetails)
	httpRg.POST("/service", api.SvcController().DeploySVC)
	httpRg.DELETE("/service/:name", api.SvcController().DeleteSvc)

	// Virtual Service
	httpRg.GET("/virtual-service", api.VirtualServiceController().GetVirtualServiceList)
	httpRg.GET("/virtual-service/:name", api.VirtualServiceController().GetVirtualServiceDetails)
	httpRg.POST("/virtual-service", api.VirtualServiceController().DeployVirtualService)
	httpRg.DELETE("/virtual-service/:name", api.VirtualServiceController().DeleteVirtualService)

	// Volume Snapshot
	httpRg.GET("/volume-snapshot", api.VolumeSnapshotController().GetVolumeSnapshotList)
	httpRg.GET("/volume-snapshot/:name", api.VolumeSnapshotController().GetVolumeSnapshotDetails)
	httpRg.POST("/volume-snapshot", api.VolumeSnapshotController().DeployVolumeSnapshot)
	httpRg.DELETE("/volume-snapshot/:name", api.VolumeSnapshotController().DeleteVolumeSnapshot)

	// Volume Snapshot Content
	httpRg.GET("/volume-snapshot-content", api.VolumeSnapshotContentController().GetVolumeSnapshotContentList)
	httpRg.GET("/volume-snapshot-content/:name", api.VolumeSnapshotContentController().GetVolumeSnapshotContentDetails)

	// Volume Snapshot Class
	httpRg.GET("/volume-snapshot-class", api.VolumeSnapshotClassController().GetVolumeSnapshotClassList)
	httpRg.GET("/volume-snapshot-class/:name", api.VolumeSnapshotClassController().GetVolumeSnapshotClassDetails)
}
