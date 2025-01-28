package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/auth/controllers"
	middleware "github.com/krack8/lighthouse/pkg/auth/middlewares"
	"github.com/krack8/lighthouse/pkg/controller/api"
)

// @title           Swagger API
// @version         1.0

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

// Declare the userService as a global variable
var userController *controllers.UserController

// Declare the userService as a global variable
var rbacController *controllers.RbacController

// Declare the userService as a global variable

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
}
