package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handleCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Id uuid.UUID `json:"feed_id"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	feed, err := config.db.GetFeedByID(r.Context(), params.Id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No Feed found")
		return
	}

	follows := database.CreateFeedFollowerParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follower, err := config.db.CreateFeedFollower(r.Context(), follows)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to follow Feed")
	}

	response := databaseFeedFollowerToFeedFollower(follower)

	respondWithJson(w, http.StatusCreated, response)
}
