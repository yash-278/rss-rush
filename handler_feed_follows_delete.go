package main

import (
	"net/http"
	"rss-rush/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	type status struct {
		Status string `json:"status"`
	}

	idParam := chi.URLParam(r, "id")

	if idParam == "" || len(idParam) != 36 {
		respondWithError(w, http.StatusInternalServerError, "invalid feed follow id")
		return
	}

	err := cfg.DB.DeleteFeedFollow(r.Context(), uuid.MustParse(idParam))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "invalid feed follow id")
		return
	}

	respondWithJSON(w, http.StatusOK, status{
		Status: "ok",
	})
}
