package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// this places the executed file in the default servemux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs := r.URL.Path
		err = tpl.Execute(w, fs)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
