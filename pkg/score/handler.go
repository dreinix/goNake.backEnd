package score

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/go-chi/render"
)

var (
	score Score
)

func getAllScore(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Inix")
	}
}

func addScore(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("usr").(auth.User)
		json.NewDecoder(r.Body).Decode(&score)
		if _, err := db.Exec(`INSERT INTO tbl_score (score , usr ,scored_on )
			VALUES ($1,$2,$3);`, score.Value, user.ID, time.Now()); err != nil {
			msg := "Something went wrong. Please, try again later."
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, msg)
			return
		}
		msg := "scored saved"
		render.JSON(w, r, msg)
	}
}

//func(s *Storage)
