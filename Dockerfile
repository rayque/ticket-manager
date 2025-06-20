# --- Estágio de Build ---
FROM golang:1.24.4 AS builder

WORKDIR /app

# Copia arquivos de dependências primeiro (para melhor cache)
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copia o código-fonte
COPY . .

# --- Estágio Final ---
FROM golang:1.24.4-alpine

WORKDIR /app

# Instala dependências mínimas necessárias
RUN apk add --no-cache ca-certificates

# Instala golang-migrate
RUN wget -O migrate.tar.gz "https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz" && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/local/bin/ && \
    rm migrate.tar.gz && \
    chmod +x /usr/local/bin/migrate

# Copia o código-fonte do estágio de build
COPY --from=builder /app .

# Copia as migrações para o container
COPY internal/infrastructure/database/migrations ./internal/infrastructure/database/migrations

#COPY .env .env


EXPOSE 8080

# Compila e executa diretamente
CMD ["go", "run", "./cmd/api"]