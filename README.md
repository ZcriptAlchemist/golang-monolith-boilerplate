# golang-monolith-boilerplate

This is boilerplate code for a monolith backend application

### packages used
- [Gin](https://gin-gonic.com/) for http router
- [goose](https://github.com/pressly/goose) for database migrations
- [sqlc](https://sqlc.dev/) for generating functions that directly execute your SQL queries
- [pgx](https://pkg.go.dev/github.com/jackc/pgx/v5) as postgres driver
- [air](https://github.com/air-verse/air) for live reloading