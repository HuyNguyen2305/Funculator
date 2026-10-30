package routes

import (
	"funculator/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/calculate", controllers.Calculate)
	}

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
}
