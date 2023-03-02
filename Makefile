services-up:
	docker compose -f docker-compose.dev.yml up

services-down:
	docker compose -f docker-compose.dev.yml down

migrateup:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/test_task_6?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/test_task_6?sslmode=disable" -verbose down

up:
	docker compose up

down:
	docker compose down -v

sqlc:
	sqlc generate

server:
	go run ./cmd/main.go

.PHONY: services-up services-down  migrate-create migrate-up sqlc up down server