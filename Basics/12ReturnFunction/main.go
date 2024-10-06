package main
import (
	"fmt"
)

func operation() (int, int) {
	fmt.Println("Enter your first number:")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Enter your second number:")
	var num2 int
	fmt.Scanln(&num2)

	sum := num1 + num2
	diff := num1 - num2

	fmt.Print("The sum and difference is: ")
	return sum, diff
} 

// Multiple return values 

func getLanguages() (string, string, string) {
	return "Go Lang", "Swift", "Julia"
}

func main() {
	fmt.Println(operation())
	fmt.Println(getLanguages())
}