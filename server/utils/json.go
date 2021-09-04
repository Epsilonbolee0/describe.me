package utils

import (
	"encoding/json"
	"net/http"
)

func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
