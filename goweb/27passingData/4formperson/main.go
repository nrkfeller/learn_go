package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", form)
	http.ListenAndServe(":8080", nil)
}

func form(w http.ResponseWriter, r *http.Request) {
	fn := r.FormValue("first")
	ln := r.FormValue("last")
	sub := r.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{fn, ln, sub})
	if err != nil {
		log.Fatalln(err)
	}
}
