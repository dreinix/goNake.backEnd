package user

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/dreinix/gonake/pkg/database"
)

var usr User

func TestUserModel(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)
	id := rand.Intn(161299)
	usr = User{ID: id, Name: "test", Username: "testusr" + strconv.Itoa(id), Password: "testpwd", Status: "active"}
	//fmt.Println(id)
	if usr.ID != id {
		t.Errorf("create user 'test' failed, expected %v, got %v", id, usr.ID) // to indicate test failed
	}
	if usr.Name != "test" {
		t.Errorf("create user 'test' failed, expected %v, got %v", "test", usr.Name) // to indicate test failed
	}
	if usr.Username != "testusr"+strconv.Itoa(id) {
		t.Errorf("create user 'test' failed, expected %v, got %v", ("testusr" + strconv.Itoa(id)), usr.Username) // to indicate test failed
	}
	if usr.Password != "testpwd" {
		t.Errorf("create user 'test' failed, expected %v, got %v", "testpwd", usr.Password) // to indicate test failed
	}
	if usr.Status != "active" {
		t.Errorf("create user 'test' failed, expected %v, got %v", "active", usr.Status) // to indicate test failed
	}

}
func TestAddUser(t *testing.T) {
	r, err := database.Conect("gonaketest")
	if err != nil {
		t.Errorf("connect to goNakeTest_db test failed expected no error got %v", err)
	}
	body := fmt.Sprintf(`{"Username":"%s","Password":"%s" , "name": "%s"}`, usr.Username, usr.Password, usr.Name)
	var jsonStr = []byte(body)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addUser(r))
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("addUser failed got %v expected status code: %v",
			http.StatusOK, status)
		t.Log(rr.Body.String())
	}
}

func TestGetAllUsers(t *testing.T) {
	r, err := database.Conect("gonaketest")
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		fmt.Println("db error")
		t.Fatal("backend error")
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllUsers(r))
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getAllUsers failed got %v expected status code: %v",
			http.StatusOK, status)
	}
}
func TestLoginUser(t *testing.T) {
	r, err := database.Conect("gonaketest")
	if err != nil {
		t.Errorf("connect to goNakeTest_db test failed expected no error got %v", err)
	}
	body := fmt.Sprintf(`{"Username":"%s","Password":"%s"}`, usr.Username, usr.Password)
	var jsonStr = []byte(body)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logIn(r))
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("LoginUser failed got %v expected status code: %v",
			http.StatusOK, status)
		t.Log(rr.Body.String())
	}
}

// protected routes can't be tested
