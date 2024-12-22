package app

import (
	"log"
	"sqlc-demo/internal"

	"sqlc-demo/internal/core"
	"sqlc-demo/internal/router"
)

// ------------------------------------------------------------------------------------------
// App starts here - Intializes all the necessary config files and setup for db, server, etc
// ------------------------------------------------------------------------------------------
func StartApp() error {
	// Loads environment variable from .env
	err := core.LoadConfig()
	if err != nil {
		return err
	}
	// Fetching config environment variables struct
	cfg := core.GlobalAppConfig
	// Connect to DB
	db, err := core.StartDB(cfg.DBString)
	if err != nil {
		log.Println(err)
	}

	// Passing db instance into internal for using db operations
	internal.SetDB(db)

	// Initializing Gin Router web Server
	routerInstance, RouterGroup := core.IntializeRouter()

	// Passing router group to routes
	router.AdminRoutes(RouterGroup)
	router.MerchantRoutes(RouterGroup)

	// Start the web server on port 8080
	log.Println("Starting the web server on port 8080")
	routerInstance.Run(":" + cfg.Port)

	return nil
}
