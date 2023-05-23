package user

import (
	"blog-api/model"
	"blog-api/pkg/errs"
	"blog-api/rest/req"
	"blog-api/store"

	"golang.org/x/crypto/bcrypt"
)

type SignInRequest struct {
	Login    string `json:"login" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

type SignInResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Service) Signin(ctx *req.Ctx) error {
	var r SignInRequest

	err := ctx.ParseJSON(&r)
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.InvalidJSON)
		return nil
	}

	user, err := s.userStore.GetUserByLogin(store.GetUserOpts{
		Login: r.Login,
	})
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.IncorrectLoginOrPassword)
		return nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		errs.ReturnError(ctx.Writer, errs.IncorrectLoginOrPassword)
		return nil
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(model.NewUser(user))
	refreshToken, err := s.tokenManager.GenerateRefreshToken()

	err = s.sessionStore.CreateSession(store.CreateSessionOpts{
		UserId:       user.ID,
		RefreshToken: refreshToken,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
