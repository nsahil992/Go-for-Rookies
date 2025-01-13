package main

import (
	"fmt"
	"golearn/Advanced/02Modules/03ImportsAndExports/greetings"
)

func main() {
	greetings.SayHello()
	fmt.Println(greetings.GlobalMessage)

	// Works because it's exported using the first letter as capital

	//greetings.secretMessage() // Error: secretMessage is not exported
}
