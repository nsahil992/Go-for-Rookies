package main

import "fmt"

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hey %v, Welcome onboard\n", name)

	fmt.Println("Enter your email address: ")
	var email string
	fmt.Scanln(&email)
	fmt.Printf("Hey %v, please confirm your email address: %v\n",name,email)
}
