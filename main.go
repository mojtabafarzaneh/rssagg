package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mojtabafarzaneh/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal()
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("couldn't find the db url")
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("couldn't connect to the db")
	}

	queries := database.New(conn)
	apiCon := apiConfig{
		DB: queries,
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:" + portString,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handlerReadiness)
	v1router.Get("/err", HandleErrResponse)
	v1router.Post("/users", apiCon.handlerCreateUser)
	router.Mount("/v1", v1router)

	log.Printf("server is starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
