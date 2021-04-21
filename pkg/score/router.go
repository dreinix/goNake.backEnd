package score

import (
	"github.com/go-chi/chi/v5"
)

func ScoreRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAllScore())
	router.Post("/", addScore())
	return router
}
