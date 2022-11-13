package main

import (
	"io"
	"log"
	"net/http"
)

var visitCounter int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { io.WriteString(w, `<a href="shopping">Go shopping!</a>`) })
	http.HandleFunc("/shopping", shopping)
	http.HandleFunc("/insight", insight)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Println("Starting server...")

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func shopping(w http.ResponseWriter, r *http.Request) {

	myCookies := r.Cookies()
	isSame := isSameUser(myCookies)

	if isSame {
		visitCounter++
		log.Println("Visit number: ", visitCounter)
	} else {
		c := &http.Cookie{
			Name:  "my-visit",
			Value: "shopping here",
			//Domain:  "localhost.com",
			Path: "/",
			//Expires: time.Now().Add(30 * time.Minute),
		}
		http.SetCookie(w, c)

		w.Header().Set("Location", "/insight")
		w.WriteHeader(http.StatusSeeOther)
	}
}

func isSameUser(cs []*http.Cookie) bool {
	var isLoggedIn bool
	for _, v := range cs {
		if v.Name == "my-visit" {
			isLoggedIn = true
			break
		}
	}
	return isLoggedIn
}

func insight(w http.ResponseWriter, r *http.Request) {

	cSniff := &http.Cookie{
		Name:  "my-sniff",
		Value: "checking here",
		//Domain:  "just-checking.com",
		Path: "/",
		//Expires: time.Now().Add(30000 * time.Minute),
	}
	http.SetCookie(w, cSniff)

	w.Header().Set("Tampered", "maybe")
	io.WriteString(w, "See you there!")
}

func expire(w http.ResponseWriter, r *http.Request) {

	cVisit, err := r.Cookie("my-visit")
	if err != http.ErrNoCookie {
		cVisit.MaxAge = -1
	}
	cSniff, err := r.Cookie("my-sniff")
	if err != http.ErrNoCookie {
		cSniff.MaxAge = -1
	}

	http.SetCookie(w, cVisit)
	http.SetCookie(w, cSniff)

	http.Redirect(w, r, "/", http.StatusOK)
}
