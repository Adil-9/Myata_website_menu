package main

import (
	"fmt"
	"myata_website_menu/render"
	"net/http"
)

func main() {
	http.HandleFunc("/", render.HomePage)
	http.HandleFunc("/cart", render.CartPage)
	fmt.Println("Server running on: http://localhost:3000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}
