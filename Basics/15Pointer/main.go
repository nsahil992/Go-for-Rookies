// A pointer is a variable that holds memory address of another variable and also point where the memory is located

package main

import (
	"fmt"
)

func main() {

	// Address Operator
	x := 77
	fmt.Println("Memory address of x:",&x) // 0x1400009a020

	// Dereference Operator
	fmt.Println("The value of x:",*(&x))  // 77

	// Declaring a Pointer
	var i *int 
	var s *string
	fmt.Println(i,s)   // <nil> <nil>

	// Initializing a Pointer
	a := "hello"
	var b *string = &a
	fmt.Println(b)
	var c *string = &a
	fmt.Println(c)
	var d *string = &a
	fmt.Println(d)

	// Dereferencing a pointer
	var deref string = *(&a)
	fmt.Println(deref)
}