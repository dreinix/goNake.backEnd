package Router

import (
	"net/http"

	"github.com/dreinix/gonake/pkg/score"
	"github.com/go-chi/chi/v5"
)

func StartServer() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})
	router.Mount("/api/scores", score.ScoreRoutes())
	http.ListenAndServe(":3001", router)
	return router
}
