.PHONY: migrate run format install dbshell

run:
	@go run .

migrate:
	@go run scripts/migrate.go

format:
	@go fmt ./...

install:
	@go mod tidy

dbshell:
	@sqlite3 -table articles.db

build:
	@go build -o dist/ 
