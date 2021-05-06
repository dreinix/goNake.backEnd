package user

import (
	"database/sql"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/go-chi/chi/v5"
)

func UserRoutes(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", (getAllUsers(db)))
	router.Get("/getMe", auth.Authentication(getMe(db)))
	router.Get("/{id}", auth.Authentication(getUser(db)))
	router.Post("/", addUser(db))
	router.Post("/login", logIn(db))
	router.Post("/logout", auth.Authentication(LogOut(db)))
	return router
	//auth.Authenticate
}
