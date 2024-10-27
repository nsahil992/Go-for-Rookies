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

	for i := 0; i < 3; i++ {
		select {
		case msg := <-done:
			fmt.Println("Received:", msg)
		case <-time.After(4 * time.Second):
			fmt.Println("Task took too long and timed out!")
			<-done
		}
	}
	fmt.Println("Cake is ready for the party!")
}

func mixIngredients(done chan<- string) {
	fmt.Println("Alice is mixing ingredients...")
	time.Sleep(2 * time.Second)
	done <- "Ingredients are mixed"
	fmt.Println("Alice: Finished mixing ingredients.")
}

func bakeCake(done chan<- string) {
	fmt.Println("Bob is baking cake...")
	time.Sleep(5 * time.Second)
	done <- "Cake is baked."
	fmt.Println("Bob: Finished baking cake.")
}

func decorateCake(done chan<- string) {
	fmt.Println("Carol is decorating cake...")
	time.Sleep(2 * time.Second)
	done <- "Cake is decorated."
	fmt.Println("Carol: Finished decorating cake.")
}
