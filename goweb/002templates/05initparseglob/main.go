package main

import (
	"html/template"
	"os"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	tpls.ExecuteTemplate(os.Stdout, "potato.gohtml", nil)
}
