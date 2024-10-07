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
	a := 10
	var ptr_i *int = &a
	fmt.Println(ptr_i)

	z := "z"
	ptr_z := &z
	fmt.Println(ptr_z)

	// dereferencing a Pointer
	y := "Hello"
	fmt.Println(y)
	ptr_y := &y
	*ptr_y = "World"
	fmt.Println(y)

}