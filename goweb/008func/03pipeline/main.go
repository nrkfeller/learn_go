package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"dub": double,
	"sq":  square,
	"rt":  sqRoot,
}

func double(x int) int {
	return x + x
}

func square(x int) int {
	return x * x
}

func sqRoot(x int) int {
	return x / 2
}

// the output of one thing becomes the input of the next thing. this is called pipelining
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}
}
