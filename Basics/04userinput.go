package main
import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d",&age)
	fmt.Println("My age is",age)

	// MULTIPLE INPUT IN GO

	var a string 
	var b int
	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d",&a, &b)
	fmt.Println("Count: ", count)
	fmt.Println("Error: ", err)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}