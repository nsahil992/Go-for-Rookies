package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Starting main function")
	defer fmt.Println("This is a deferred function")
	fmt.Println("This will print first")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File is opened and ready for operations.")
}
