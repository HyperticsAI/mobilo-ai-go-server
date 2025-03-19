package routes

import (
	"github.com/gin-gonic/gin"

	controllers "go-server/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	// add v1 prefix
	v1 := router.Group("/api/v1")

	// Settings routes
	v1.POST("/settings", controllers.CreateOrganizationSetting)
	v1.GET("/settings/:organizationId", controllers.GetOrganizationSettings)
	v1.GET("/settings/:organizationId/:key", controllers.GetOrganizationSetting)
	v1.PUT("/settings/:organizationId", controllers.UpdateOrganizationSetting)

	// AI Response routes
	v1.POST("/ai-responses", controllers.CreateAIResponse)
	v1.GET("/ai-responses/:organizationId", controllers.GetOrganizationAIResponses)
	v1.GET("/ai-responses/:organizationId/:id", controllers.GetOrganizationAIResponse)
	v1.POST("/ai-responses/:organizationId/:id/feedback", controllers.CreateAIResponseFeedback)

	return router
}
