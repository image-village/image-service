package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ErrorContent - Shape of error message
type ErrorContent struct {
	Message string `json:"message"`
}

// Errs - Slice of error structs
type Errs []ErrorContent

// JSON - Shape response data
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR shape error response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	var errors Errs = make(Errs, 0)
	error := ErrorContent{
		Message: err.Error(),
	}
	errors = append(errors, error)

	if err != nil {
		JSON(w, statusCode, struct {
			Errors Errs `json:"errors"`
		}{
			Errors: errors,
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}