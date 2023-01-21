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
	Category    string
}

type Items struct {
	Category string
	Dishes   []Item
}

var data []Item = []Item{
	{1, "Ceasar", "anchovies, olive oil, lemon juice, egg, and Parmesan cheese, garnished with croutons", 1750, "Salads"},
	{2, "Caprese", "Sliced mozzarella. Sliced tomatoes. Sweet basil. Olive oil. Salt.", 1250, "Salads"},
	{3, "Cobb", "Chopped salad greens, tomato, crispy bacon, chicken breast, hard-boiled eggs, avocado, chives, Roquefort cheese and red wine vinaigrette.", 950, "Salads"},
	{4, "Margarita", "tomatoes, mozzarella cheese, garlic, fresh basil, and extra-virgin olive oil", 2950, "Pizzas"},
	{5, "Sicilian", "dough topped with mozzarella cheese and tomato sauce.", 2350, "Pizzas"},
	{6, "Quattro formaggi", "topped with a combination of four kinds of cheese", 3150, "Pizzas"},
	{7, "Gazpacho", "tomatoes, garlic, onions, pepper and olive oil", 1750, "Soups"},
	{8, "Tom Yum", "several spices and herbs, including lemongrass, galangal and kaffir lime leaves.", 2150, "Soups"},
	{9, "Ramen", "broth based on chicken, seasoned with taré and served with pasta", 1450, "Soups"},
	{10, "Bolognese", "spaghetti and a sauce made of minced beef, tomatoes, onion, bacon, red wine and herbs", 2390, "Pasta"},
	{11, "Fettuccine Alfredo", "fettuccine (flat pasta ribbons) tossed with parmesan cheese and butter.", 2290, "Pasta"},
	{12, "Carbonara", "spaghetti with a cream-based sauce with ham or pancetta", 2490, "Pasta"},
}

func getData(data []Item, count int) []Items {
	items := make([]Items, count+1)
	for i := 0; i < len(data); i++ {
		for j := i; j > 0 && data[j-1].Category > data[j].Category; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
	j := 0
	for i := 0; i < len(data); i++ {
		if data[i].Category != items[j].Category {
			j++
		}
		items[j].Category = data[i].Category
		items[j].Dishes = append(items[j].Dishes, Item{data[i].Id, data[i].Dish, data[i].Description, data[i].Price, data[i].Category})
	}
	return items[1:]
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
	arr := getData(data, 4)
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
