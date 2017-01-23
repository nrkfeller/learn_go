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
	mathiens := []string{"Gogel", "Turing", "Gauss"}
	tpl.Execute(os.Stdout, mathiens)
}
