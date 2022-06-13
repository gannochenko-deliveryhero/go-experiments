ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Application related commands.
# ---------------------------------

install:
	@go install github.com/joho/godotenv/cmd/godotenv@latest
	@go install github.com/cosmtrek/air@latest
	@brew install golang-migrate
	@brew install postgresql
	@brew install kyoh86/tap/richgo
	@go mod download

create_resources: ## Create local resources, such as data tables, buckets, etc.

migrate_database: ## Migrate the database

run: ## Run the application locally on the host machine
	@godotenv -f $(ROOT_DIR)/../../.env.local go run $(ROOT_DIR)/cmd/service.go

run_live:
	@godotenv -f $(ROOT_DIR)/../../.env.local air
