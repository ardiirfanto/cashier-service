.PHONY: help run build test clean install db-setup

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

run: ## Run the application
	go run ./cmd/server

build: ## Build the application
	go build -o bin/cashier-api ./cmd/server

test: ## Run tests
	go test -v ./...

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf tmp/

db-setup: ## Setup database (requires MySQL to be running)
	@echo "Setting up database..."
	@mysql -u root -p < database_setup.sql
	@echo "Database setup completed!"

dev: ## Run in development mode with auto-reload (requires air: go install github.com/cosmtrek/air@latest)
	air

lint: ## Run linter (requires golangci-lint)
	golangci-lint run

format: ## Format code
	go fmt ./...

docker-build: ## Build Docker image
	docker build -t cashier-api:latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 --env-file .env cashier-api:latest
