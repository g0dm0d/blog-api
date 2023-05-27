package article

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *Service) GetArticle(ctx *req.Ctx) error {
	path := chi.URLParam(ctx.Request, "path")

	article, err := s.articleStore.GetArticle(store.GetArticleOpts{
		Path: path,
	})
	if err != nil {
		return err
	}

	user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
	if err != nil {
		return err
	}

	return ctx.JSON(dto.NewArticle(model.NewArticle(article), dto.NewUserPublic(model.NewUser(user))))
}

func (s *Service) GetArticleForFeed(ctx *req.Ctx) error {
	pageStr := ctx.Request.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	articles, err := s.articleStore.GetArticleForFeed(store.GetArticleFeed{Page: page})
	if err != nil {
		return err
	}

	var users []store.User
	for _, article := range articles {
		user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	return ctx.JSON(dto.NewArticles(model.NewArticles(articles), dto.NewUsersPublic(model.NewUsers(users))))
}

func (s *Service) SearchArticle(ctx *req.Ctx) error {
	tagsStr := ctx.Request.URL.Query().Get("tags[]")
	text := ctx.Request.URL.Query().Get("text")
	author := ctx.Request.URL.Query().Get("author")
	pageStr := ctx.Request.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}

	var tags interface{}
	if toInterface(tagsStr) == nil {
		tags = nil
	} else {
		tags = strings.Split(tagsStr, ",")
	}

	articles, err := s.articleStore.SearchArticle(store.SearchArticleOpts{
		Page:     page,
		Tags:     tags,
		Text:     toInterface(text),
		AuthorID: toInterface(author),
	})
	if err != nil {
		return err
	}

	var users []store.User
	for _, article := range articles {
		user, err := s.userStore.GetUserByID(store.GetUserOpts{ID: article.Author_id})
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	return ctx.JSON(dto.NewArticles(model.NewArticles(articles), dto.NewUsersPublic(model.NewUsers(users))))
}

func toInterface(param interface{}) interface{} {
	if param == "" {
		return nil
	}
	return param
}
