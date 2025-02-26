package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseData adalah struktur untuk response data
type ResponseData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PaginationData adalah struktur untuk data pagination
type PaginationData struct {
	Total       int64       `json:"total"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	LastPage    int         `json:"last_page"`
	Data        interface{} `json:"data"`
}

// ErrorResponse adalah struktur untuk response error
type ErrorData struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
}

// SuccessResponse mengirim response sukses
func SuccessResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ResponseData{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse mengirim response error
func ErrorResponse(w http.ResponseWriter, code int, message string, errors interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorData{
		Status:  false,
		Message: message,
		Errors:  errors,
	})
}

// ValidationErrorResponse mengirim response error validasi
func ValidationErrorResponse(w http.ResponseWriter, message string, errors interface{}) {
	ErrorResponse(w, 422, message, errors)
}

// PaginationResponse mengirim response dengan pagination
func PaginationResponse(w http.ResponseWriter, message string, total int64, perPage, page int, data interface{}) {
	lastPage := int(total) / perPage
	if int(total)%perPage != 0 {
		lastPage++
	}

	paginationData := PaginationData{
		Total:       total,
		PerPage:     perPage,
		CurrentPage: page,
		LastPage:    lastPage,
		Data:        data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseData{
		Status:  true,
		Message: message,
		Data:    paginationData,
	})
}

// NotFoundResponse mengirim response 404
func NotFoundResponse(w http.ResponseWriter, message string) {
	ErrorResponse(w, http.StatusNotFound, message, nil)
}

// ServerErrorResponse mengirim response 500
func ServerErrorResponse(w http.ResponseWriter, err error) {
	ErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", map[string]string{
		"error": err.Error(),
	})
}

// UnauthorizedResponse mengirim response 401
func UnauthorizedResponse(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
}

// ForbiddenResponse mengirim response 403
func ForbiddenResponse(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusForbidden, "Forbidden", nil)
}
