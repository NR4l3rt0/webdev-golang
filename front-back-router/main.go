package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func greetUser(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	xs := strings.Split(p, "/")
	username := xs[len(xs)-1]

	tpl := template.Must(template.ParseFiles("helloUser.gohtml"))
	err := tpl.Execute(w, username)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
		return
	}
	return
}

func main() {
	http.HandleFunc("/", greetUser)
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
