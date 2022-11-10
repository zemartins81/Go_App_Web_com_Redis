package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contact", contactHandler).Methods("GET")
	r.HandleFunc("/about", aboutHandler).Methods("GET")
	r.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "This is index page!")
	if err != nil {
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "This is the contact page!")
	if err != nil {
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "This is about page!")
	if err != nil {
		return
	}
}
