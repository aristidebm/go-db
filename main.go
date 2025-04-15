package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
)

func main() {

	// parse command line arguments
	var dataSource string
	flag.StringVar(&dataSource, "datsource", "", "SQLite datasource")
	flag.StringVar(&dataSource, "d", "", "SQLite datasource")
	flag.Parse()

	if dataSource == "" {
		log.Fatal("")
	}

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
	pool, err := InitDB(ctx, "sqlite3", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// migrate tables
	if err := CreateTables(ctx, pool); err != nil {
		log.Fatal(err)
	}

	// instantiate queries
	// articles := NewArticleQuery(pool)
	comments := NewCommentQuery(pool)

	// // // add an article
	// article, err := articles.Add(ctx, Article{
	// 	Title: "How to learn golang in 30 days ?",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = comments.Add(ctx, Comment{
	// 	Content: "Very interesting post",
	// 	Article: article.ID,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// err = comments.Add(ctx, Comment{
	// 	Content: "Nicely done",
	// 	Article: article.ID,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// artciles, err := articles.ListWithComments(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// JSONResponse(os.Stdout, artciles)

	// results, err := comments.Filter(ctx, WithLimit(1))
	results, err := comments.Filter(ctx, WithLimit(2), WithArticleId(100))
	if err != nil {
		log.Fatal("cannot return a result")
	}

	fmt.Print(results)

}

func JSONResponse(w io.Writer, data any) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return err
	}
	io.Copy(w, &buf)
	return nil
}
