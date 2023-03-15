package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
