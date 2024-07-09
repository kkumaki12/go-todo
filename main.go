package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello, World!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDb()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.ListenAndServe(":8080", nil)
}
