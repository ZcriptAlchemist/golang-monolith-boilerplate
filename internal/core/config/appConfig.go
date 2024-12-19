package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ----------------------------------------
// Struct with environment variable fields
// ----------------------------------------
type AppConfig struct {
	Port     string
	DBString string
}

// ------------------------------------
// Globally accessible config variable
// ------------------------------------
var GlobalAppConfig *AppConfig

// --------------------------------------
// Loads environment variables from .env
// --------------------------------------
func LoadConfig() error {
	// envPath := GetEnvPath()
	// Load environment variables from a file
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		return errors.New("failed to load .env file")
	}

	// Retrieve and check the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("missing PORT environment variable")
	}

	// Retrieve and check the DB Connection String variable
	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		return errors.New("missing DB Connection string environment variable")
	}

	// Mapping fetched environment variables from .env to GlobalAppConfig
	GlobalAppConfig = &AppConfig{
		Port:     port,
		DBString: dbString,
	}

	return nil
}

// // -------------------------
// // Getting global .env path
// // -------------------------
// func GetEnvPath() string {
// 	// Get the directory of the executable
// 	execPath, err := os.Executable()
// 	if err != nil {
// 		log.Println("Error getting executable path:", err)
// 		return ""
// 	}

// 	// Construct the path to .env in the same directory as the binary
// 	execDir := filepath.Dir(execPath)
// 	envPath := filepath.Join(execDir, ".env")

// 	// Verify if the .env file exists
// 	if _, err := os.Stat(envPath); err != nil {
// 		log.Println("Error: .env file not found at:", envPath, "Error:", err)
// 		return ""
// 	}

// 	log.Println("Resolved .env path:", envPath)
// 	return envPath
// }
