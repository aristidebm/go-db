.PHONY: migrate run format install dbshell generate

run: build
	@./dist/db -d articles.db

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

generate:
	@go tool jet -source=sqlite -dsn="./articles.db" -path=./gen -rel-model-path=./entity

