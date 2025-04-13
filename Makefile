.PHONY: migrate run format install

run:
	@go run main.go

migrate:
	@go run scripts/migrate.go

format:
	@go fmt ./...

install:
	@go mod tidy
