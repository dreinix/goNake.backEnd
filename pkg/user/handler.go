package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
)

var (
	user User
)

func getAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT * FROM tbl_user`)
		if err != nil {
			render.JSON(w, r, "Something went wrong")
			return
		}
		if !rows.Next() {
			render.JSON(w, r, "There's not user on database!! add someone")
			return
		}
		var users []User
		for rows.Next() {
			var u User

			err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Password)
			if err != nil {
				json.NewEncoder(w).Encode("Something went wrong")
				log.Fatal(err)
			}
			users = append(users, u)
		}
		//log.Fatal(rows)
		render.JSON(w, r, users)
		//json.NewEncoder(w).Encode(users)
	}
}

func addUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		if _, err := db.Exec(`INSERT INTO tbl_user (full_name , usrn , pwd )
		VALUES ($1,$2,$3);`, user.Name, user.Username, user.Password); err != nil {
			log.Fatal(err)
		}

		msg := "You sucessfully added user " + (user.Name)
		render.JSON(w, r, msg)
	}
}
