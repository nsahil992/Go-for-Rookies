package main
import (
	"fmt"
)

func modify(s *string) {
	*s = "hello"
}

func main() {
	a := "world"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)
}