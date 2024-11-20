package main

import (
	"fmt"
)

func main() {
	person := Person{Name: "Sahil", Age: 21}
	person.greet()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) greet() {
	fmt.Println("Hello, My name is ", p.Name, "and I am ", p.Age, "years old.")
}
