package main
import (
	"fmt"
)

func main() {
	var rank int
	fmt.Println("Enter your rank: ")
	fmt.Scanln(&rank)

	if(rank == 1) {
		fmt.Println("Congratulation! You won a GOLD medal 🥇")
	} else if (rank == 2) {
		fmt.Println("Congratulations! You won a SILVER medal 🥈")
	} else if (rank == 3) {
		fmt.Println("Congratulations! You won a BRONZE medal 🥉")
	} else {
		fmt.Println("Congratulations for your effortss")
	}
}