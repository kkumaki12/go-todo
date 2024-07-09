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
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.ListenAndServe(":8080", nil)
}
