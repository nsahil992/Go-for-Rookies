package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func coffeeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Here's your coffee! ☕️")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func snacksHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Here is a free cookie for you! 🍪")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the Cafe! Choose /coffee or /snacks")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homepage)
	r.HandleFunc("/coffee", coffeeHandler)
	r.HandleFunc("/snacks", snacksHandler)

	fmt.Println("Cafe is open at port 9090...")
	err := http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
