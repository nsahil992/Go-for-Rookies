package greetings

import "fmt"

func SayHello() {
	fmt.Println("Hello from the greetings package")
}

func secret() {
	fmt.Println("This is a secret function which cannot be exported only used within the package")
}

const GlobalMessage = "Go lang is the best programming language"

var secretMessage = "But Go is little bit tricky"
