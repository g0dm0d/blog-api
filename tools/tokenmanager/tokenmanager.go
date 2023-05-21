package tokenmanager

import (
	"blog-api/model"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager interface {
	GenerateAccessToken(user model.User) (string, error)
	GenerateRefreshToken() (string, error)
	ValidateJWTToken(token string) (*Claims, error)
}

type Tool struct {
	secret string
}

func New(secret string) *Tool {
	return &Tool{
		secret: secret,
	}
}

const (
	AccessLiveTime     = 24 * time.Hour
	RefreshTokenLength = 16
)

type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	Role     int16  `json:"role"`
	jwt.RegisteredClaims
}

func (t *Tool) GenerateAccessToken(user model.User) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessLiveTime)),
		},
		Username: user.Username,
		UserID:   user.Id,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (t *Tool) GenerateRefreshToken() (string, error) {
	refreshToken, err := GenerateRandomSalt(RefreshTokenLength)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (t *Tool) ValidateJWTToken(token string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		return &Claims{}, err
	}

	if !tkn.Valid {
		return &Claims{}, nil
	}

	return claims, nil
}

func GenerateRandomSalt(length int) (string, error) {
	salt := make([]byte, length)

	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	saltHex := hex.EncodeToString(salt)

	return saltHex, nil
}
