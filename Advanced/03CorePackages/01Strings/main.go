package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains :- If the string contains a specific word

	str := "Go is awesome"

	if strings.Contains(str, "awesome") {
		fmt.Println("string contains \"awesome\"")
	} else {
		fmt.Println("string not contains \"awesome\"")
	}

	if strings.Contains(str, "fantastic") {
		fmt.Println("string contains \"fantastic\"")
	} else {
		fmt.Println("string not contains \"fantastic\"")
	}

	// Count :- How many times the word appeared in a string

	str2 := "Go is very powerful and best programming language. I love Go lang."

	count := strings.Count(str2, "Go")
	fmt.Printf("The substring 'GO' appears %d times.\n", count)

	count2 := strings.Count(str, "Python")
	fmt.Printf("The substring 'Python' appears %d times in the string.\n", count2)
}
