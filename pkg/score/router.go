package score

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func ScoreRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAllScore(db))
	router.Post("/", addScore(db))
	return router
}
