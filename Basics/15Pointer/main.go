package main

import "fmt"

func main() {
	// Initializing the king variable
	king := "Dhritarashtra"
	fmt.Println("King to become is:", king)

	// Declaring a pointer variable vidur, which will point to the address of king
	var vidur *string
	vidur = &king // vidur now stores the address of the king variable

	// Printing the address where vidur points to (the address of the king)
	fmt.Printf("Vidur comes to know that the king of hastinapur (%v) will be %v\n", vidur, king)

	// Vidur suggests a new king by changing the value at the address
	fmt.Println("Vidur suggests, Pandu should be king of Hastinapur")
	*vidur = "Pandu" // Dereferencing vidur (changing the value at that address)

	// Final value of king after Vidur’s suggestion
	fmt.Println("The king of Hastinapur becomes ", king)
}
