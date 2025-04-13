package main

import (
	"fmt"
	"time"
)

type Article struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Comments []Comment `json:"comments"`
}

func (a *Article) String() string {
	return fmt.Sprintf("Article{Title: %s, Author: %s, Comments: %v}", a.Title, a.Author, a.Comments)
}

type Comment struct {
	ID      int       `json:"id"`
	Content string    `json:"Content"`
	Author  string    `json:"author"`
	Article int       `json:"article"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment{Article: %d, Author: %s}", c.Article, c.Author)
}
