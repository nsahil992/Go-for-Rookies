package main
import "fmt"

func main() {
	var name string = "Sahil"
	var age int = 21

	fmt.Printf("Hello, my name is %v and I am %d years old\n", name, age)

	// shorthand method

	pi := 3.14
	fmt.Println("The value of pi is: ",pi)

	// 2 variable declaration of same type

	var x,y string = "Go", "Mascot"   
	fmt.Println(x)
	fmt.Println(y)

	// Shorthand way of two different variables
	var (
		a string = "Gopher"      
		i int = 8999
	)
	fmt.Println(a)
	fmt.Println(i)
}