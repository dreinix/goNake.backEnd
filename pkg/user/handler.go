package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var (
	user User
)

func getAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT usr_id,full_name,usrn,stat FROM tbl_user`)
		if err != nil {
			w.WriteHeader(500)
			render.JSON(w, r, "Something went wrong")
			return
		}
		if !rows.Next() {
			render.JSON(w, r, "There's not user on database!! add someone")
			return
		}
		//where stat = $1 "actv"
		if err := db.QueryRow(`SELECT usr_id,full_name,usrn,stat FROM tbl_user where stat = $1 `, "actv").
			Scan(&user.ID, &user.Name, &user.Username, &user.Status); err != nil {
			render.JSON(w, r, err)
			return
		}
		var users []User
		//The first value is ignore because "next"
		users = append(users, user)
		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Status)
			if err != nil {
				w.WriteHeader(500)
				fmt.Println(err.Error())
				render.JSON(w, r, "Something went wrong, please try again later")
			}
			users = append(users, u)
		}

		// if there's only one result there's no next
		render.JSON(w, r, users)
	}
}

func getUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		var user User
		if err := db.QueryRow(`SELECT usr_id,full_name,usrn FROM tbl_user where usr_id = $1 and stat = $2`, id, "actv").Scan(&user.ID, &user.Name, &user.Username); err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, "This user does not exist.")
			return
		}
		render.JSON(w, r, user)
	}
}

func addUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		if _, err := db.Exec(`INSERT INTO tbl_user (full_name , usrn , pwd, stat )
		VALUES ($1,$2,$3,$4);`, user.Name, user.Username, user.Password, "actv"); err != nil {

			if strings.Contains(err.Error(), "unique") {
				msg := "Username already exist, we couldn't create your account."
				w.WriteHeader(http.StatusBadRequest)
				render.JSON(w, r, msg)
				return
			}
			msg := "Something went wrong. Please, try again later."
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, msg)
			return
		}
		user.Password = ""
		msg := "You sucessfully added user " + (user.Name)
		render.JSON(w, r, msg)
	}
}

func logIn(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&user)
		if err := db.QueryRow(`SELECT usr_id,full_name,usrn FROM tbl_user where usrn = $1 and pwd = $2 and stat = $3`, user.Username, user.Password, "actv").
			Scan(&user.ID, &user.Name, &user.Username); err != nil {
			render.JSON(w, r, "username or password incorrect")
			//render.JSON(w, r, "This user does not exist.")
			return
		}
		token, err := auth.CreateToken(user.Username)
		if err != nil {
			render.JSON(w, r, "We couldn't log you in, please try again")
		}
		expirationTime := time.Now().Add(5 * 24 * time.Hour)
		//fmt.Printf("you saved jwt  \n" + token)
		http.SetCookie(w, &http.Cookie{
			Name:    "jwt",
			Value:   token,
			Expires: expirationTime,
			Secure:  false,
			Path:    "/",
		})
		type Message struct {
			Msg   string
			Token string
		}
		msg := Message{Msg: "login successfully", Token: token}
		user.Password = ""
		render.JSON(w, r, msg)
	}
}

func LogOut(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expirationTime := time.Now().Add(2 * time.Second)
		http.SetCookie(w, &http.Cookie{
			Name:    "jwt",
			Value:   "loggedout",
			Expires: expirationTime,
			Path:    "/",
		})
		msg := "logout successfully, come back soon!"

		render.JSON(w, r, msg)
	}
}
