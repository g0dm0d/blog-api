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
	_, err := s.db.Exec("CALL create_article($1, $2, $3, $4, $5, $6)",
		opts.Title, opts.Path, opts.Markdown, pq.Array(opts.Tags), opts.Preview, opts.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionStore) GetArticle(opts store.GetArticleOpts) (store.Article, error) {
	var article store.Article

	res := s.db.QueryRow("SELECT * FROM get_article_by_path($1)", opts.Path)

	err := res.Scan(&article.ID, &article.Title, &article.Path, &article.Markdown, pq.Array(&article.Tags), &article.Preview, &article.Author_id, &article.Created_at)
	if err != nil {
		return store.Article{}, err
	}

	return article, nil
}

func (s *SessionStore) GetArticleForFeed(opts store.GetArticleFeed) ([]store.Article, error) {
	var articles []store.Article

	res, err := s.db.Query("SELECT * FROM get_article_feed($1)", opts.Page)
	if err != nil {
		return []store.Article{}, err
	}

	for res.Next() {
		var article store.Article
		err = res.Scan(&article.ID, &article.Title, &article.Path, &article.Markdown, pq.Array(&article.Tags), &article.Preview, &article.Author_id, &article.Created_at)
		if err != nil {
			return []store.Article{}, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (s *SessionStore) SearchArticle(opts store.SearchArticleOpts) ([]store.Article, error) {
	var articles []store.Article

	var authorIDParam interface{}
	if opts.AuthorID != "" {
		authorIDParam = opts.AuthorID
	} else {
		authorIDParam = nil
	}

	res, err := s.db.Query("SELECT * FROM search_article($1, $2, $3, $4)", opts.Page, pq.Array(opts.Tags), authorIDParam, opts.Text)
	if err != nil {
		return []store.Article{}, err
	}

	for res.Next() {
		var article store.Article
		err = res.Scan(&article.ID, &article.Title, &article.Path, &article.Markdown, pq.Array(&article.Tags), &article.Preview, &article.Author_id, &article.Created_at)
		if err != nil {
			return []store.Article{}, err
		}

		articles = append(articles, article)
	}

	return articles, nil
}
