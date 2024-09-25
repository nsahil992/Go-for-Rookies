package main
import (
	"fmt"
)

func main() {
	var cartoons [5]string = [5]string{"Hagemaru", "Courage the cowardly dog", "Phineas and Ferb", "Ben10", "Mr.Bean"}
	fmt.Println(cartoons)

	marks := [3]int{10, 20, 30}   // shorthand way
	fmt.Println(marks)

	languages := [...]string{"Swift", "Golang", "Julia", "Java"}
	fmt.Println(languages)
	fmt.Println(len(languages))
	fmt.Println(languages[0])

	marks[0] = 99
	fmt.Println(marks) // [99 20 30]

	// Looping through array
	var grades [4]int = [4]int{87, 99, 100, 72}
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}

	for index, element := range grades {
		fmt.Println(index, "=>", element)
	} 

	// Multidimensional array
	arr := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	// [3] refers to the size of the array
	// [2] refers to each element in a array
	fmt.Println(arr[0][1])  // prints 20
}