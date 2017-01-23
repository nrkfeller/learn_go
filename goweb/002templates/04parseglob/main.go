package main

import (
	"html/template"
	"os"
)

func main() {
	tpls, _ := template.ParseGlob("templates/*.gohtml")

	tpls.Execute(os.Stdout, nil)
	tpls.ExecuteTemplate(os.Stdout, "potato.gohtml", nil)
}
