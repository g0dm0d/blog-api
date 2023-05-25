package user

import (
	"blog-api/dto"
	"blog-api/model"
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"

	"github.com/go-chi/chi/v5"
)

func (s *Service) GetByUsername(ctx *req.Ctx) error {
	username := chi.URLParam(ctx.Request, "username")

	user, err := s.userStore.GetUserByUsername(store.GetUserOpts{Username: username})

	if err != nil {
		errs.ReturnError(ctx.Writer, errs.UserNotFound)
		return err
	}

	return ctx.JSON(dto.NewUserPublic(model.NewUser(user)))
}
