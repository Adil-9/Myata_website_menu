package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("Myata_premium_session_key")))

type Position struct {
	Id          int
	Dish        string
	Description string
	Price       int
	Category    string
}

type Items struct {
	Category string
	Dishes   []Position
}

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

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		errorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	tmpl, err := template.ParseFiles("./static/templates/main.html")
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	arr := getData(data(), count())
	if err := tmpl.Execute(w, arr); err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func count() int {
	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		fmt.Println("No database found")
		// panic(err)
	}
	defer db.Close()
	row, err := db.Query("SELECT COUNT(DISTINCT category) FROM menu")
	if err != nil {
		fmt.Println("Error getting data from database")
		// panic(err)
	}
	defer row.Close()
	count := 0
	for row.Next() {
		var err = row.Scan(&count)
		if err != nil {
			fmt.Println("Error scanning data from database")
			// panic(err)
		}
	}
	return count
}

func data() []Position {
	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		fmt.Println("No database found")
		// panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM menu")
	if err != nil {
		fmt.Println("Error getting data from database")
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
			var err = rows.Scan(&id, &price, &dish, &descr, &categ)
			if err != nil {
				fmt.Println("Error scanning data from database")
				// panic(err)
			}
			pos := Position{id, dish, descr, price, categ}
			data = append(data, pos)
		}
	}
	return data
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

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl, err := template.ParseFiles("./static/templates/login.html")
	if err != nil {
		panic(err.Error())
	}
	tmpl.Execute(w, nil)
}

func Check_login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	login := r.Form.Get("login")
	passwd := r.Form.Get("password")

	db, err := sql.Open("sqlite3", "admin.db")
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("SELECT password FROM admin WHERE login=? LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(login)
	if err != nil {
		panic(err.Error())
	}
	var password string
	for rows.Next() {
		rows.Scan(&password)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwd)); err != nil {
		fmt.Println("The password is incorrect")
		http.Redirect(w, r, "login", http.StatusSeeOther)
	} else {
		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		session, _ := store.Get(r, "admin")
		// Set some session values.
		session.Values["UserID"] = login
		session.Values["Id"] = 1
		// Save it before we write to the response/return from the handler.
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		}
	}
}

func Editing(w http.ResponseWriter, r *http.Request) {

	sessions, _ := store.Get(r, "admin")
	value, ok := sessions.Values["UserID"]
	fmt.Println("ok:", value, " Id:", sessions.Values["Id"])

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		fmt.Println("No database found")
		// panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM menu")
	if err != nil {
		fmt.Println("Error getting data from database")
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
			var err = rows.Scan(&id, &price, &dish, &descr, &categ)
			if err != nil {
				fmt.Println("Error scanning data from database")
				// panic(err)
			}
			pos := Position{id, dish, descr, price, categ}
			data = append(data, pos)
		}
	}

	tmpl, err := template.ParseFiles("./static/templates/editing_page.html")
	if err != nil {
		fmt.Println("Error parsing template")
		// panic(err)
	}

	if len(data) == 0 {
		tmpl.Execute(w, nil)
	} else {
		tmpl.Execute(w, data)
	}
}

func Insertion_page(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "admin")
	value, ok := sessions.Values["UserID"]
	fmt.Println("ok:", value, " Id:", sessions.Values["Id"])

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Redirect(w, r, "/editing", http.StatusSeeOther)
	} else {
		tmpl, err := template.ParseFiles("./static/templates/insert.html")
		if err != nil {
			panic(err.Error())
		}
		tmpl.Execute(w, nil)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "admin")
	value, ok := sessions.Values["UserID"]
	fmt.Println("ok:", value, " Id:", sessions.Values["Id"])

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Redirect(w, r, "/editing", http.StatusSeeOther)
		return
	} else {
		r.ParseForm()
		price, err := strconv.Atoi(r.Form.Get("price"))
		if err != nil {
			fmt.Println("In price non integer value was passed")
		}
		dish := r.Form.Get("dish")
		desc := r.Form.Get("description")
		categ := r.Form.Get("category")

		if price == 0 || dish == "" || desc == "" || categ == "" {
			if price == 0 {
				fmt.Println("price is not passed")
			}
			if dish == "" {
				fmt.Println("dish is not passed")
			}
			if desc == "" {
				fmt.Println("decsription is not passed")
			}
			if categ == "" {
				fmt.Println("categ is not passed")
			}
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		} else {
			db, err := sql.Open("sqlite3", "./menu.db")
			if err != nil {
				fmt.Println("Error opening database")
				// panic(err)
			}
			defer db.Close()

			// Insert an integer value
			stmt, err := db.Prepare("INSERT INTO menu (price, dish, description, category) VALUES(?,?,?,?)")
			if err != nil {
				fmt.Println("Error prepareing database")
				// panic(err)
			}
			defer stmt.Close()

			res, err := stmt.Exec(price, dish, desc, categ) //here parse from html and add
			if err != nil {
				fmt.Println("Error inserting values to database")
				// panic(err)
			}

			fmt.Printf("Inserted value %v\n", res)
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		}
	}
}

