package main

import "fmt"

// Vidur attempt to change the king
func suggestKing(suggestion string) {
	fmt.Println("Vidur suggests a new king...")
	suggestion = "Yudhistira" // This change affects only the copy, not the original variable
	fmt.Println("Inside Vidur's mind, king is:", suggestion)
}

func main() {
	// Initializing the king as Dhritarashtra
	king := "Dhritarashtra"
	fmt.Println("King to become is:", king)

	// Passing the value of king (not its address) to the function
	suggestKing(king)

	// Final value of king even after Vidur’s suggestion
	fmt.Println("The king of Hastinapur is still:", king)
}
