package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dreinix/gonake/pkg/database"
	"github.com/go-chi/render"
)

func CreateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_TOKEN")))

}

func tokenValid(r *http.Request) error {
	tokenString := extractToken(r)
	if tokenString == "lo" {
		return fmt.Errorf("no active session")
	}
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func extractToken(r *http.Request) string {

	if r.Header["Authorization"] != nil {
		bearerToken := r.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			return strings.Split(bearerToken, " ")[1]
		}
		return ""
	} else {
		c, err := r.Cookie("jwt")
		if err != nil {
			return "lo"
		}
		token := c.Value
		if token != "" {
			return token
		}
	}
	return ""

}

func ExtractTokenID(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		usrn := fmt.Sprintf("%s", claims["user_id"])
		if usrn != "" {
			return usrn, nil
		}
		return "", nil
	}
	return "", nil
}

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
type Message struct {
	Msg string
}

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/* w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*") */
		db, _ := database.Conect()
		err := tokenValid(r)

		if err != nil {
			msg := Message{err.Error()}
			w.WriteHeader(400)
			render.JSON(w, r, msg)
			return
		}

		id, _ := ExtractTokenID(r)
		var user User
		if err := db.QueryRow(`SELECT usr_id,full_name,usrn FROM tbl_user where usrn = $1 and stat=$2`, id, "actv").Scan(&user.ID, &user.Name, &user.Username); err != nil {
			w.WriteHeader(200)
			msg := Message{"This user does not exist anymore"}
			render.JSON(w, r, msg)
			return
		}

		ctx := context.WithValue(r.Context(), "usr", user)
		next(w, r.WithContext(ctx))
	}
}

/*package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

var TokenAuth *jwtauth.JWTAuth

func GoDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func CreateToken(user_id string) (string, error) {
	token := GoDotEnvVariable(os.Getenv("JWT_TOKEN"))
	TokenAuth = jwtauth.New("HS256", []byte(token), nil)
	_, tokenString, _ := TokenAuth.Encode(map[string]interface{}{"user_id": user_id})
	return tokenString, nil
}

func Authenticate(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(401)
				render.JSON(w, r, "UNAUTHORIZED")
				return
			}
			w.WriteHeader(401)
			render.JSON(w, r, "invalid token")
			return
		}
		us_id, err := TokenAuth.Decode(c.Value)
		if err != nil {
			w.WriteHeader(401)
			render.JSON(w, r, us_id)
			render.JSON(w, r, "your token has expire or the user changed the information")
			return
		}
		endpoint(w, r)
	})
}
*/
func Verify() {

}

/*
func Authenticate(endpoint http.HandlerFunc) http.HandlerFunc {

	return nil
}*/
