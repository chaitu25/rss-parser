package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (config *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
	}

	dbUser := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := config.db.CreateUser(r.Context(), dbUser)

	response := databaseUserToUser(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user")
	}

	respondWithJson(w, http.StatusCreated, response)
}

func (config *apiConfig) handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, http.StatusOK, user)
}

func (config *apiConfig) handlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := config.db.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  5,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to fetch latest posts")
	}
	respondWithJson(w, http.StatusOK, databasePostsToPosts(posts))
}
