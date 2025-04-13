package main

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateTables(ctx context.Context, pool *sql.DB) error {
	stmt := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255),
		content TEXT NULL,
		author VARCHAR(50) NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		article_id INTEGER
		content TEXT NULL,
		author VARCHAR(50) NULL,
		created DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(article_id) REFERENCES articles(id)
	);

	CREATE TRIGGER IF NOT EXISTS update_articles_updated
	AFTER UPDATE ON articles
	WHEN old.updated <> CURRENT_TIMESTAMP
	BEGIN
		 UPDATE articles
		SET updated = CURRENT_TIMESTAMP
		WHERE id = old.id;
	END;

	CREATE TRIGGER IF NOT EXISTS update_comments_updated
	AFTER UPDATE ON comments
	WHEN old.updated <> CURRENT_TIMESTAMP
	BEGIN
		 UPDATE comments
		SET updated = CURRENT_TIMESTAMP
		WHERE id = old.id;
	END;
	`

	if _, err := pool.ExecContext(ctx, stmt); err != nil {
		return fmt.Errorf("Unable to create tables %v", err)
	}

	return nil
}
