package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", 42)
}
