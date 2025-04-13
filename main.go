package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DataSource = "articles.db"

func main() {
	// connection establishment
	pool, err := sql.Open("sqlite3", DataSource)
	if err != nil {
		log.Fatalf("Unable to use datsource %v: %v", DataSource, err)
	}
	defer pool.Close()

	// optionally configuring the pool
	pool.SetConnMaxLifetime(0)
	pool.SetConnMaxIdleTime(3)
	pool.SetMaxOpenConns(3)

	// setup a cancellation context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Ping the datasource to make sure everything it is available
	if err := Ping(ctx, pool); err != nil {
		log.Fatal(err)
	}
}

func Ping(ctx context.Context, pool *sql.DB) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		return fmt.Errorf("Unable to use datsource %v: %v", DataSource, err)
	}
	return nil
}
