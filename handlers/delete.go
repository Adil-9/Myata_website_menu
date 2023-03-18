package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
