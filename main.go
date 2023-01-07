package main

import(
	"net/http"
	"fmt"
	"myata_website_menu/render"
)

func main() {
	http.HandleFunc("/", render.Homepage)
	fmt.Println("Server running on: http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}