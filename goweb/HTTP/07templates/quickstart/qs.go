package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	// returns a pointer to a template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// where do you want the output to go, we can pass Name to {{.}}
	err = tpl.Execute(os.Stdout, "Nick Feller")
	if err != nil {
		log.Fatalln(err)
	}
}
