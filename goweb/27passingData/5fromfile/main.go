package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", formFile)
	http.ListenAndServe(":8080", nil)
}

func formFile(w http.ResponseWriter, r *http.Request) {
	var s string
	//	fmt.Println(r.Method)

	if r.Method == http.MethodPost {

		//open
		f, _, err := r.FormFile("q")
		if err != nil {
			log.Fatalln(err)
			return
		}
		defer f.Close()

		//read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
