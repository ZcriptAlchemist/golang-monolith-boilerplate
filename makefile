hello:
	echo "testing makefile bish!"

# // runs app locally in the terminal
# run:
# 	cd cmd && go run .

# // builds and runs app locally in the terminal with hot module reload
run:
	air

# // builds docker image
dockerBuild:
	docker build -t server-image .

# // runs built docker image in a container
dockerRun: dockerBuild
	docker run -d -p 8080:8080 --name server-container server-image

# Build development mode in Docker with hot module reload
dockerHotBuild:
	docker build -t server-dev --target development .

# Run development mode in Docker with hot module reload
dockerHotRun: dockerHotBuild
	docker run --rm -p 8080:8080 --name server-container-dev -v $(shell pwd):/server server-dev

# Build production mode Docker image
dockerProdBuild:
	docker build -t server-prod --target production .

# Run production mode Docker container in detached mode
dockerProdRun: dockerProdBuild
	docker run -d --name server-prod -p 8080:8080 server-prod

# // Generates sqlc code for db operations
sqlcGenerate:
	sqlc generate

# // Migrates schema changes to DB
# // note - these goose commands fetch directory path, driver, DB connection string from .env. For more customized migrations use custom goose commands as per your needs
gooseUp:
	source .env && goose -dir $$GOOSE_MIGRATION_DIR $$GOOSE_DRIVER "$$GOOSE_DBSTRING" up

# // rollbacks schema changes to DB
# // note - these goose commands fetch directory path, driver, DB connection string from .env. For more customized migrations use custom goose commands as per your needs
gooseDown:
	source .env && goose -dir $$GOOSE_MIGRATION_DIR $$GOOSE_DRIVER "$$GOOSE_DBSTRING" down
