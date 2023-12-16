package main

import (
	"net/http"
)

func (cfg *apiConfig) readinessSuccessHandler(w http.ResponseWriter, r *http.Request) {
	type status struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, status{
		Status: "ok",
	})
}

func (cfg *apiConfig) readinessErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
