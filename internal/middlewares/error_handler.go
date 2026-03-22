package middleware

import (
	"errors"
	"gem/pkg/response"
	"log/slog"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorHandler(h AppHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {

			var apiErr *response.ApiException

			if errors.As(err, &apiErr) {

				// if apiErr.StatusCode >= 500 {
				// 	logFailure(apiErr.Message, apiErr.Err, apiErr.StackTrace)
				// }

				logFailure(apiErr.Message, apiErr.Err, apiErr.StackTrace)

				response.Failed(w, apiErr.StatusCode, apiErr.Message)
				return
			}

			// Unknown Error (the "catch-all")
			logFailure(apiErr.Error(), err, "none")
			response.Failed(w, http.StatusInternalServerError, "An unexpected error occurred")
		}
	}
}

func logFailure(msg string, err error, stack string) {
	slog.Error(msg,
		"details", err,
		"stack", stack,
	)
}
