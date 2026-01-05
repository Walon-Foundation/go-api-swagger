.PHONY: swag dev test format migrate-up migrate-down

#This is a demo DATABASE_URL	
DB_URL=postgres://walon:password@localhost:5432/app_db?sslmode=disable
MIGRATION_PATH=cmd/db/migrations

swag:
	swag init -g cmd/api/main.go --parseDependency --parseInternal

dev:
	go run ./cmd/api

test:
	go test ./...

format:
	go fmt ./...
	
migrate-up:
	migrate -path "$(MIGRATION_PATH)" -database "$(DB_URL)" up

migrate-down:
	migrate -path "$(MIGRATION_PATH)" -database "$(DB_URL)" down 1
	
	
