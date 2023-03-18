package handlers

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
