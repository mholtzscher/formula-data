DB_DRIVER := postgres
DB_STRING := "host=localhost user=postgres password=postgres dbname=formula-data sslmode=disable"

# Define the directory where migration files are stored
MIGRATIONS_DIR := sql/migrations

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make create-migration name=<migration_name>"
	@echo "  make up"
	@echo "  make down"
	@echo "  make status"
	@echo "  make gen"

# Create a new migration
.PHONY: create-migration
create-migration:
	@if [ -z "$(name)" ]; then echo "Please provide a migration name like this: make create-migration name=your_migration"; exit 1; fi
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Migrate the DB to the most recent version available
.PHONY: up
up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) up

# Roll back the version by 1
.PHONY: down
down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) down

# Display the status of all migrations
.PHONY: status
status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) status

# Generate the code for the SQLC
.PHONY: gen
gen:
	sqlc generate
