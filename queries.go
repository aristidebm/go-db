package main

import (
	"context"
	"database/sql"
	"errors"
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
	query := `
	 UPDATE articles
	 SET 
	    title = ?,
		content = ?,
		author = ?
	 WHERE id = ?
	 RETURNING id, created, updated;
	`

	exists, err := q.Exists(ctx, instance.ID)
	if err != nil {
		return Article{}, err
	}

	if !exists {
		err := errors.New("Unable to update")
		return Article{}, fmt.Errorf("%w: the article %v does not exists", err, instance.ID)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return Article{}, fmt.Errorf("Unable to update user: %v", err)
	}

	article := Article{
		Title:   instance.Title,
		Content: instance.Content,
		Author:  instance.Author,
	}
	if err := stmt.QueryRowContext(ctx, instance.Title, instance.Content, instance.Author, instance.ID).Scan(&article.ID, &article.Created, &article.Updated); err != nil {
		return Article{}, fmt.Errorf("Unable to update user: %v", err)
	}

	return article, err
}

func (q *ArticleQuery) List(ctx context.Context) ([]Article, error) {
	return nil, nil
}

func (q *ArticleQuery) GetById(ctx context.Context, id int) (Article, error) {
	query := `
	 SELECT id, title, content, author, created, updated
	 FROM articles
	 WHERE id = ?;
	`

	exists, err := q.Exists(ctx, id)
	if err != nil {
		return Article{}, err
	}

	if !exists {
		err := errors.New("Unable to update")
		return Article{}, fmt.Errorf("%w: the retrieve the article %v does not exists", err, id)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return Article{}, fmt.Errorf("Unable to retrieve user: %v", err)
	}

	article := Article{}
	if err := stmt.QueryRowContext(ctx, id).Scan(&article.ID, &article.Title, &article.Content, &article.Author, &article.Created, &article.Updated); err != nil {
		return Article{}, fmt.Errorf("Unable to retrieve user: %v", err)
	}

	return article, err
}

func (q *ArticleQuery) Remove(ctx context.Context, id int) error {
	query := `
	 DELETE FROM articles
	 WHERE id = ?;
	`

	exists, err := q.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		err := errors.New("Unable to remove")
		return fmt.Errorf("%w: the article %v does not exists", err, id)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to remove user: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("Unable to remove user: %v", err)
	}

	return nil
}

func (q *ArticleQuery) Exists(ctx context.Context, id int) (bool, error) {
	query := `
	SELECT EXISTS(SELECT 1 FROM articles WHERE id = ?);
	`
	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	var exists bool
	if err := stmt.QueryRowContext(ctx, id).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}
