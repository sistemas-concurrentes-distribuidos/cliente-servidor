package util

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//RespondWithError create JSON response with error message and status code
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	params := map[string]string{
		"message": msg,
		"code":    strconv.Itoa(code),
		"error":   http.StatusText(code),
	}
	RespondWithJSON(w, code, params)
}

//RespondWithJSON create JSON response with some structure given
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
