package main

import (
	"encoding/json"
	"net/http"
	"rss-rush/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId string `json:"feed_id"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot decode parameters")
		return
	}

	newUuid, err := uuid.NewUUID()
	if err != nil || len(params.FeedId) != 36 {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create UUID")
		return
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        newUuid,
		UserID:    user.ID,
		FeedID:    uuid.MustParse(params.FeedId),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := cfg.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "cannot create a feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feed))
}
