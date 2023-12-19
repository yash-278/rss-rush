package main

import (
	"encoding/json"
	"net/http"
	"rss-rush/internal/database"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	type response struct {
		Feed       `json:"feed"`
		FeedFollow `json:"feed_follow"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	userUuid, err := uuid.NewUUID()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create UUID")
		return
	}

	feedData := database.CreateFeedParams{
		ID:        userUuid,
		UserID:    user.ID,
		Name:      params.Name,
		Url:       params.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), feedData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	followUuid, err := uuid.NewUUID()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create UUID")
		return
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        followUuid,
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), feedFollowParams)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "cannot create a feed follow")
		return
	}

	respondWithJSON(w, 200, response{
		databaseFeedToFeed(feed),
		databaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (cfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, "No feeds found")
		return
	}
	feedResponse := []Feed{}

	for _, feed := range feeds {
		feedResponse = append(feedResponse, databaseFeedToFeed(feed))
	}

	respondWithJSON(w, 200, feedResponse)
}
