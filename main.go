package main

import (
	"fmt"
	"myata/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(handlers.HomePage))
	r.Handle("/editing", http.HandlerFunc(handlers.Editing))
	r.Handle("/insert", http.HandlerFunc(handlers.Insert))
	r.Handle("/edit/{id}", http.HandlerFunc(handlers.Edit))
	r.Handle("/insertion-page", http.HandlerFunc(handlers.Insertion_page))
	r.Handle("/change/{id}", http.HandlerFunc(handlers.Change))
	r.Handle("/delete/{id}", http.HandlerFunc(handlers.Delete))
	r.Handle("/check-login", http.HandlerFunc(handlers.Check_login))
	r.Handle("/login", http.HandlerFunc(handlers.Login))
	fmt.Println("Server running on: http://localhost:8000")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8000", r)
}
