package score

import (
	"database/sql"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/go-chi/chi/v5"
)

func ScoreRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAllScore(db))
	router.Get("/top", getTop(db))
	router.Post("/", auth.Authentication(addScore(db)))
	return router
}
