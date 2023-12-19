package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"rss-rush/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to DB")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	r := chi.NewRouter()
	apiRouter := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)

	r.Mount("/v1", apiRouter)

	apiRouter.Get("/readiness", apiCfg.readinessSuccessHandler)
	apiRouter.Get("/err", apiCfg.readinessErrHandler)

	apiRouter.Post("/users", apiCfg.handleCreateUser)
	apiRouter.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUserByApiKey))

	apiRouter.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeeds))
	apiRouter.Get("/feeds", apiCfg.handleGetFeeds)

	apiRouter.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowsGet))
	apiRouter.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handleFeedFollowsCreate))
	apiRouter.Delete("/feed_follows/{id}", apiCfg.middlewareAuth(apiCfg.handleFeedFollowDelete))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
