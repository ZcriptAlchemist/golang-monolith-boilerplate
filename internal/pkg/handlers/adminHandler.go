package handlers

import (
	"net/http"
	"sqlc-demo/internal/pkg/db/sqlc"
	"sqlc-demo/internal/pkg/service"

	"github.com/gin-gonic/gin"
)

// -------------------
// Creates admin user
// -------------------
func CreateAdmin(context *gin.Context) {
	var params sqlc.CreateAdminParams

	// Bind the JSON payload to the struct and validate
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateAdmin(params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status code": http.StatusInternalServerError, "message": "Server error"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status code": http.StatusOK, "message": "created admin successfully"})
}

// ----------------------
// Fetches Admin details
// ----------------------
func GetAllAdmins(context *gin.Context) {
	admins, err := service.FetchAdmins()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status code": http.StatusOK,
		"message":     admins,
	})
}
