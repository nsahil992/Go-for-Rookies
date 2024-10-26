package main

import (
	"fmt"
	"time"
)

func mixIngredients(done chan<- string) {
	fmt.Println("Mixing the ingredients...")
	time.Sleep(2 * time.Second)
	done <- "Ingredients are mixed"
	fmt.Println("Finished mixing ingredients")
}

func bakeCake(done chan string) {
	msg := <-done
	fmt.Println("Received:", msg)
	fmt.Println("Baking the cake...")
	time.Sleep(3 * time.Second)
	done <- "Cake is baked"
	fmt.Println("Finished baking the cake.")
}

func decorateCake(done chan string) {
	msg := <-done
	fmt.Println("Received:", msg)
	fmt.Println("Decorating the cake...")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished decorating the cake.")
}

func main() {
	done := make(chan string, 3)

	go mixIngredients(done)
	go bakeCake(done)
	go decorateCake(done)

	time.Sleep(8 * time.Second)

	close(done)

	fmt.Println("Cake is ready!")
}
