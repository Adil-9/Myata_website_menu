package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

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
			err := rows.Scan(&id, &price, &dish, &descr, &categ)
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

		// Opening database
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
			err := rows.Scan(&price, &dish, &descr, &categ)
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
