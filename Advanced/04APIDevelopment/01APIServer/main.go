package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello User form the handler package")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Cafe is open at Port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
}
