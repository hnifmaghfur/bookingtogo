.PHONY: migrate-up migrate-down seed run-dev

# Database migration up
migrate-up:
	@echo "Running database migrations..."
	@go run cmd/db/main.go migrate up

# Database migration down
migrate-down:
	@echo "Reverting database migrations..."
	@go run cmd/db/main.go migrate down

# Run database seeder
seed:
	@echo "Seeding database..."
	@go run cmd/db/main.go seed

# Run application in development mode
run-dev:
	@echo "Starting development server..."
	@go run main.go

# Setup development environment (migrate + seed + run)
dev: migrate-up seed run-dev

# Clean up (migrate down + up)
reset-db: migrate-down migrate-up seed

help:
	@echo "Available commands:"
	@echo "  make migrate-up    - Run database migrations"
	@echo "  make migrate-down  - Rollback database migrations"
	@echo "  make seed         - Seed the database"
	@echo "  make run-dev      - Run the application in development mode"
	@echo "  make dev          - Setup and run development environment (migrate + seed + run)"
	@echo "  make reset-db     - Reset database (migrate down + up + seed)"
	@echo "  make help         - Show this help message"