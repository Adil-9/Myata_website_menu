package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

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

		if price == 0 || dish == "" || categ == "" {
			if price == 0 {
				fmt.Println("price is not passed")
			}
			if dish == "" {
				fmt.Println("dish is not passed")
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

			res, err := stmt.Exec(price, dish, desc, categ) // here parse from html and add
			if err != nil {
				fmt.Println("Error inserting values to database")
				// panic(err)
			}

			fmt.Printf("Inserted value %v\n", res)
			http.Redirect(w, r, "/editing", http.StatusSeeOther)
		}
	}
}
