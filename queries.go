package main

import (
	"context"
	"database/sql"
	"fmt"
)

type ArticleQuery struct {
	db *sql.DB
}

func (q *ArticleQuery) Add(ctx context.Context, instance Article) (Article, error) {
	query := `
	 INSERT INTO articles(title, content, author) VALUES
	 (?, ?, ?)
	 RETURNING id, created, updated;
	`

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return Article{}, fmt.Errorf("Unable to add user: %v", err)
	}

	article := Article{
		Title:   instance.Title,
		Content: instance.Content,
		Author:  instance.Author,
	}
	if err := stmt.QueryRowContext(ctx, instance.Title, instance.Content, instance.Author).Scan(&article.ID, &article.Created, &article.Updated); err != nil {
		return Article{}, fmt.Errorf("Unable to add user: %v", err)
	}

	return article, err
}

func (q *ArticleQuery) Update(ctx context.Context, instance Article) (Article, error) {
	return Article{}, nil
}

func (q *ArticleQuery) List(ctx context.Context) ([]Article, error) {
	return nil, nil
}

func (q *ArticleQuery) GetById(ctx context.Context, id int) (Article, error) {
	return Article{}, nil
}

func (q *ArticleQuery) Remove(ctx context.Context, id int) error {
	return nil
}

func (q *ArticleQuery) Exists(ctx context.Context, id int) (bool, error) {
	return false, nil
}
