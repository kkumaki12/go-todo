package main

import (
	"log"
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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
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

	rows, err := db.Query("SELECT id, content, created_at FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Content, &todo.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	tmpl.Execute(w, todos)
}
