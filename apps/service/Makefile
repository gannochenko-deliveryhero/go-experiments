help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Application related commands.
# ---------------------------------

install:
	go mod download

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	echo "Create!"

migrate_database: ## Migrate the database
