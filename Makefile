.PHONY: migrate run format install

run:
	@go run .

migrate:
	@go run scripts/migrate.go

format:
	@go fmt ./...

install:
	@go mod tidy
