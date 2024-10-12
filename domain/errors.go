package domain

import "net/http"

type APIError struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Status  int    `json:"status,omitempty"`
	Details any    `json:"details,omitempty"`
}

func NewAPIError(statusCode int, details any) *APIError {
	return &APIError{
		Type:    getTypeByStatusCode(statusCode),
		Title:   http.StatusText(statusCode),
		Status:  statusCode,
		Details: details,
	}
}

func getTypeByStatusCode(statusCode int) string {
	switch statusCode {
	case http.StatusBadRequest:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400"
	case http.StatusUnauthorized:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401"
	case http.StatusForbidden:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403"
	case http.StatusNotFound:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404"
	case http.StatusConflict:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409"
	case http.StatusInternalServerError:
		return "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500"
	}
	return ""
}
