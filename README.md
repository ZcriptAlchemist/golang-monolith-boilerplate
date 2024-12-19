# golang-monolith-boilerplate

This is boilerplate code for a monolith backend application

### packages used
- [Gin](https://gin-gonic.com/) for http router
- [goose](https://github.com/pressly/goose) for database migrations
- [sqlc](https://sqlc.dev/) for generating functions that directly execute your SQL queries
- [pgx](https://pkg.go.dev/github.com/jackc/pgx/v5) as postgres driver
- [air](https://github.com/air-verse/air) for live reloading

This setup was done on wsl ubuntu 24.04.1 LTS

### steps to run 

- clone this repository
- install all dependencies - note that some dependencies like sqlc, goose can be installed on cli
```bash
go mod tidy
```
- run command - make sure to have installed `air` for hot reload
```bash
air
```
### Goose CLI Commands

- Installing Goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Load Environment Variables
```bash
source .env
```

- Create a Migration script
```bash
goose -dir $GOOSE_MIGRATIONS_DIR create file_name sql
```
- Run Migrations (Up)
```bash
goose -dir $GOOSE_MIGRATIONS_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" up
```
- Rollback Migrations (Down)
```bash
goose -dir $GOOSE_MIGRATIONS_DIR $GOOSE_DRIVER "$GOOSE_DBSTRING" down
```
