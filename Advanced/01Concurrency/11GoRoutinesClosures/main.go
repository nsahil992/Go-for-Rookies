package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{"Mixing vanilla ingredients", "Baking vanilla cake", "Decorating with white frosting",
		"Mixing chocolate ingredients", "Baking chocolate cake", "Decorating with chocolate frosting",
		"Mixing strawberry ingredients", "Baking strawberry cake", "Decorating with pink frosting"}

	for _, task := range tasks {
		currentTask := task
		go taskHelper("Helper:", currentTask)
	}
	time.Sleep(3 * time.Second)
}
func taskHelper(name string, task string) {
	fmt.Printf("%s is doing the task: %s\n", name, task)
	time.Sleep(2 * time.Second)
	fmt.Printf("%s finished task: %s\n", name, task)
}
