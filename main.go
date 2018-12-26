package main

import (
	"net/http"
)

// Item ...
type Item struct {
	name, image, category string
	price                 float32
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
	w.Write([]byte("items page"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", users)
	http.HandleFunc("/items", items)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
