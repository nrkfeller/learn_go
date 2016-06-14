package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// this places the executed file in the default servemux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tEmpty := template.New("template test")
		tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n")) //empty pipeline following if
		tEmpty.Execute(w, nil)

		tWithValue := template.New("template test")
		tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n")) //non empty pipeline following if condition
		tWithValue.Execute(w, nil)

		tIfElse := template.New("template test")
		tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n")) //non empty pipeline following if condition
		tIfElse.Execute(w, nil)
	})

	log.Fatalln(http.ListenAndServe(":9000", nil))
}
