package handlers

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
}

type Items struct {
	Category string
	Dishes   []Item
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
	salads := []Item {
		{1, "Ceasar", "anchovies, olive oil, lemon juice, egg, and Parmesan cheese, garnished with croutons", 1750},
		{2, "Caprese", "Sliced mozzarella. Sliced tomatoes. Sweet basil. Olive oil. Salt.", 1250},
		{3, "Cobb", "Chopped salad greens, tomato, crispy bacon, chicken breast, hard-boiled eggs, avocado, chives, Roquefort cheese and red wine vinaigrette.", 950},
	}
	pizzas := []Item {
		{4, "Margarita", "tomatoes, mozzarella cheese, garlic, fresh basil, and extra-virgin olive oil", 2950},
		{5, "Sicilian", "dough topped with mozzarella cheese and tomato sauce.", 2350},
		{6, "Quattro formaggi", "topped with a combination of four kinds of cheese", 3150},
	}
	soups := []Item {
		{7, "Gazpacho", "tomatoes, garlic, onions, pepper and olive oil", 1750},
		{8, "Tom Yum", "several spices and herbs, including lemongrass, galangal and kaffir lime leaves.", 2150},
		{9, "Ramen", "broth based on chicken, seasoned with taré and served with pasta", 1450},
	}
	pasta := []Item {
		{10, "Bolognese", "spaghetti and a sauce made of minced beef, tomatoes, onion, bacon, red wine and herbs", 2390},
		{11, "Fettuccine Alfredo", "fettuccine (flat pasta ribbons) tossed with parmesan cheese and butter.", 2290},
		{12, "Carbonara", "spaghetti with a cream-based sauce with ham or pancetta", 2490},
	}
	arr := []Items{
		{"Salads", salads},
		{"Pizzas", pizzas},
		{"Soups", soups},
		{"Pasta", pasta},
	}
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
