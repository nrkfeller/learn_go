package main

import (
	"io"
	"net/http"
)

// MyHandler is an int
type MyHandler int

func main() {
	var h MyHandler

	http.ListenAndServe(":9000", h)
}

func (h MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// takes in as html code
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch req.URL.Path {
	// user can define the url. not related to anything else in directory/system/etc
	case "/cat":
		io.WriteString(res, "<strong>miaw<strong>")
	case "/dog":
		io.WriteString(res, `<img src="http://r.ddmcdn.com/s_f/o_1/cx_633/cy_0/cw_1725/ch_1725/w_720/APL/uploads/2014/11/too-cute-doggone-it-video-playlist.jpg">`)
	}
}
