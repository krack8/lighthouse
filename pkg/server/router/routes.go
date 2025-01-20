package router

import (
	"github.com/gin-gonic/gin"
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
	//httpRg.GET("api/v1/namespace", k8s.NamespaceController().GetNamespaceList)
}
