package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	// this places the executed file in the default servemux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("name").Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
		t.Execute(w, nil)

		t1 := template.Must(template.New("name1").Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
		t1.Execute(w, nil)

		t2 := template.Must(template.New("name2").Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
		t2.Execute(w, nil)
	})

	log.Fatalln(http.ListenAndServe(":9000", nil))
}
