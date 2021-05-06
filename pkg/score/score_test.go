package score

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dreinix/gonake/pkg/auth"
	"github.com/dreinix/gonake/pkg/database"
)

var scr Score

func TestUserModel(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	id := rand.Intn(161299)
	usr := auth.User{ID: 12, Name: "test", Username: "testusr"}

	scr = Score{ID: id, Value: 16, User: usr, Date: time.Date(99, 12, 16, 00, 00, 00, 00, time.Local)}
	//fmt.Println(id)
	if scr.ID != id {
		t.Errorf("create score 'test' failed, expected %v, got %v", id, scr.ID) // to indicate test failed
	}
	if scr.Value != 16 {
		t.Errorf("create score 'test' failed, expected %v, got %v", "test", scr.Value) // to indicate test failed
	}
	if scr.User.ID != usr.ID {
		t.Errorf("create score 'test' failed, expected %v, got %v", usr.ID, scr.User.ID) // to indicate test failed
	}
	if scr.Date.Month() != 12 {
		t.Errorf("create score 'test' failed, expected %v, got %v", "December", scr.Date.Month().String()) // to indicate test failed
	}

}
func TestGetAllScores(t *testing.T) {
	r, err := database.Conect("gonaketest")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/scores", nil)
	if err != nil {
		fmt.Println("db error")
		t.Fatal("backend error")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllScore(r))
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getAllUsers failed got %v expected status code: %v",
			http.StatusOK, status)
	}
}

func TestGetTop(t *testing.T) {
	r, err := database.Conect("gonaketest")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/scores/top", nil)
	if err != nil {
		fmt.Println("db error")
		t.Fatal("backend error")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getTop(r))
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getAllUsers failed got %v expected status code: %v",
			http.StatusOK, status)
	}
}
