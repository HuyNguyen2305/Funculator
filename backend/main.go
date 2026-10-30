package main

import (
	"funculator/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	routes.SetupRoutes(router)

	router.Run(":8080")
}
