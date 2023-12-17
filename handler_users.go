package main

import (
	"encoding/json"
	"net/http"
	"rss-rush/internal/database"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create UUID")
		return
	}

	userData := database.CreateUserParams{
		ID:        uuid,
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	user, err := cfg.DB.CreateUser(r.Context(), userData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (cfg *apiConfig) handleGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Authorization")

	key := strings.Split(apiKey, " ")
	if len(key) < 2 {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key")
		return
	}

	user, err := cfg.DB.GetUserByApiKey(r.Context(), key[1])
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid API Key")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
