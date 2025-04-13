package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(ctx context.Context, driver string, datasource string) (*sql.DB, error) {
	pool, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, fmt.Errorf("Unable to use datsource %v: %v", datasource, err)
	}

	// Optionally configuring the pool
	pool.SetConnMaxLifetime(0)
	pool.SetConnMaxIdleTime(3)
	pool.SetMaxOpenConns(3)

	// Ping the datasource to make sure everything it is available
	if err := Ping(ctx, pool, datasource); err != nil {
		return nil, err
	}

	return pool, nil
}

func Ping(ctx context.Context, pool *sql.DB, datasource string) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.PingContext(ctx); err != nil {
		return fmt.Errorf("Unable to use datsource %v: %v", DataSource, err)
	}
	return nil
}
