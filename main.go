package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// this serves the index page to test the game on
	http.HandleFunc("/", mainHandler)

	// this gets the libraries for your game
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.Handle("/frontend/", http.StripPrefix("/frontend/", http.FileServer(http.Dir("./frontend"))))
	http.Handle("/backend/", http.StripPrefix("/backend/", http.FileServer(http.Dir("./backend"))))
	// listen and serve
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	//cwd, _ := os.Getwd()

	//p := filepath.Join(cwd, "./index.html")
	//fmt.Printf("%v", p)
	t := template.Must(template.ParseFiles("./frontend/index.html"))

	if err := t.Execute(w, nil); err != nil {
		log.Print(err.Error())
	}
}
