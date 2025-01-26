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
}
