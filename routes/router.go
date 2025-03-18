package routes

import (
	"github.com/gin-gonic/gin"

	controllers "go-server/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Settings routes
	router.POST("/settings", controllers.CreateOrganizationSetting)
	router.GET("/settings/:organizationId", controllers.GetOrganizationSettings)
	router.GET("/settings/:organizationId/:key", controllers.GetOrganizationSetting)
	router.PUT("/settings/:organizationId", controllers.UpdateOrganizationSetting)

	// AI Response routes
	router.POST("/ai-responses", controllers.CreateAIResponse)
	router.GET("/ai-responses/:organizationId", controllers.GetOrganizationAIResponses)
	router.GET("/ai-responses/:organizationId/:id", controllers.GetOrganizationAIResponse)
	router.POST("/ai-responses/:organizationId/:id/feedback", controllers.CreateAIResponseFeedback)

	return router
}
