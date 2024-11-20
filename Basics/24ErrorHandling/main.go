package main

import (
	"fmt"
)

func main() {
	var num1, num2 int
	fmt.Println("Enter your first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Error occured while reading the first number: ", err)
		return
	}

	fmt.Println("Enter your second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Error occurred while reading the second number: ", err)
		return
	}

	if num2 == 0 {
		fmt.Println("Error: Cannot divide by zero")
	} else {
		result := num1 / num2
		fmt.Println(result)
	}
}
