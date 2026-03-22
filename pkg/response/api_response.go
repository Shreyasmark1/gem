package response

import (
	"encoding/json"
	"net/http"
)

type ResponseType struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Send(w http.ResponseWriter, data any) error {
	res := ResponseType{
		Status:  true,
		Message: "Request completed successfully",
		Data:    data,
	}
	return sendJSON(w, http.StatusOK, res)
}

func Successful(w http.ResponseWriter) error {
	res := ResponseType{
		Status:  true,
		Message: "Request completed successfully",
		Data:    nil,
	}

	return sendJSON(w, http.StatusOK, res)
}

func Failed(w http.ResponseWriter, statusCode int, message string) error {
	if message == "" {
		message = "failed to process your request"
	}
	res := ResponseType{
		Status:  false,
		Message: message,
		Data:    nil,
	}
	return sendJSON(w, statusCode, res)
}

func sendJSON(w http.ResponseWriter, statusCode int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(payload)
}
