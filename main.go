package main

import (
	"fmt"
	"myata_website_menu/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePage)
	fmt.Println("Server running on: http://localhost:3000")
	http.Handle("/static/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":3000", nil)
}
