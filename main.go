package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Item ...
type Item struct {
	name     string
	image    string
	category string
	price    float32
}

// User ...
type User struct {
	name    string
	items   []Item
	balance float32
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users page"))
}

func items(w http.ResponseWriter, r *http.Request) {
	var items []Item
	w.Write([]byte("items page"))

	db, dberr := sql.Open("mysql", dbcredentials)
	checkErr(dberr)

	rows, err := db.Query("SELECT * FROM items")
	checkErr(err)

	for rows.Next() {
		var name string
		var image string
		var category string
		var price float32

		err := rows.Scan(&name, &image, &category, &price)
		checkErr(err)

		item := Item{name, image, category, price}
		fmt.Println(item.name)

		items = append(items, item)

		w.Write([]byte(item.name))
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", users)
	http.HandleFunc("/items", items)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
