package main

import (
	"fmt"
	"myata_website_menu/render"
	"net/http"
)

func main() {
	http.HandleFunc("/", render.HomePage)
	http.HandleFunc("/cart", render.CartPage)
	fmt.Println("Server running on: http://localhost:3000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":3000", nil)
}

// import (
//   "database/sql"
//   "fmt"

//   _ "github.com/mattn/go-sqlite3"
// )

// func main() {
//   // Open database connection
//   db, err := sql.Open("sqlite3", "./my.db")
//   if err != nil {
//     panic(err)
//   }
//   defer db.Close()

//   // Create the table
//   sqlStmt := `
//   CREATE TABLE IF NOT EXISTS mytable (
//     id INTEGER PRIMARY KEY,
//     value INTEGER,
//     dish TEXT,
//     description TEXT
//   );`
//   _, err = db.Exec(sqlStmt)
//   if err != nil {
//     panic(err)
//   }

//   fmt.Println("Table created successfully")

//   // Insert an integer value
//   stmt, err := db.Prepare("INSERT INTO mytable (value, dish, description) VALUES(?,?,?)")
//   if err != nil {
//     panic(err)
//   }
//   defer stmt.Close()

//   res, err := stmt.Exec(10, "apple", "is an apple")
//   if err != nil {
//     panic(err)
//   }

//   fmt.Printf("Inserted value %v\n", res)

//   // Query the table
//   rows, err := db.Query("SELECT * FROM mytable")
//   if err != nil {
//     panic(err)
//   }
//   defer rows.Close()

//   // Iterate over the rows and print out the values
//   for rows.Next() {
//     var id int
//     var value int
//     var dish string
//     var dis string
//     err = rows.Scan(&id, &value, &dish, &dis)
//     if err != nil {
//       panic(err)
//     }

//     fmt.Printf("id: %d, value: %d\n, dish: %v, dis: %v\n", id, value, dish, dis)
//   }
//}
