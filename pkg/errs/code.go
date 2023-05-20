package errs

import "net/http"

type Code struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
}

var (
	AccessTokenHasExpired       = Code{Code: 1001, Description: "Invalid credentials, try again."}
	AccessTokenInvalidFormat    = Code{Code: 1002, Description: "Access token has invalid format."}
	AccessTokenInvalidSignature = Code{Code: 1003, Description: "Access token has invalid signature."}
)

func HTTPStatusCode(c uint16) int {
	switch c {
	case AccessTokenHasExpired.Code:
		return http.StatusUnauthorized
	case AccessTokenInvalidFormat.Code:
		return http.StatusUnauthorized
	case AccessTokenInvalidSignature.Code:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
