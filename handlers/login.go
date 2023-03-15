package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

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

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "admin")
	delete(session.Values, "UserID")
	session.Save(r, w)
	http.Redirect(w, r, "login", http.StatusSeeOther)
}
