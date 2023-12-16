package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	type customErrors struct {
		Error string `json:"error"`
	}

	respBody := customErrors{
		Error: msg,
	}

	respondWithJSON(w, code, respBody)
}
