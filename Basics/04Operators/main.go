package main
import (
	"fmt"
)

func main() {

	// Comparison Operators
	var num1 = 18
	var num2 = 19
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 > num2)
	fmt.Println(!(num1 != num2))

	// Arithmetic Operators
	var a, b string = "foo", "bar"
	fmt.Println(a + b)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
	fmt.Println(num1 / num2)
	fmt.Println(num1 % num2)

	// Logical Operator
	var x int = 10
	fmt.Println((x < 100) && (x > 100))
	fmt.Println((x < 100) || (x > 100))
	fmt.Println(!(x < 100))
}