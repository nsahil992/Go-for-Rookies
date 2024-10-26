package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go task("Boiling water", &wg)
	go task("Making tea", &wg)
	go task("Serving tea", &wg)

	wg.Wait()
	fmt.Println("Tea is ready!")
}

func task(name string, wg *sync.WaitGroup) {
	fmt.Println("Started", name)
	time.Sleep(time.Second)
	fmt.Println("Finished", name)

	wg.Done()
}
