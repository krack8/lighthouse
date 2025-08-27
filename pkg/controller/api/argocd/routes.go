package argocd

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all ArgoCD routes
func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	// Agent management routes
	agents := router.Group("/agents")
	{
		agents.POST("/register", handler.RegisterAgent)
		agents.POST("/heartbeat", handler.Heartbeat)
		agents.GET("", handler.ListAgents)
	}

	// ArgoCD routes
	argocd := router.Group("/argocd")
	{
		// Application routes
		apps := argocd.Group("/applications")
		{
			apps.GET("", handler.ListApplications)
			apps.POST("", handler.CreateApplication)
			apps.GET("/:name", handler.GetApplication)
			apps.PUT("/:name", handler.UpdateApplication)
			apps.DELETE("/:name", handler.DeleteApplication)
			apps.POST("/:name/sync", handler.SyncApplication)
			apps.POST("/:name/rollback", handler.RollbackApplication)
		}

		// Project routes
		projects := argocd.Group("/projects")
		{
			projects.GET("", handler.ListProjects)
			projects.POST("", handler.CreateProject)
		}

		// Repository routes
		repos := argocd.Group("/repositories")
		{
			repos.GET("", handler.ListRepositories)
			repos.POST("", handler.CreateRepository)
		}
	}
}
