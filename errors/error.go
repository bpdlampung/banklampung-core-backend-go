package errors

import "net/http"

type ErrorString struct {
	statusCode   int
	responseCode string
	message      string
}

func (e ErrorString) StatusCode() int {
	return e.statusCode
}

func (e ErrorString) ResponseCode() string {
	return e.responseCode
}

func (e ErrorString) Error() string {
	return e.message
}

func (e ErrorString) Message() string {
	return e.message
}

// Error
func Code(errorCode string) error {
	return &ErrorString{
		responseCode: errorCode,
	}
}

// Error
func CodeMsg(errorCode, msg string) error {
	return &ErrorString{
		responseCode: errorCode,
		message:      msg,
	}
}

// BadRequest will throw if the given request-body or params is not valid
func BadRequest(msg string) error {
	return &ErrorString{
		statusCode: http.StatusBadRequest,
		message:    msg,
	}
}

// NotFound will throw if the requested item is not exists
func NotFound(msg string) error {
	return &ErrorString{
		statusCode: http.StatusNotFound,
		message:    msg,
	}
}

// NotFound will throw if the requested item is not exists
func MethodNotAllowed(msg string) error {
	return &ErrorString{
		statusCode: http.StatusMethodNotAllowed,
		message:    msg,
	}
}

// Conflict will throw if the current action already exists
func Conflict(msg string) error {
	return &ErrorString{
		statusCode: http.StatusConflict,
		message:    msg,
	}
}

// InternalServerError will throw if any the Internal Server Error happen,
// Database, Third Party etc.
func InternalServerError(msg string) error {
	return &ErrorString{
		statusCode: http.StatusInternalServerError,
		message:    msg,
	}
}

func UnauthorizedError(msg string) error {
	return &ErrorString{
		statusCode: http.StatusUnauthorized,
		message:    msg,
	}
}

func ForbiddenError(msg string) error {
	return &ErrorString{
		statusCode: http.StatusForbidden,
		message:    msg,
	}
}
