package wrapper

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, code int, result *Result, message string) {
	if code < 200 || code >= 300 {
		result.Err = ErrInternalServer
		ResponseError(w, result)
		return
	}
	responseData := httpResponse{
		Success: true,
		Data:    result.Data,
		Message: message,
		Code:    code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseData)
	return
}

// ResponseError will respond http error with json serializer
func ResponseError(w http.ResponseWriter, result *Result) {
	code := http.StatusBadRequest
	responseData := httpResponse{
		Success: false,
		Code:    code,
		Message: result.Err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(responseData)
	return
}
