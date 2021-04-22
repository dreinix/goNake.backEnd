package score

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
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
		json.NewDecoder(r.Body).Decode(&score)
		msg := "You got " + strconv.Itoa(score.Value)
		json.NewEncoder(w).Encode(msg)
		/*
			stmt, err := db.Prepare(`INSERT INTO tbl_score (Value, User)
				VALUES ('?', '?', 'Product Manager', NOW(), NOW());`)
			r, err := stmt.Exec(score.Value, score.User.ID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(r.RowsAffected())*/
	}
}

//func(s *Storage)
