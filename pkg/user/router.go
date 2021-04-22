package user

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func ScoreRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAllUsers(db))
	router.Post("/", addUser(db))

	router.Get("/{id}", getUser(db))
	return router
}
