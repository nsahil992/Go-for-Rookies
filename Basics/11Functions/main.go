package main
import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("Greetings for the day, %s\n", name)
}

func add(a int, b int) {
	sum := a + b
	fmt.Printf("The sum of %d and %d is: %d\n",a,b,sum)
}

func auth() {

	passSave := []string{}
	fmt.Println("Enter your password:")
	var password string 
	fmt.Scanln(&password)

	passSave = append(passSave, password)
	fmt.Println(passSave)

}

func main() {
	greet("Sahil")
	add(1, 2)
	auth()
}