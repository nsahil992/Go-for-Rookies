package main

import (
	"fmt"
)

type Car struct {
	name string
	color string
	model int
}

func main() {
	var car1 Car
	car1.name = "Lexus ES 300H"
	car1.color = "White"
	car1.model = 2024

	fmt.Println("The name of the car is",car1.name)
	fmt.Printf("The color of %v is %v\n",car1.name,car1.color)
	fmt.Printf("The model of %v is %v\n",car1.name,car1.model)

	// Second method

	car2 := Car {
		name: "Maserati",
		color: "Navy",
		model: 2023,
	}

	fmt.Printf("The %v color %v %v model has been started!\n",car2.color,car2.name,car2.model)

	// Third method(new keyword)

	var car3 = new(Car) 
	car3.name = "Range Rover Sports"
	car3.color = "Black"
	car3.model = 2024

	fmt.Println(car3)
}