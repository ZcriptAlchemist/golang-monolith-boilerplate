package internal

import (
	"log"
	"sqlc-demo/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *sqlc.Queries

// --------------------
// Getting DB Instance
// --------------------
func SetDB(database *pgxpool.Pool) {
	if database == nil {
		log.Println("Received nil database connection")
	}

	DB = sqlc.New(database)
}
