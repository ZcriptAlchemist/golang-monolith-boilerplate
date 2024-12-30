package internal

import (
	"log"
	"sqlc-demo/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	DB     *sqlc.Queries
	TestDB *sqlc.Queries
)

// ----------------------------
// Mapping DB Instance to sqlc
// ----------------------------
func SetDB(database *pgxpool.Pool) {
	if database == nil {
		log.Println("Received nil database connection")
	}

	DB = sqlc.New(database)
}

// ------------------------------------------------
// Mapping Staging DB Instance for testing to sqlc
// ------------------------------------------------
func SetTestDB(database *pgxpool.Pool) {
	if database == nil {
		log.Println("Received nil database connection")
	}

	TestDB = sqlc.New(database)
}
