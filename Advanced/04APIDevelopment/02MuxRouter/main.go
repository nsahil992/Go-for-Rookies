package main

import (
	"fmt"
	"net/http"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products = []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequest() {

	http.HandleFunc("/", homePage)
	http.ListenAndServe("localhost:10000", nil)
}

func main() {

	Products = []Product{
		{1, "Laptop", 10, 10000},
		{2, "Mobile", 10, 10000},
		{3, "Tablet", 10, 10000},
	}
	handleRequest()
}
