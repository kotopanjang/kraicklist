package wrapper

import (
	"encoding/json"
	"net/http"
)

// ResponseSuccess return http success response
func ResponseSuccess(w http.ResponseWriter, result *Result, message string) {
	responseData := httpResponse{
		Success: true,
		Data:    result.Data,
		Message: message,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	return
}

// ResponseError return http error response
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
