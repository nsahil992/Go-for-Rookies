package main

import "fmt"

// Function to change the king
func suggestKing(suggestion *string) {
	fmt.Println("Vidur suggests a new king...")
	*suggestion = "Pandu" // Dereferencing the pointer to change the value at the address
}

func main() {
	// Initializing the king variable
	king := "Dhritarashtra"
	fmt.Println("King to become is:", king)

	// Passing the address of king to the function
	suggestKing(&king)

	// Final value of king after Vidur’s suggestion
	fmt.Println("The king of Hastinapur is:", king)
}
