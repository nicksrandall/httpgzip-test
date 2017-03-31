package main

import (
	"net/http"

	"github.com/shurcooL/httpgzip"
)

func main() {
	// this will give infinate redirects
	http.Handle("/static/", http.StripPrefix("/static/", httpgzip.FileServer(http.Dir("static"), httpgzip.FileServerOptions{})))

	// this works
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}
