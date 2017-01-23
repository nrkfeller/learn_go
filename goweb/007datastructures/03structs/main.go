package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type mathiens struct {
	Name     string
	Invented string
}

func main() {
	// same with maps
	kfg := mathiens{"Karl Fredrich Gauss", "normal distro"}
	kg := mathiens{"Kurt Godel", "Computational logic"}
	le := mathiens{"Leohart Euler", "Euler's Laws"}

	allm := []mathiens{kfg, kg, le}

	tpl.Execute(os.Stdout, allm)
}
