package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string)
	go mixIngredients(done)
	go bakeCake(done)

	time.Sleep(5 * time.Second)
	fmt.Println("Cake is ready")
}

func mixIngredients(done chan string) {
	fmt.Println("Alex: Mixing ingredients...")
	time.Sleep(2 * time.Second)
	done <- "Ingredient are ready!"
	fmt.Println("Alex: Ingredients are mixed.")
}

func bakeCake(done chan string) {
	msg := <-done
	fmt.Println("Sam received: ", msg)
	fmt.Println("Sam baking cake...")
	time.Sleep(3 * time.Second)
	fmt.Println("Sam: Cake is baked!")
}
