package core

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sqlc-demo/internal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// --------------------------------------------------------------------------------
// StartDB initializes a connection to the database and returns a connection pool.
// --------------------------------------------------------------------------------
func StartDB(dbString string) (*pgxpool.Pool, error) {
	// Create a connection pool
	dbpool, err := pgxpool.New(context.Background(), dbString)
	fmt.Println(&dbpool)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	log.Println("Connected to the database successfully!")
	return dbpool, nil
}

// --------------------------------------------------------------------------------------------
// StartTestDB initializes a connection to the staging database and returns a connection pool.
// --------------------------------------------------------------------------------------------
func StartTestDB() (*pgxpool.Pool, error) {
	// Load environment variables from a file
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
		return nil, errors.New("failed to load .env file")
	}

	// Retrieve and check the DB Connection String variable
	testDBString := os.Getenv("STAGING_DATABASE_URL")
	if testDBString == "" {
		return nil, errors.New("missing staging DB Connection string environment variable")
	}

	// Create a connection pool
	dbpool, err := pgxpool.New(context.Background(), testDBString)
	fmt.Println(&dbpool)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to test database: %w", err)
	}

	// mapping db instance to sqlc
	internal.SetTestDB(dbpool)

	log.Println("Connected to test database successfully! Now you can use test db with `TestDB`")
	return dbpool, nil
}
