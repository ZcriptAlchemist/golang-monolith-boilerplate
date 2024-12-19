package routes

import (
	"sqlc-demo/internal/pkg/handlers"

	"github.com/gin-gonic/gin"
)

// -------------
// Admin routes
// -------------
func AdminRoutes(api *gin.RouterGroup) {
	// Creates Admin
	api.POST("/admin/create", handlers.CreateAdmin)
	// Get All Admins
	api.GET("/admin/get-all-admins", handlers.GetAllAdmins)
}
