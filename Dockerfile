FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

FROM golang:1.24.4-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates

RUN wget -O migrate.tar.gz "https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz" && \
    tar -xzf migrate.tar.gz && \
    mv migrate /usr/local/bin/ && \
    rm migrate.tar.gz && \
    chmod +x /usr/local/bin/migrate

COPY --from=builder /app .

COPY internal/infrastructure/database/migrations ./internal/infrastructure/database/migrations

EXPOSE 8080

CMD ["go", "run", "./cmd/api"]