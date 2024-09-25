package main
import (
	"fmt"
)

func main() {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Welcome onboard, %s!\n",name)

	fmt.Print("Enter your email: ")
	var email string
	fmt.Scan(&email)
	fmt.Printf("Hey %s, please confirm your email: %s",name,email)
}