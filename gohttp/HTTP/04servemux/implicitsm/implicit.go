/*

there is a default servemux that can be used insteads

*/

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/euro/", europe)
	http.HandleFunc("/afro/", africa)

	// http package uses a default servemux!
	http.ListenAndServe(":9000", nil)
}

func europe(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "EUROPE FTW")
}

func africa(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "AFRICA FTW")
}
