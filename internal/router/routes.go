package router

import (
	"sqlc-demo/internal/apis"

	"github.com/gin-gonic/gin"
)

// -------------
// Admin routes
// -------------
func AdminRoutes(api *gin.RouterGroup) {
	// Creates Admin
	api.POST("/admin/create", apis.CreateAdmin)
	// Get All Admins
	api.GET("/admin/get-all-admins", apis.GetAllAdmins)
}

// -------------
// Admin routes
// -------------
func MerchantRoutes(api *gin.RouterGroup) {
	// Creates Admin
	api.POST("/merchant/create", apis.CreateMerchant)
	// Get All Admins
	api.GET("/merchant/get-all-merchant", apis.GetAllMerchants)
}
