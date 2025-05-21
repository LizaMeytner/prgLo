PROJECT_NAME = prgLo
GO = go
PROTOC = protoc

# Цели (targets)
.PHONY: all build clean test migrate-up migrate-down gen-proto run

all: build

## Сборка проекта
build:
	@echo "Building all services..."
	@$(GO) build -o bin/user-service ./cmd/user-service
	@$(GO) build -o bin/forum-service ./cmd/forum-service
	@$(GO) build -o bin/chat-service ./cmd/chat-service

## Запуск тестов
test:
	@echo "Running tests..."
	@$(GO) test -v -cover -race ./...

## Генерация gRPC кода
gen-proto:
	@echo "Generating protobuf files..."
	@$(PROTOC) --go_out=. --go-grpc_out=. ./api/proto/*.proto

## Миграции базы данных
migrate-up:
	@echo "Applying migrations..."
	@migrate -path ./migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up

migrate-down:
	@echo "Reverting migrations..."
	@migrate -path ./migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" down 1

## Запуск сервисов
run-user:
	@$(GO) run ./cmd/user-service

run-forum:
	@$(GO) run ./cmd/forum-service

run-chat:
	@$(GO) run ./cmd/chat-service

## Очистка
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@$(GO) clean