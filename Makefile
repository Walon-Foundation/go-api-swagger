.PHONY: swag dev test format migrate-up migrate-down

#This is a demo DATABASE_URL	
DB_URL=postgres://walon:password@localhost:5432/app_db?sslmode=disable
MIGRATION_PATH=cmd/db/migrations

#

swag:
	swag init -g main.go -d cmd/api,cmd/utils,cmd/models,cmd/db
	
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
	
docker-build:
	docker build -t my-api:1.0 .

docker-run:
	docker run -d --name my-api -p 8080:8080 my-api:1.0

