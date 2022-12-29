package main

import(
	"net/http"
	"fmt"
	"myata_website_menu/render"
)

func main() {
	http.HandleFunc("/", render.Homepage)
	fmt.Println("Server running on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}