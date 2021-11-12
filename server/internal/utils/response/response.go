package response

import (
	"encoding/json"
	"net/http"
)

type resp struct {
	Error string `json:"error" example:"message"`
}

func WithError(w http.ResponseWriter, err error) {
	msg := err.Error()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(StatusByMessage(msg))

	json.NewEncoder(w).Encode(Message(msg))
}

func WithoutError(w http.ResponseWriter, msg string, obj interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(StatusByMessage(msg))

	json.NewEncoder(w).Encode(obj)
}

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
