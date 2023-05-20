package middleware

import (
	"blog-api/pkg/errs"
	"net/http"
	"strings"
)

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")
		authorizationSplitted := strings.Split(authorization, " ")

		if len(authorizationSplitted) != 2 {
			errs.ReturnError(w, errs.AccessTokenInvalidFormat)
			return
		}

		if authorizationSplitted[0] != "Bearer" {
			errs.ReturnError(w, errs.AccessTokenInvalidFormat)
			return
		}

		_, err := m.tokenManager.ValidateJWTToken(authorizationSplitted[1])
		if err != nil {
			errs.ReturnError(w, errs.AccessTokenInvalidSignature)
			return
		}

		next.ServeHTTP(w, r)
	})
}
