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

var Articles = newArticlesTable("", "articles", "")

type articlesTable struct {
	sqlite.Table

	// Columns
	ID      sqlite.ColumnInteger
	Title   sqlite.ColumnString
	Content sqlite.ColumnString
	Author  sqlite.ColumnString
	Created sqlite.ColumnTimestamp
	Updated sqlite.ColumnTimestamp

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
	DefaultColumns sqlite.ColumnList
}

type ArticlesTable struct {
	articlesTable

	EXCLUDED articlesTable
}

// AS creates new ArticlesTable with assigned alias
func (a ArticlesTable) AS(alias string) *ArticlesTable {
	return newArticlesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArticlesTable with assigned schema name
func (a ArticlesTable) FromSchema(schemaName string) *ArticlesTable {
	return newArticlesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArticlesTable with assigned table prefix
func (a ArticlesTable) WithPrefix(prefix string) *ArticlesTable {
	return newArticlesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArticlesTable with assigned table suffix
func (a ArticlesTable) WithSuffix(suffix string) *ArticlesTable {
	return newArticlesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArticlesTable(schemaName, tableName, alias string) *ArticlesTable {
	return &ArticlesTable{
		articlesTable: newArticlesTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newArticlesTableImpl("", "excluded", ""),
	}
}

func newArticlesTableImpl(schemaName, tableName, alias string) articlesTable {
	var (
		IDColumn       = sqlite.IntegerColumn("id")
		TitleColumn    = sqlite.StringColumn("title")
		ContentColumn  = sqlite.StringColumn("content")
		AuthorColumn   = sqlite.StringColumn("author")
		CreatedColumn  = sqlite.TimestampColumn("created")
		UpdatedColumn  = sqlite.TimestampColumn("updated")
		allColumns     = sqlite.ColumnList{IDColumn, TitleColumn, ContentColumn, AuthorColumn, CreatedColumn, UpdatedColumn}
		mutableColumns = sqlite.ColumnList{TitleColumn, ContentColumn, AuthorColumn, CreatedColumn, UpdatedColumn}
		defaultColumns = sqlite.ColumnList{CreatedColumn, UpdatedColumn}
	)

	return articlesTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:      IDColumn,
		Title:   TitleColumn,
		Content: ContentColumn,
		Author:  AuthorColumn,
		Created: CreatedColumn,
		Updated: UpdatedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
