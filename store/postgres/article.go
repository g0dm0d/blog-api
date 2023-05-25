package postgres

import (
	"blog-api/store"
	"database/sql"

	"github.com/lib/pq"
)

type ArticleStore struct {
	db *sql.DB
}

func NewArticleStore(db *sql.DB) store.ArticleStore {
	return &SessionStore{
		db: db,
	}
}

func (s *SessionStore) CreateArticle(opts store.CreateArticleOpts) error {
	_, err := s.db.Exec("CALL create_article($1, $2, $3, $4, $5)",
		opts.Title, opts.Markdown, pq.Array(opts.Tags), opts.Preview, opts.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionStore) GetArticle(opts store.GetArticleOpts) (store.Article, error) {
	var article store.Article

	res := s.db.QueryRow("SELECT * FROM get_article_by_id($1)", opts.ID)

	err := res.Scan(&article.Title, &article.Markdown, pq.Array(&article.Tags), &article.Preview, &article.Author_id, &article.Created_at)
	if err != nil {
		return store.Article{}, err
	}

	return article, nil
}
