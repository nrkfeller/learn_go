/*
This can come in handy but it might be best to just process on you side and pass in data static
separation of concerns
*/

package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func main() {
	var err error

	tpl := template.New("tpl.gohtml")

	tpl = tpl.Funcs(template.FuncMap{
		"upcase": func(s string) string {
			return strings.ToUpper(s)
		},
	})

	tpl, err = tpl.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: "nicks title",
		Body:  "Workds a sswe and ai ra",
	})
}

type Page struct {
	Title string
	Body  string
}
