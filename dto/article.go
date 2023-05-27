package dto

import (
	"blog-api/model"
	"time"
)

type Article struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Path      string     `json:"path"`
	Markdown  string     `json:"markdown"`
	Tags      []string   `json:"tags"`
	Preview   string     `json:"preview"`
	CreatedAt time.Time  `json:"created_at"`
	Author    UserPublic `json:"author"`
}

func NewArticle(a model.Article, u UserPublic) Article {
	return Article{
		ID:        a.ID,
		Title:     a.Title,
		Path:      a.Path,
		Markdown:  a.Markdown,
		Tags:      a.Tags,
		Preview:   a.Preview,
		CreatedAt: a.CreatedAt,
		Author:    u,
	}
}

func NewArticles(a []model.Article, u []UserPublic) []Article {
	var articles []Article
	for i := range a {
		articles = append(articles, NewArticle(a[i], u[i]))
	}
	return articles
}
