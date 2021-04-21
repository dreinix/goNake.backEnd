package score

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	score Score
)

func getAllScore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Inix")
	}
}

func addScore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&score)
		msg := "You got " + strconv.Itoa(score.Value)
		json.NewEncoder(w).Encode(msg)
	}
}
