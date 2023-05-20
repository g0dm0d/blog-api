package user

import (
	"blog-api/model"
	"blog-api/rest/req"
	"blog-api/store"
	"fmt"

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
		return fmt.Errorf("error parsing sign up request -> %w", err)
	}

	user, err := s.userStore.GetUser(store.GetUserOpts{
		Login: r.Login,
	})
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil {
		return err
	}
	accessToken, err := s.tokenManager.GenerateAccessToken(model.NewUser(user))
	refreshToken, err := s.tokenManager.GenerateRefreshToken()

	err = s.sessionStore.CreateSession(store.CreateSessionOpts{
		UserId:       user.Id,
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
