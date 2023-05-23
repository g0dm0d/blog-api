package errs

import "net/http"

type Code struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
}

var (
	// System
	InvalidJSON = Code{Code: 100, Description: "Invalid JSON body"}

	// Auth codes
	AccessTokenHasExpired       = Code{Code: 1001, Description: "Invalid credentials, try again."}
	AccessTokenInvalidFormat    = Code{Code: 1002, Description: "Access token has invalid format."}
	AccessTokenInvalidSignature = Code{Code: 1003, Description: "Access token has invalid signature."}
	IncorrectLoginOrPassword    = Code{Code: 1004, Description: "Invalid login credentials"}

	// User codes
	HasActiveRoom     = Code{Code: 1500, Description: "User has an active room, you can't join another one."}
	UserNotFound      = Code{Code: 1501, Description: "User not found, make sure user id is correct."}
	UserAlreadyExists = Code{Code: 1502, Description: "User already exists."}

	// Article codes
)

func HTTPStatusCode(c uint16) int {
	switch c {
	case AccessTokenHasExpired.Code:
		return http.StatusUnauthorized
	case AccessTokenInvalidFormat.Code:
		return http.StatusUnauthorized
	case AccessTokenInvalidSignature.Code:
		return http.StatusUnauthorized
	case IncorrectLoginOrPassword.Code:
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}
