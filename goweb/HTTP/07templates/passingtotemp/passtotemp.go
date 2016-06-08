package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, []int{1, 1, 2, 3, 5, 8, 13})
	if err != nil {
		log.Fatalln(err)
	}
}
