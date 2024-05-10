package main

import (
	"net/http"

	"github.com/chaitu25/rss-aggregator/internal/auth"
	"github.com/chaitu25/rss-aggregator/internal/database"
)

type authorizedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (config *apiConfig) middlewareAuth(next authorizedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Not Authorized to read user")
			return
		}

		databaseUser, err := config.db.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, http.StatusBadRequest, "No User found")
			return
		}

		next(w, r, databaseUser)
	}
}
