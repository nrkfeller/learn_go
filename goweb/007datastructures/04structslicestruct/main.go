package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type data struct {
	Animals []animal
	Fruits  []fruit
}

type animal struct {
	Kind string
	Food string
}

type fruit struct {
	Kind     string
	Quantity int
}

func main() {
	a := fruit{"Apple", 12}
	b := fruit{"banana", 6}
	c := animal{"cat", "whiskers"}
	d := animal{"dog", "pedigree"}

	all := data{
		[]animal{c, d},
		[]fruit{a, b},
	}
	tpl.Execute(os.Stdout, all)
}
