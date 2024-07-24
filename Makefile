DB_DRIVER := postgres
DB_STRING := "host=localhost user=postgres password=postgres dbname=formula-data sslmode=disable"

# Define the directory where migration files are stored
MIGRATIONS_DIR := sql/migrations

# Default target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make migration-create name=<migration_name>"
	@echo "  make migration-up"
	@echo "  make migration-down"
	@echo "  make migration-status"
	@echo "  make sqlc-gen"
	@echo "  make mock-gen"
	@echo "  make buf-gen"
	@echo "  make run"

# Create a new migration
.PHONY: migration-create
migration-create:
	@if [ -z "$(name)" ]; then echo "Please provide a migration name like this: make create-migration name=your_migration"; exit 1; fi
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Migrate the DB to the most recent version available
.PHONY: migration-up
migration-up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) up

# Roll back the version by 1
.PHONY: migration-down
migration-down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) down

# Display the status of all migrations
.PHONY: migration-status
migration-status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) $(DB_STRING) status

# Run the server
.PHONY: run
run: 
	go run cmd/server/server.go

# Generate the code for the sqlc
.PHONY: sqlc-gen
sqlc-gen:
	sqlc generate

# Generate code from the proto files
.PHONY: buf-gen
buf-gen:
	buf generate

# Generate mocks for testing
.PHONY: mock-gen
mock-gen:
	mockery

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run all gen commands
.PHONY: gen
gen: buf-gen migration-up sqlc-gen mock-gen 

# Run all 
.PHONY: all
all: buf-gen migration-up sqlc-gen mock-gen run


.PHONY: get-f1db-release
get-f1db-release:
	@if [ -z "$(tag)" ]; then echo "Please provide a release tag name like this: make get-f1db-release tag=your_tag"; exit 1; fi
	rm -rf /tmp/f1db-$(tag)
	mkdir -p /tmp/f1db-$(tag)
	curl -Lo /tmp/f1db-$(tag)/f1db-sql-sqlite.zip https://github.com/f1db/f1db/releases/download/$(tag)/f1db-sql-sqlite.zip
	unzip /tmp/f1db-$(tag)/f1db-sql-sqlite.zip -d /tmp/f1db-$(tag)/
	$(eval CLEANED_TAG := $(shell echo $(tag) | tr -d 'a-zA-Z-.'))
	sed '/^INSERT/d' /tmp/f1db-$(tag)/f1db-sql-sqlite.sql > ./sql/schema/$(CLEANED_TAG)-f1db-sql-sqlite.sql	
	curl -Lo /tmp/f1db-$(tag)/f1db-sqlite.zip https://github.com/f1db/f1db/releases/download/$(tag)/f1db-sqlite.zip
	unzip /tmp/f1db-$(tag)/f1db-sqlite.zip -d /tmp/f1db-$(tag)/
	cp /tmp/f1db-$(tag)/f1db.db ./f1db.db

