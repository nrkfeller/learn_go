package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// this places the executed file in the default servemux
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fs := strings.Split(req.URL.Path, "/")
		err = tpl.Execute(res, fs[len(fs)-1])
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.ListenAndServe(":9000", nil)
}
