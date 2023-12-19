package main

import (
	"net/http"
	"rss-rush/internal/database"
)

func (cfg *apiConfig) handleFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := cfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "cannot find a feed follow for user")
		return
	}

	feedResponse := []FeedFollow{}

	for _, feedFollow := range feedFollows {
		feedResponse = append(feedResponse, databaseFeedFollowToFeedFollow(feedFollow))
	}
	respondWithJSON(w, http.StatusOK, feedResponse)
}
