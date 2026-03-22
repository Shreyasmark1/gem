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

func InternalServerError(msg string) *ApiException {
	return InternalErrorServerError(msg, nil)
}

func InternalErrorServerError(msg string, err error) *ApiException {
	return &ApiException{
		Message:    msg,
		StatusCode: http.StatusInternalServerError,
		Err:        err,
		StackTrace: string(debug.Stack()),
	}
}

func NotFound(msg string) *ApiException {
	return &ApiException{Message: msg, StatusCode: http.StatusNotFound}
}
