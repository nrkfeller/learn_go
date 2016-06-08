package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Capital are important!
	err = tpl.Execute(os.Stdout, Page{
		Title: "Nick Title",
		Body:  "Works as a swe and researcher",
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// Page is injected into web template
type Page struct {
	Title string
	Body  string
}
