package dbSetup

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
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
