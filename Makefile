help:
	@printf "usage: make [target] ...\n\n"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

# ---------------------------------
# Common project related commands.
# ---------------------------------

install: ## Install external dependencies and resources.
	yarn
	make -f ./apps/service/Makefile install

create_resources: ## Create the local resources, such as data tables, buckets, etc.
	make -f ./apps/service/Makefile create_resources

migrate_databases: ## Migrate databases of all applications
	make -f ./apps/service/Makefile migrate_database

run_infra: ## Run the infrastructure locally
	docker-compose -f docker-compose.yml up -d
