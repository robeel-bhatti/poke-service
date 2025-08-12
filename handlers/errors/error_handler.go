// Package errors provides custom app-wide error handling
package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"poke-ai-service/util/constants"
	"time"
)

// vars to be used in app-wide error messages and can be used to unwrap errors
var (
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("forbidden")

	ErrMap = map[error]int{
		ErrNotFound:            http.StatusNotFound,
		ErrInternalServerError: http.StatusInternalServerError,
		ErrBadRequest:          http.StatusBadRequest,
		ErrUnauthorized:        http.StatusUnauthorized,
		ErrForbidden:           http.StatusForbidden,
	}
)

// AppError struct defines API error response body fields
type AppError struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Path      string    `json:"path"`
	Status    int       `json:"status"`
	Reason    string    `json:"reason"`
	RequestId string    `json:"request_id"`
}

func NewAppError(message string, status int, path string, reqId string) *AppError {
	return &AppError{
		Message:   message,
		Timestamp: time.Now(),
		Status:    status,
		Reason:    http.StatusText(status),
		Path:      path,
		RequestId: reqId,
	}
}

func BuildAppError(err error, path string, reqId string) *AppError {
	var httpCode int
	for e, code := range ErrMap {
		if errors.Is(err, e) {
			httpCode = code
		}
	}
	return NewAppError(err.Error(), httpCode, path, reqId)
}

// CreateErrorResponse creates and returns custom error HTTP response
func CreateErrorResponse(err error, w http.ResponseWriter, r *http.Request) {
	appErr := BuildAppError(err, r.URL.Path, r.Header.Get(constants.RequestIdKey))
	w.WriteHeader(appErr.Status)
	w.Header().Set(constants.ContentTypeKey, constants.ContentTypeValue)
	w.Header().Set(constants.RequestIdKey, appErr.RequestId)
	json.NewEncoder(w).Encode(appErr)
}
