package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in Go lang")

	for i := 1; i < 5; i++ {
		fmt.Println("GOpher")
	}

	number := 5
	count := 1

	// Break

	for {
		var userInput string 
		fmt.Println("Press c to continue, any other letter to stop.")
		fmt.Scanln(&userInput)

		if userInput == "c" {
			fmt.Println(number * count)
			count ++
		} else {
			fmt.Println("Quitting.......")
			fmt.Println("Bye!")
			break
		}
	}

	// Continue

	for num := 1; num <= 5; num++ {

		if num == 3 {
			continue
		}
		fmt.Println(num)
	}
}
