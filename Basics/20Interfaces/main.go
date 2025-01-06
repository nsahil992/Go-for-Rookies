package main

import "fmt"

func main() {
	dog := Dog{}
	cat := Cat{}

	introduce(dog)
	introduce(cat)
}

type Speaker interface {
	Speak() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) Speak() string {
	return "Meow!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}
