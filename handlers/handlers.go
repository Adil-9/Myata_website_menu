package handlers

import (
	"database/sql"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("Myata_premium_session_key")))

func getData(data []Position, count int) []Items {
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
		items[j].Dishes = append(items[j].Dishes, Position{data[i].Id, data[i].Dish, data[i].Description, data[i].Price, data[i].Category})
	}
	return items[1:]
}

func count() int {
	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		ErrorLog.Println("No database found")
		// panic(err)
	}
	defer db.Close()
	row, err := db.Query("SELECT COUNT(DISTINCT category) FROM menu")
	if err != nil {
		ErrorLog.Println("Error getting data from database")
		// panic(err)
	}
	defer row.Close()
	count := 0
	for row.Next() {
		err := row.Scan(&count)
		if err != nil {
			ErrorLog.Println("Error scanning data from database")
			// panic(err)
		}
	}
	return count
}

func data() []Position {
	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		ErrorLog.Println("No database found")
		// panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM menu")
	if err != nil {
		ErrorLog.Println("Error getting data from database")
		// panic(err)
	}
	defer rows.Close()

	data := []Position{}

	if err == nil {
		// Iterate over the rows and print out the values
		for rows.Next() {
			var id int
			var price int
			var dish string
			var descr string
			var categ string
			err := rows.Scan(&id, &price, &dish, &descr, &categ)
			if err != nil {
				ErrorLog.Println("Error scanning data from database")
				// panic(err)
			}
			pos := Position{id, dish, descr, price, categ}
			data = append(data, pos)
		}
	}
	return data
}

// var sampleData []Position = []Position{
// 	{1, "Ceasar", "anchovies, olive oil, lemon juice, egg, and Parmesan cheese, garnished with croutons", 1750, "Salads"},
// 	{2, "Caprese", "Sliced mozzarella. Sliced tomatoes. Sweet basil. Olive oil. Salt.", 1250, "Salads"},
// 	{3, "Cobb", "Chopped salad greens, tomato, crispy bacon, chicken breast, hard-boiled eggs, avocado, chives.", 950, "Salads"},
// 	{4, "Margarita", "tomatoes, mozzarella cheese, garlic, fresh basil, and extra-virgin olive oil", 2950, "Pizzas"},
// 	{5, "Sicilian", "dough topped with mozzarella cheese and tomato sauce.", 2350, "Pizzas"},
// 	{6, "Quattro formaggi", "topped with a combination of four kinds of cheese", 3150, "Pizzas"},
// 	{7, "Gazpacho", "tomatoes, garlic, onions, pepper and olive oil", 1750, "Soups"},
// 	{8, "Tom Yum", "several spices and herbs, including lemongrass, galangal and kaffir lime leaves.", 2150, "Soups"},
// 	{9, "Ramen", "broth based on chicken, seasoned with taré and served with pasta", 1450, "Soups"},
// 	{10, "Bolognese", "spaghetti and a sauce made of minced beef, tomatoes, onion, bacon, red wine and herbs", 2390, "Pasta"},
// 	{11, "Fettuccine Alfredo", "fettuccine (flat pasta ribbons) tossed with parmesan cheese and butter.", 2290, "Pasta"},
// 	{12, "Carbonara", "spaghetti with a cream-based sauce with ham or pancetta", 2490, "Pasta"},
// }

// func create_table() {
// 	// Open database connection
// 	db, err := sql.Open("sqlite3", "./menu.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	// Create the table
// 	sqlStmt := `
// 	CREATE TABLE IF NOT EXISTS menu (
// 		id INTEGER PRIMARY KEY,
// 		price INTEGER,
// 		dish TEXT,
// 		description TEXT,
// 		category TEXT
// 	);`
// 	_, err = db.Exec(sqlStmt)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ErrorLog.Println("Table created successfully")
// }
