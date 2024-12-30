package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ------------------------------
// Initializes the Gin Router Server
// ------------------------------
func IntializeRouter() (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()

	// Define the /health endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "Application up and running successfully",
		})
	})

	// Group API under /api/v1
	api := router.Group("/api/v1")

	return router, api
}
