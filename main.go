package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome</h1>")

	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<h1>Contact <a href=\"mailto:#\">mail us</a></h1>")
	}

}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3030", nil)
}
