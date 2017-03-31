package main

import (
	"net/http"

	"github.com/shurcooL/httpgzip"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", httpgzip.FileServer(http.Dir("static"), httpgzip.FileServerOptions{})))
	http.ListenAndServe(":8080", nil)
}
