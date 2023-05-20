package user

import (
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Service) Refresh(ctx *req.Ctx) error {
	var r RefreshRequest

	ctx.ParseJSON(&r)

	refreshToken, err := s.tokenManager.GenerateRefreshToken()
	if err != nil {
		return err
	}

	userID, err := s.sessionStore.UpdateSession(store.UpdateSessionOpts{
		NewToken: refreshToken,
		OldToken: r.RefreshToken,
	})
	if err != nil {
		return err
	}

	user, err := s.userStore.GetUserByID(userID)
	if err != nil {
		return err
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(model.NewUser(user))
	if err != nil {
		return err
	}

	return ctx.JSON(RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
