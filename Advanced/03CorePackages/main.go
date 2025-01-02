package main

import (
	"fmt"
	"strings"
)

func main() {
	learning := "Learning standard library in Go"
	fun := "Library in go"

	result := strings.Contains(learning, fun)
	fmt.Println(result)
}
