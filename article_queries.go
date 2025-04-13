package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type ArticleQuery struct {
	db       *sql.DB
	comments *CommentQuery
}

func NewArticleQuery(db *sql.DB) *ArticleQuery {
	articleQuery := &ArticleQuery{
		db: db,
	}
	articleQuery.comments = &CommentQuery{
		db:       db,
		articles: articleQuery,
	}
	return articleQuery
}

func (q *ArticleQuery) Add(ctx context.Context, instance *Article) error {
	query := `
	 INSERT INTO articles(title, content, author) VALUES
	 (?, ?, ?)
	 RETURNING id, created, updated;
	`

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to add article: %v", err)
	}

	if err := stmt.QueryRowContext(ctx, instance.Title, instance.Content, instance.Author).Scan(&instance.ID, &instance.Created, &instance.Updated); err != nil {
		return fmt.Errorf("Unable to add article: %v", err)
	}

	return nil
}

func (q *ArticleQuery) Update(ctx context.Context, instance *Article) error {
	query := `
	 UPDATE articles
	 SET 
	    title = ?,
		content = ?,
	 WHERE id = ?
	 RETURNING id, created, updated;
	`

	exists, err := q.Exists(ctx, instance.ID)
	if err != nil {
		return err
	}

	if !exists {
		err := errors.New("Unable to update")
		return fmt.Errorf("%w: the article %v does not exists", err, instance.ID)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to update article: %v", err)
	}

	if err := stmt.QueryRowContext(ctx, instance.Title, instance.Content, instance.ID).Scan(&instance.ID, &instance.Created, &instance.Updated); err != nil {
		return fmt.Errorf("Unable to update article: %v", err)
	}

	return nil
}

func (q *ArticleQuery) List(ctx context.Context) ([]*Article, error) {
	query := `
	 SELECT id, title, content, author, created, updated
	 FROM articles;
	`
	articles := []*Article{}

	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Unable to list articles: %v", err)
	}

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Author, &article.Created, &article.Updated)
		if err != nil {
			break
		}
		articles = append(articles, article)
	}

	// Check for errors during rows "Close".
	// This may be more important if multiple statements are executed
	// in a single batch and rows were written as well as read.
	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	// Check for row scan error.
	if err != nil {
		return nil, err
	}

	// Check for errors during row iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (q *ArticleQuery) GetById(ctx context.Context, id int) (*Article, error) {
	query := `
	 SELECT id, title, content, author, created, updated
	 FROM articles
	 WHERE id = ?;
	`

	exists, err := q.Exists(ctx, id)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := errors.New("Unable to update")
		return nil, fmt.Errorf("%w: the retrieve the article %v does not exists", err, id)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve article: %v", err)
	}

	article := &Article{}
	if err := stmt.QueryRowContext(ctx, id).Scan(&article.ID, &article.Title, &article.Content, &article.Author, &article.Created, &article.Updated); err != nil {
		return nil, fmt.Errorf("Unable to retrieve article: %v", err)
	}

	return article, nil
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
		return fmt.Errorf("Unable to remove article: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("Unable to remove article: %v", err)
	}

	return nil
}

func (q *ArticleQuery) ListWithComments(ctx context.Context) ([]*Article, error) {
	articles, err := q.List(ctx)
	if err != nil {
		return nil, err
	}

	articleIds := []int{}
	for _, artcile := range articles {
		articleIds = append(articleIds, artcile.ID)
	}

	comments, err := q.comments.List(ctx, articleIds...)
	if err != nil {
		return nil, err
	}

	commentsPerArticle := map[int][]Comment{}
	for _, comment := range comments {
		commentsPerArticle[comment.Article] = append(commentsPerArticle[comment.Article], *comment)
	}

	for _, article := range articles {
		article.Comments = append(article.Comments, commentsPerArticle[article.ID]...)
	}

	return articles, nil
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
