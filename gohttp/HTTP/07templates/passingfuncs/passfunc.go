package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	var err error

	tpl := template.New("tpl.gohtml")

	// this just gives the template the right to call another func
	tpl = tpl.Funcs(template.FuncMap{
		"myfunc": func() string {
			return "passed text from myfunc"
		},
		"myotherfunc": func() string {
			return "more text from other func"
		},
	})

	// tpl = tpl.Funcs(template.FuncMap{
	// 	"myotherfunc": func() string {
	// 		return "more text from other func"
	// 	},
	// })

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
