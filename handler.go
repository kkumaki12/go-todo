package main

import (
	"net/http"
	"time"
)

func newHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/new.html")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")
	_, err := db.Exec("INSERT INTO todos (content, created_at) VALUES (?, ?)", content, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
