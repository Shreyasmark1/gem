package response

import (
	"net/http"
	"runtime/debug"
)

type ApiException struct {
	Message    string
	StatusCode int
	Err        error
	StackTrace string
}

func (e *ApiException) Error() string {
	return e.Message
}

func BadRequest(msg string) *ApiException {
	return &ApiException{Message: msg, StatusCode: http.StatusBadRequest}
}

func Unauthorized(msg string) *ApiException {
	return &ApiException{Message: msg, StatusCode: http.StatusUnauthorized}
}

func InternalServerError(msg string, err ...error) *ApiException {
	var e error
	if len(err) > 0 {
		e = err[0]
	}
	return &ApiException{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
		Err:        e,
		StackTrace: string(debug.Stack()),
	}
}

func NotFound(msg string) *ApiException {
	return &ApiException{Message: msg, StatusCode: http.StatusNotFound}
}
