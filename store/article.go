package store

import (
	"database/sql"
	"time"
)

type Article struct {
	ID        int
	Title     string
	Path      string
	Markdown  string
	Tags      []string
	Preview   sql.NullString
	AuthorID  int
	CreatedAt time.Time
}

type CreateArticleOpts struct {
	Title    string
	Path     string
	Markdown string
	Tags     []string
	Preview  string
	AuthorID int
}

type GetArticleOpts struct {
	Path string
}

type GetArticleFeed struct {
	Page int
}

type SearchArticleOpts struct {
	Page     int
	Tags     interface{}
	Text     interface{}
	AuthorID interface{}
}

type ArticleStore interface {
	CreateArticle(opts CreateArticleOpts) error
	GetArticle(opts GetArticleOpts) (Article, error)
	SearchArticle(opts SearchArticleOpts) ([]Article, error)
	GetArticleForFeed(opts GetArticleFeed) ([]Article, error)
}
