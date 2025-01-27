package router

import (
	"github.com/gin-gonic/gin"
	"github.com/krack8/lighthouse/pkg/controller/api"
)

// @title           Swagger API
// @version         1.0

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func AddApiRoutes(httpRg *gin.RouterGroup) {
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
}
