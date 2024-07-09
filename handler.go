package main

import (
	"net/http"
	"text/template"
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	todos := []Todo{
		{ID: 1, Content: "Write a blog", CreatedAt: time.Now()},
		{ID: 2, Content: "Take a walk", CreatedAt: time.Now()},
	}

	tmpl.Execute(w, todos)
}
