package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Create("create.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Successfully created file")

	content := "Welcome to the world of Go"
	_, err = io.WriteString(file, content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully written to file")

	// Buffer

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer[:n]))
	}
}
