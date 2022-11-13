package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("stuff"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, `<a href="/resources"> Go to page</a>`)
	})

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
