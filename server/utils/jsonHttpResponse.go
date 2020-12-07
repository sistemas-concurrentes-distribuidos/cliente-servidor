package util

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	params := map[string]string{
		"message": msg,
		"code":    strconv.Itoa(code),
		"error":   http.StatusText(code),
	}
	RespondWithJson(w, code, params)
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
