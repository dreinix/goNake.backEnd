package Router

import (
	"fmt"
	"net/http"

	"github.com/dreinix/gonake/pkg/database"
	"github.com/dreinix/gonake/pkg/score"
	"github.com/dreinix/gonake/pkg/user"
	"github.com/go-chi/chi/v5"
)

func StartServer() *chi.Mux {

	r, err := database.Conect()

	if err != nil {
		fmt.Println(err)
	}
	router := chi.NewRouter()

	router.Mount("/api/scores", score.ScoreRoutes(r))
	router.Mount("/api/users", user.UserRoutes(r))
	http.ListenAndServe("127.0.0.1:3001", router)
	return router
}
