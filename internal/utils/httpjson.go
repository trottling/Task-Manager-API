package utils

import (
	"encoding/json"
	"net/http"
)

// WriteJSON пишет ответ v в формате JSON с кодом статуса code.
func WriteJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
