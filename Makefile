-include .env

# Determine which docker-compose command to use.
# If the new docker compose command is available it is used,
# otherwise, the legacy docker-compose command is used.
DOCKER_COMPOSE := $(shell if command -v docker >/dev/null 2>&1 && docker compose version >/dev/null 2>&1; then echo "docker compose"; else echo "docker-compose"; fi)

# Create the .env file if it does not exist by copying from .env.example.
.env:
	@if [ ! -f .env ]; then \
	  cp .env.example .env; \
	fi

# install: Set up the complete development environment.
# It creates the .env file, starts the containers, applies migrations, and seeds the database.
install: .env up migrate_up seed

# up: Start docker containers in detached mode, building images as necessary.
up:
	$(DOCKER_COMPOSE) up -d --build

# down: Stop and remove all running docker containers.
down:
	$(DOCKER_COMPOSE) down

# seed: Execute the seed command to populate the database.
seed:
	$(DOCKER_COMPOSE) exec app go run cmd/seed/main.go

# logs: Display the real-time logs of the app container.
logs:
	$(DOCKER_COMPOSE) logs -f app

# migrate_up: Apply database migrations (up).
migrate_up:
	$(DOCKER_COMPOSE) exec app migrate -path=internal/infrastructure/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

# migrate_down: Roll back the applied database migrations (down).
migrate_down:
	$(DOCKER_COMPOSE) exec app migrate -path=internal/infrastructure/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

# generate-mocks: Generate mocks for interfaces in the project.
generate-mocks:
	sh scripts/generate-mocks.sh

# Define directories for calculating test coverage.
COVER_DIRS=./internal/application/usecases

# test-command: Run tests with race detector and code coverage enabled.
test-command:
	go test -race -coverpkg=$(COVER_DIRS) -v -coverprofile=coverage.out ./...

# test: Execute tests and display the coverage report.
test:
	make test-command
	go tool cover -func=coverage.out

.PHONY: help
# help: Display all available commands in this Makefile.
help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-18s\033[0m %s\n", $$1, $$2}'