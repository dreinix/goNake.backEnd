package Router

import (
	"fmt"
	"net/http"

	"github.com/dreinix/gonake/pkg/database"
	"github.com/dreinix/gonake/pkg/score"
	"github.com/dreinix/gonake/pkg/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func StartServer() *chi.Mux {

	r, err := database.Conect()

	if err != nil {
		fmt.Println(err)
	}
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Gonected to goNake backend"))
	})
	router.Mount("/api/scores", score.ScoreRoutes(r))
	router.Mount("/api/users", user.UserRoutes(r))
	http.ListenAndServe("127.0.0.1:3001", router)
	return router
}
