package apis

import (
	"context"
	"net/http"
	"sqlc-demo/internal"
	"sqlc-demo/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// -------------------------
// Creates Merchant profile
// --------------------------
func CreateMerchant(c *gin.Context) {
	var params sqlc.CreateMerchantParams

	// Bind the JSON payload to the struct and validate
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := internal.DB.CreateMerchant(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status code": http.StatusInternalServerError, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status code": http.StatusOK, "message": "created admin successfully"})
}

// ------------------------------
// Fetches All Merchant profiles
// ------------------------------
func GetAllMerchants(c *gin.Context) {
	admins, err := internal.DB.ListAllMerchants(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status code": http.StatusOK,
		"message":     admins,
	})
}
