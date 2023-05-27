package user

import (
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/tools/tokenmanager"
)

type CheckResponse struct {
	Status bool `json:"status"`
}

func (s *Service) Check(ctx *req.Ctx) error {
	user, err := s.tokenManager.ValidateJWTToken(ctx.BearerToken())
	if err != nil {
		return errs.ReturnError(ctx.Writer, errs.AccessTokenHasExpired)
	}

	if user == (&tokenmanager.Claims{}) {
		return ctx.JSON(CheckResponse{
			Status: false,
		})
	}

	return ctx.JSON(CheckResponse{
		Status: true,
	})
}
