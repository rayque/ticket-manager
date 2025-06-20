# Verificar se .env existe, se não, criar a partir do .env.example
-include .env

# Detectar automaticamente qual versão do Docker Compose está disponível
DOCKER_COMPOSE := $(shell if command -v docker >/dev/null 2>&1 && docker compose version >/dev/null 2>&1; then echo "docker compose"; else echo "docker-compose"; fi)

.env:
	@if [ ! -f .env ]; then \
		echo "Arquivo .env não encontrado. Criando a partir do .env.example..."; \
		cp .env.example .env; \
		echo "Arquivo .env criado! Por favor, configure as variáveis conforme necessário."; \
	fi

install: .env up migrate_up up_seed

up:
	$(DOCKER_COMPOSE) up -d --build

down:
	$(DOCKER_COMPOSE) down

seed:
	$(DOCKER_COMPOSE) exec app go run cmd/seed/main.go

up_seed: up
	@echo "Waiting for MongoDB to be ready..."
	@sleep 3
	@make seed

migrate_up:
	$(DOCKER_COMPOSE) exec app migrate -path=internal/infrastructure/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	$(DOCKER_COMPOSE) exec app migrate -path=internal/infrastructure/database/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down

generate-mocks:
	sh scripts/generate-mocks.sh

COVER_DIRS=./internal/application/usecases

test-command:
	go test -race -coverpkg=$(COVER_DIRS) -v -coverprofile=coverage.out ./...

test:
	make test-command
	go tool cover -func=coverage.out

logs:
	$(DOCKER_COMPOSE) logs -f app