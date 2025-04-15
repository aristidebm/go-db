package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"example/db/gen/entity"
	"example/db/gen/table"

	"github.com/go-jet/jet/v2/sqlite"
)

type CommentQuery struct {
	db       *sql.DB
	articles *ArticleQuery
}

func NewCommentQuery(db *sql.DB) *CommentQuery {
	CommentQuery := &CommentQuery{
		db: db,
	}
	CommentQuery.articles = &ArticleQuery{
		db:       db,
		comments: CommentQuery,
	}
	return CommentQuery
}

func (q *CommentQuery) Add(ctx context.Context, instance *Comment) error {
	query := `
	 INSERT INTO comments(content, author, article_id) VALUES
	 (?, ?, ?)
	 RETURNING id, created, updated;
	`

	exists, err := q.articles.Exists(ctx, instance.Article)
	if err != nil {
		return err
	}

	if !exists {
		err := errors.New("Unable to update")
		return fmt.Errorf("%w: the article %v does not exists", err, instance.Article)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to add comment: %v", err)
	}

	if err := stmt.QueryRowContext(ctx, instance.Content, instance.Author, instance.Article).Scan(&instance.ID, &instance.Created, &instance.Updated); err != nil {
		return fmt.Errorf("Unable to add comment: %v", err)
	}

	return nil
}

func (q *CommentQuery) Update(ctx context.Context, instance *Comment) error {
	query := `
	 UPDATE comments
	 SET content = ?,
	 WHERE id = ?
	 RETURNING id, created, updated;
	`

	exists, err := q.Exists(ctx, instance.ID)
	if err != nil {
		return err
	}

	if !exists {
		err := errors.New("Unable to update")
		return fmt.Errorf("%w: the comment %v does not exists", err, instance.ID)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to update comment: %v", err)
	}

	if err := stmt.QueryRowContext(ctx, instance.Content, instance.ID).Scan(&instance.ID, &instance.Created, &instance.Updated); err != nil {
		return fmt.Errorf("Unable to update comment: %v", err)
	}

	return nil
}

func (q *CommentQuery) List(ctx context.Context, articles ...int) ([]*Comment, error) {
	query := `
	 SELECT id, content, author, article_id, created, updated
	 FROM comments
	`

	// construct dynamic filter on articles
	placeholders := ""
	if placeholders = strings.Repeat("?,", len(articles)); placeholders != "" {
		placeholders = "WHERE article_id IN ( " + strings.TrimRight(placeholders, ",") + ")"
	}
	query = query + placeholders + ";"

	comments := []*Comment{}

	// NOTE: Convert []int to []interface{}
	// for a reason, that I don't know is not possible.
	// Either to pass a slice of []int or to cast this.
	// slice as []any([]int), perhaps because only down-casting
	// is supported, I don't know, I have to look at it

	args := make([]interface{}, len(articles))
	for i, v := range articles {
		args[i] = v
	}

	rows, err := q.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Unable to list comments: %v", err)
	}

	for rows.Next() {
		comment := &Comment{}
		err := rows.Scan(&comment.ID, &comment.Content, &comment.Author, &comment.Article, &comment.Created, &comment.Updated)
		if err != nil {
			break
		}
		comments = append(comments, comment)
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

	return comments, nil
}

func (q *CommentQuery) GetById(ctx context.Context, id int) (*Comment, error) {
	query := `
	 SELECT id, content, author, created, updated
	 FROM comments
	 WHERE id = ?;
	`

	exists, err := q.Exists(ctx, id)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := errors.New("Unable to update")
		return nil, fmt.Errorf("%w: the retrieve the comment %v does not exists", err, id)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve comment: %v", err)
	}

	article := &Comment{}
	if err := stmt.QueryRowContext(ctx, id).Scan(&article.ID, &article.Content, &article.Author, &article.Created, &article.Updated); err != nil {
		return nil, fmt.Errorf("Unable to retrieve comment: %v", err)
	}

	return article, nil
}

func (q *CommentQuery) Remove(ctx context.Context, id int) error {
	query := `
	 DELETE FROM comments
	 WHERE id = ?;
	`

	exists, err := q.Exists(ctx, id)
	if err != nil {
		return err
	}

	if !exists {
		err := errors.New("Unable to remove")
		return fmt.Errorf("%w: the comment %v does not exists", err, id)
	}

	stmt, err := q.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("Unable to remove comment: %v", err)
	}

	if _, err := stmt.ExecContext(ctx, id); err != nil {
		return fmt.Errorf("Unable to remove comment: %v", err)
	}

	return nil
}

func (q *CommentQuery) Exists(ctx context.Context, id int) (bool, error) {
	query := `
	SELECT EXISTS(SELECT 1 FROM comments WHERE id = ?);
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

type filterOption struct {
	limit     int
	offset    int
	articleId int
	name      string
}

type FilterOption func(option *filterOption)

func WithLimit(limit int) FilterOption {
	return func(option *filterOption) {
		option.limit = limit
	}
}

func WithOffset(offset int) FilterOption {
	return func(option *filterOption) {
		option.offset = offset
	}
}

func WithArticleId(articleId int) FilterOption {
	return func(option *filterOption) {
		option.articleId = articleId
	}
}

func (q *CommentQuery) Filter(ctx context.Context, options ...FilterOption) ([]entity.Comments, error) {
	option := &filterOption{}
	for _, fn := range options {
		fn(option)
	}
	fmt.Print(option)

	stmt := sqlite.SELECT(table.Comments.AllColumns).FROM(table.Comments)

	if option.limit != 0 {
		stmt = stmt.LIMIT(int64(option.limit))
	}

	if option.offset != 0 {
		stmt = stmt.OFFSET(int64(option.offset))
	}

	if option.articleId != 0 {
		stmt = stmt.WHERE(table.Comments.ArticleID.EQ(sqlite.Int64(int64(option.articleId))))
	}

	comments := []entity.Comments{}
	if err := stmt.QueryContext(ctx, q.db, &comments); err != nil {
		return nil, err
	}
	fmt.Printf("%#+v", comments)
	return comments, nil
}
