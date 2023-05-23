package store

import "time"

type Article struct {
	ID         int
	Name       string
	Markdown   string
	Tags       []string
	Preview    string
	Author_id  int
	Created_at time.Time
}

type CreateArticleOpts struct {
	Name     string
	Markdown string
	Tags     []string
	Preview  string
	AuthorID int
}

type GetArticleOpts struct {
	ID int
}

type ArticleStore interface {
	CreateArticle(opts CreateArticleOpts) error
	GetArticle(opts GetArticleOpts) (Article, error)
}
