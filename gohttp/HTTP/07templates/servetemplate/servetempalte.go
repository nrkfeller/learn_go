package main

import (
	"html/template"
	"log"
	"net/http"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// http://localhost:9000/tpl.gohtml
	http.ListenAndServe(":9000", nil)

}
