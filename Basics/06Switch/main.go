package main
import (
	"fmt"
)

func main() {
	fmt.Println("Enter your phone model: ")
	var phone string
	fmt.Scan(&phone)

	switch phone {
	case "iPhone":
		fmt.Println("You are using iOS")
	case "Blackberry":
		fmt.Println("You are using Blackberry")
	case "Nokia Lumia":
		fmt.Println("You are using Windows")
	default:
		fmt.Println("You are usign Android")
	}

	// fallthrough

	var rank int = 1

	switch rank {
	case 1:
		fmt.Println("Congratulation! You won a GOLD medal 🥇")
		fallthrough
	case 2:
		fmt.Println("Congratulations! You won a SILVER medal 🥈")
	case 3:
		fmt.Println("Congratulations! You won a BRONZE medal 🥉")
	default:
		fmt.Println("Congratulations for your efforts.")
	}
}