package main
import (
	"fmt"
)

func divide(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}

func main() {
	fmt.Println("Error handling in functions:")
	ans := divide(10, 4)
	fmt.Println("The division is: ",ans)
}