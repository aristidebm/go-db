//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var Comments = newCommentsTable("", "comments", "")

type commentsTable struct {
	sqlite.Table

	// Columns
	ID        sqlite.ColumnInteger
	ArticleID sqlite.ColumnInteger
	Content   sqlite.ColumnString
	Author    sqlite.ColumnString
	Created   sqlite.ColumnTimestamp
	Updated   sqlite.ColumnTimestamp

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
	DefaultColumns sqlite.ColumnList
}

type CommentsTable struct {
	commentsTable

	EXCLUDED commentsTable
}

// AS creates new CommentsTable with assigned alias
func (a CommentsTable) AS(alias string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CommentsTable with assigned schema name
func (a CommentsTable) FromSchema(schemaName string) *CommentsTable {
	return newCommentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CommentsTable with assigned table prefix
func (a CommentsTable) WithPrefix(prefix string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CommentsTable with assigned table suffix
func (a CommentsTable) WithSuffix(suffix string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCommentsTable(schemaName, tableName, alias string) *CommentsTable {
	return &CommentsTable{
		commentsTable: newCommentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newCommentsTableImpl("", "excluded", ""),
	}
}

func newCommentsTableImpl(schemaName, tableName, alias string) commentsTable {
	var (
		IDColumn        = sqlite.IntegerColumn("id")
		ArticleIDColumn = sqlite.IntegerColumn("article_id")
		ContentColumn   = sqlite.StringColumn("content")
		AuthorColumn    = sqlite.StringColumn("author")
		CreatedColumn   = sqlite.TimestampColumn("created")
		UpdatedColumn   = sqlite.TimestampColumn("updated")
		allColumns      = sqlite.ColumnList{IDColumn, ArticleIDColumn, ContentColumn, AuthorColumn, CreatedColumn, UpdatedColumn}
		mutableColumns  = sqlite.ColumnList{ArticleIDColumn, ContentColumn, AuthorColumn, CreatedColumn, UpdatedColumn}
		defaultColumns  = sqlite.ColumnList{CreatedColumn, UpdatedColumn}
	)

	return commentsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		ArticleID: ArticleIDColumn,
		Content:   ContentColumn,
		Author:    AuthorColumn,
		Created:   CreatedColumn,
		Updated:   UpdatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
