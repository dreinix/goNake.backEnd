package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func CreateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 5) //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_TOKEN")))

}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	c, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	token := c.Value
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

//Pretty display the claims licely in the terminal
func Pretty(data interface{}) {
	_, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := TokenValid(r)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, "invalid token")
			return
		}
		next(w, r)
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
