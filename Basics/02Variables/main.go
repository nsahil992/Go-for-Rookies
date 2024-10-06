package main

import "fmt"

func main() {
	var name string = "Sahil"
	fmt.Println("My name is", name)

	var age int = 21
	fmt.Printf("My age is %d\n",age)

	const lang = "Go"
	fmt.Printf("Currently, I am learning %slang\n",lang)
	
	const pi = 3.14
	fmt.Printf("The value of pi is %.2f\n",pi)

	// shorthand

	mascot := "Gopher"
	fmt.Printf("The mascot of Go lang is %s\n",mascot)

	const value = 349.438
	fmt.Printf("The variable value is a type of %T\n",value)
}
