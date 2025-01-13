package main

import (
	g1 "golearn/Advanced/02Modules/04PackageNameCollision/greetings1" // Alias greetings1 as g1
	g2 "golearn/Advanced/02Modules/04PackageNameCollision/greetings2" // Alias greetings1 as g2
)

func main() {
	g1.Greet1()
	g2.Greet2()
}
