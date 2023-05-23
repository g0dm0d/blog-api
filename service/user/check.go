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
		errs.ReturnError(ctx.Writer, errs.AccessTokenHasExpired)
		return err
	}

	if user == (&tokenmanager.Claims{}) {
		ctx.JSON(CheckResponse{
			Status: false,
		})
	}

	ctx.JSON(CheckResponse{
		Status: true,
	})

	return nil
}
