package score

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/go-chi/render"
)

var (
	score Score
)

func getAllScore(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT * FROM tbl_score`)
		if err != nil {
			w.WriteHeader(500)
			render.JSON(w, r, "Something went wrong")
			return
		}
		if !rows.Next() {
			render.JSON(w, r, "no score saved")
			return
		}
		var id int64
		if err := db.QueryRow(`SELECT * from tbl_score`).
			Scan(&score.ID, &score.Value, &id, &score.Date); err != nil {
			render.JSON(w, r, err)
			return
		}
		var user auth.User
		if err := db.QueryRow(`SELECT usr_id,full_name,usrn FROM tbl_user where usr_id = $1 and stat = $2`, id, "actv").Scan(&user.ID, &user.Name, &user.Username); err != nil {
			fmt.Println("The user associated with " + strconv.Itoa(score.ID) + " score does not exist anymore.")
		}
		var scores []Score
		//The first value is ignore because "next"
		score.User = user
		scores = append(scores, score)
		for rows.Next() {
			var s Score
			err := rows.Scan(&s.ID, &s.Value, &id, &s.Date)
			if err != nil {
				w.WriteHeader(500)
				fmt.Println(err.Error())
				render.JSON(w, r, "Something went wrong, please try again later")
				return
			}
			if err := db.QueryRow(`SELECT usr_id,full_name,usrn FROM tbl_user where usr_id = $1 and stat = $2`, id, "actv").Scan(&user.ID, &user.Name, &user.Username); err != nil {
				fmt.Println("The user associated with " + strconv.Itoa(score.ID) + " score does not exist anymore.")
			}
			s.User = user
			scores = append(scores, s)
		}
		// if there's only one result there's no next
		render.JSON(w, r, scores)
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
