package httperror

import "fmt"

type HttpError struct {
	Message    string
	StatusCode int
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

func New(message string, status int) *HttpError {
	return &HttpError{Message: message, StatusCode: status}
}

// Helpers
func BadRequest(message string) *HttpError {
	return New(message, 400)
}

func Unauthorized(message string) *HttpError {
	return New(message, 401)
}

func Forbidden(message string) *HttpError {
	return New(message, 403)
}

func NotFound(message string) *HttpError {
	return New(message, 404)
}
