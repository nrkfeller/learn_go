package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", get)
	http.ListenAndServe(":8080", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
