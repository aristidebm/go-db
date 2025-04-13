package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

var DataSource = "articles.db"

func main() {
	// connection establishment
	sigChan := make(chan os.Signal, 3)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(context.Background())
	// Cancel the context on system interrupt signal
	// reception
	go func() {
		<-sigChan
		cancel()
	}()

	if err := InitDB(ctx, "sqlite3", DataSource); err != nil {
		log.Fatal(err)
	}
}
