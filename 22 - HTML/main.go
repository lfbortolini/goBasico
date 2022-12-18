package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	templates = template.Must(template.ParseGlob("*.html"))
	u := usuario{"João", "joao@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}

func main() {
	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
