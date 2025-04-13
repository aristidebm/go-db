package main

import (
	"context"
	"fmt"
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
	// cancel the context on system interrupt signal
	// reception
	go func() {
		<-sigChan
		cancel()
	}()

	// get a connection pool
	pool, err := InitDB(ctx, "sqlite3", DataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// migrate tables
	if err := CreateTables(ctx, pool); err != nil {
		log.Fatal(err)
	}

	// instantiate article query
	articles := ArticleQuery{db: pool}
	article, err := articles.Add(ctx, Article{
		Title: "How to learn golang in 30 days ?",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(article)
}
