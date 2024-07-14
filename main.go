package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/chaitu25/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	// Create a new instance of the server
	fmt.Println("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No PORT environment variable detected")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No DB URL environment variable detected")
	}

	con, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	queries := database.New(con)
	dbConfig := apiConfig{db: queries}

	go startScrapping(queries, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", dbConfig.handlerCreateUser)
	v1Router.Get("/users", dbConfig.middlewareAuth(dbConfig.handlerGetUserByApiKey))
	v1Router.Post("/feeds", dbConfig.middlewareAuth(dbConfig.handleCreateFeed))
	v1Router.Post("/feed_follow", dbConfig.middlewareAuth(dbConfig.handleCreateFeedFollows))
	v1Router.Get("/posts", dbConfig.middlewareAuth(dbConfig.handlerGetPostsByUser))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	fmt.Printf("Server started on port %s\n", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
