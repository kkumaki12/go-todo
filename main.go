package main

import (
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello, World!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
