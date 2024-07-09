package main

import "net/http"

func newHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new.html")
}