func Change(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "admin")
	value, ok := sessions.Values["UserID"]
	fmt.Println("ok:", value, " Id:", sessions.Values["Id"])

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	dish := r.Form.Get("dish")
	price, err1 := strconv.Atoi(r.Form.Get("price"))
	category := r.Form.Get("category")
	description := r.Form.Get("description")

	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		fmt.Println("Error opening database")
	}
	defer db.Close()

	if err1 == nil {
		stmt, err := db.Prepare("UPDATE menu SET price=? WHERE id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(price, id)
		if err != nil {
			panic(err)
		}
	}

	if description != "" {
		stmt, err := db.Prepare("UPDATE menu SET description=? WHERE id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(description, id)
		if err != nil {
			panic(err)
		}
	}

	if dish != "" {
		stmt, err := db.Prepare("UPDATE menu SET dish=? WHERE id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(dish, id)
		if err != nil {
			panic(err)
		}
	}

	if category != "" {
		stmt, err := db.Prepare("UPDATE menu SET category=? WHERE id=?")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(category, id)
		if err != nil {
			panic(err)
		}
	}
	http.Redirect(w, r, "/editing", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	sessions, _ := store.Get(r, "admin")
	value, ok := sessions.Values["UserID"]
	fmt.Println("ok:", value, " Id:", sessions.Values["Id"])

	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var price int
	var dish string
	var descr string
	var categ string
	if r.Method != "POST" {
		http.Redirect(w, r, "/editing", http.StatusSeeOther)
	} else {
		r.ParseForm()
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		//Opening database
		db, err := sql.Open("sqlite3", "./menu.db")
		if err != nil {
			fmt.Println("Error accessign database")
			// panic(err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT price, dish, description, category FROM menu WHERE id=?", id)
		if err != nil {
			fmt.Println("Error getting data from database")
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
			// panic(err)
		}
		defer rows.Close()

		for rows.Next() {
			var err = rows.Scan(&price, &dish, &descr, &categ)
			if err != nil {
				fmt.Println("Error scanning data from database")
				// panic(err)
			}
		}
		data := Position{id, dish, descr, price, categ}
		tmpl, _ := template.ParseFiles("./static/templates/edit.html")
		tmpl.Execute(w, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/editing", http.StatusSeeOther)
	} else {
		r.ParseForm()
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		// id, err := strconv.Atoi(r.Form.Get("id"))
		// if err != nil {
		// 	fmt.Println("In id non integer was passed")
		// }

		if id == 0 {
			fmt.Println("id is not passed")
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		} else {

			db, err := sql.Open("sqlite3", "./menu.db")
			if err != nil {
				panic(err)
			}
			defer db.Close()

			stmt, err := db.Prepare("DELETE FROM menu WHERE id=?")
			if err != nil {
				panic(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(id)
			if err != nil {
				panic(err)
			}
			redo_id()
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		}
	}
}

func redo_id() {
	db, err := sql.Open("sqlite3", "./menu.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE new_menu AS SELECT id, price, dish, description, category FROM menu")
	if err != nil {
		fmt.Println("REDOING Id WENT WRONG, CREATING NEW TABLE WENT WRONG")
	}
	_, err = db.Exec("DELETE FROM menu")
	if err != nil {
		fmt.Println("REDOING Id WENT WRONG, DELETING FROM TABLE WENT WRONG")
	}
	// _, err = db.Exec("DBCC CHECKIDENT (menu, RESEED, 0)")
	// if err != nil {
	// 	fmt.Println("REDOING Id WENT WRONG1")
	// 	fmt.Print(err)
	// }
	_, err = db.Exec("INSERT INTO menu (price, dish, description, category) SELECT price, dish, description, category FROM new_menu ORDER BY id ASC")
	if err != nil {
		fmt.Println("REDOING Id WENT WRONG2")
	}
	_, err = db.Exec("DROP TABLE new_menu")
	if err != nil {
		fmt.Println("REDOING Id WENT WRONG, DELETING WENT WRONG")
	}
	// fmt.Println("END OF REDOING Id")
}

func create_admin() {
	// Open database connection
	db, err := sql.Open("sqlite3", "./admin.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create the table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS admin (
		id INTEGER PRIMARY KEY,
		login TEXT,
		password TEXT
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully")

	stmt, err := db.Prepare("INSERT INTO admin (login, password) VALUES (?,?)")
	if err != nil {
		panic(err)
	}
	login := "admin"
	password := "myata_admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("Hash to store:", string(hash))

	if err != nil {
		// TODO: Properly handle error
		panic(err.Error())
	}
	stmt.Exec(login, hash)
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

// 	fmt.Println("Table created successfully")
// }
