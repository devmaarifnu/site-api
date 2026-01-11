.PHONY: help run build clean test install dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	go mod download
	go mod tidy

run: ## Run the application
	go run cmd/api/main.go

dev: ## Run the application in development mode with hot reload (requires air)
	air

build: ## Build the application
	go build -o bin/api cmd/api/main.go

build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 go build -o bin/api-linux cmd/api/main.go

build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 go build -o bin/api.exe cmd/api/main.go

test: ## Run tests
	go test -v ./...

clean: ## Clean build files
	rm -rf bin/
	rm -rf tmp/
	go clean

fmt: ## Format code
	go fmt ./...

lint: ## Run linter
	golangci-lint run

docker-build: ## Build Docker image
	docker build -t lpmaarifnu-site-api .

docker-run: ## Run Docker container
	docker run -p 8080:8080 lpmaarifnu-site-api

setup-db: ## Setup database from schema
	mysql -u root -p < database_schema.sql
