package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.Url,
		UserID:    user.ID,
	}

	userFeed, err := config.db.CreateFeed(r.Context(), feed)

	response := databaseFeedToFeed(userFeed)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create Feed")
	}

	respondWithJson(w, http.StatusCreated, response)
}
