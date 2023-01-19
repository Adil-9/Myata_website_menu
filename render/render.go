package render

import (
	"fmt"
	"net/http"
	"text/template"
)

type Item struct {
	Id          int
	Dish        string
	Description string
	Price       int
	Category    string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("templates/main.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	arr := []Item{{1, "Ceasar", "anchovies, olive oil, lemon juice, egg, and Parmesan cheese, garnished with croutons", 1750, "Salads"}, {2, "Margarita", "tomatoes, mozzarella cheese, garlic, fresh basil, and extra-virgin olive oil", 2950, "Pizzas"}}
	if err := tmpl.Execute(w, arr); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 page not found")
	} else if status == http.StatusBadRequest {
		fmt.Fprint(w, "400 bad request")
	} else if status == http.StatusInternalServerError {
		fmt.Fprint(w, "500 internal server error")
	} else if status == http.StatusMethodNotAllowed {
		fmt.Fprint(w, "Method not allowed")
	}
}
