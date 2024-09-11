# Project names
PROJECT_LOCAL := sn-local
PROJECT_PROD := sn-prod

# Colors for pretty output
GREEN := \033[0;32m
NC := \033[0m # No Color

# Get all targets
TARGETS := $(shell grep -oE '^[a-zA-Z0-9_-]+:' $(MAKEFILE_LIST) | sed 's/://')

.PHONY: $(TARGETS)

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(GREEN)%-30s$(NC) %s\n", $$1, $$2}'

up: ## Start the development environment
	@echo "Starting development environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_LOCAL) docker compose up --build -d

up-setup: ## Start the development environment with initial setup
	@echo "Starting development environment with setup..."
	COMPOSE_PROJECT_NAME=$(PROJECT_LOCAL) docker compose --profile setup up --build -d

down: ## Stop the development environment
	@echo "Stopping development environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_LOCAL) docker compose down

prod-up: ## Start the production environment
	@echo "Starting production environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_PROD) docker compose -f docker-compose.prod.yml up --build -d

prod-up-setup: ## Start the production environment with initial setup
	@echo "Starting production environment with setup..."
	COMPOSE_PROJECT_NAME=$(PROJECT_PROD) docker compose -f docker-compose.prod.yml --profile setup up --build -d

prod-down: ## Stop the production environment
	@echo "Stopping production environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_PROD) docker compose -f docker-compose.prod.yml down

logs: ## Show logs from all containers (development environment)
	@echo "Showing logs for development environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_LOCAL) docker compose logs -f

prod-logs: ## Show logs from all containers (production environment)
	@echo "Showing logs for production environment..."
	COMPOSE_PROJECT_NAME=$(PROJECT_PROD) docker compose -f docker-compose.prod.yml logs -f

clean-postgres-volume-dev: ## Stop and remove PostgreSQL volume (development)
	@echo "Removing PostgreSQL volume for development environment..."
	docker volume rm $(PROJECT_LOCAL)_postgres_data || true

clean-postgres-volume-prod: ## Stop and remove PostgreSQL volume (production)
	@echo "Removing PostgreSQL volume for production environment..."
	docker volume rm $(PROJECT_PROD)_postgres_social_data || true

clean-minio-volume-dev: ## Stop and remove PostgreSQL volume (development)
	@echo "Removing MinIO volume for development environment..."
	docker volume rm $(PROJECT_LOCAL)_minio_data || true

clean-minio-volume-prod: ## Stop and remove PostgreSQL volume (production)
	@echo "Removing MinIO volume for production environment..."
	docker volume rm $(PROJECT_PROD)_minio_data || true

# Autocomplete setup
.PHONY: _autocomplete
_autocomplete:
	@echo $(TARGETS)

