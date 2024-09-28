package main
import (
	"fmt"
)

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Denominator must not be 0"
	}
	return a / b, "nil"
}

func getUserData() (string, int) {
	return "Sahil", 21
}

func main() {
	fmt.Println("Error handling in functions:")
	ans,_ := divide(10, 0)
	fmt.Println("The division is: ",ans)

	name,_ := getUserData()
	fmt.Println("Welcome,",name)
}