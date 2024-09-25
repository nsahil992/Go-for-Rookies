package main
import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World")
	}

	multiple := 5
	n := 1

	// Break
	for {
		var input string
		fmt.Println("Type c to continue, any other letter to stop")
		fmt.Scan(&input)

		if(input == "c") {
			fmt.Println(multiple * n)
			n++
		} else {
			fmt.Println("Bye!")
			break
		}
	}

	// Continue

	for num := 1; num <= 5; num++ {
		if(num == 3) {
			continue
		}
		fmt.Println(num)
	}
}