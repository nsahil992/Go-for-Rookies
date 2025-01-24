package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{3, 8, 9, 2, 1, 4, 10, 6}
	fmt.Println("Unsorted numbers:", numbers)

	sort.Ints(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Sort with strings

	vars := []string{"Learning", "Go", "over", "Python"}
	fmt.Println("Unsorted variables:", vars)

	sort.Strings(vars)
	fmt.Println("Sorted variables:", vars)
}
