package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string, 3)

	go mixIngredients(done)
	go bakeCake(done)
	go decorateCake(done)

	time.Sleep(8 * time.Second)

	fmt.Println("Cake is ready 🎂🎂")
}

func mixIngredients(done chan string) {
	fmt.Println("Mixing Ingredients...")
	time.Sleep(2 * time.Second)
	done <- "Ingredients mixed"
	fmt.Println("Finished mixing ingredients.")
}

func bakeCake(done chan string) {
	msg := <-done
	fmt.Println("Received:", msg)
	fmt.Println("Baking the cake...")
	time.Sleep(3 * time.Second)
	done <- "Cake is baked!"
	fmt.Println("Finished baking cake.")
}

func decorateCake(done chan string) {
	msg := <-done
	fmt.Println("Received:", msg)
	fmt.Println("Decorating cake...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished decorating cake.")
}
