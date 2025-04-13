package main

import (
	"fmt"
	"time"
)

type Article struct {
	ID      int
	Title   string
	Content string
	Author  string
	Created time.Time
	Updated time.Time
}

func (a *Article) String() string {
	return fmt.Sprintf("Article(%d)", a.ID)
}
