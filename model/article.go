package model

import (
	"blog-api/store"
	"time"
)

type Article struct {
	ID        int
	Title     string
	Path      string
	Markdown  string
	Tags      []string
	Preview   string
	AuthorID  int
	CreatedAt time.Time
}

func NewArticle(a store.Article) Article {
	return Article{
		ID:        a.ID,
		Title:     a.Title,
		Path:      a.Path,
		Markdown:  a.Markdown,
		Tags:      a.Tags,
		Preview:   a.Preview.String,
		AuthorID:  a.AuthorID,
		CreatedAt: a.CreatedAt,
	}
}

func NewArticles(a []store.Article) []Article {
	var articles []Article
	for i := range a {
		articles = append(articles, NewArticle(a[i]))
	}
	return articles
}
