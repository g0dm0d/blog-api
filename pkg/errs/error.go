package errs

import (
	"encoding/json"
	"net/http"
)

func ReturnError(w http.ResponseWriter, code Code) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(HTTPStatusCode(code.Code))

	json.NewEncoder(w).Encode(code)
}
