package user

import (
	"blog-api/rest/req"
	"blog-api/store"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,max=100,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func (s *Service) Signup(ctx *req.Ctx) error {
	var r SignUpRequest

	err := ctx.ParseJSON(&r)
	if err != nil {
		return fmt.Errorf("error parsing sign up request -> %w", err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	_, err = s.userStore.CreateUser(store.CreateUserOpts{
		Username: r.Username,
		Password: string(passwordHash),
		Email:    r.Email,
	})
	return err
}
