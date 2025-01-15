package main

import "fmt"

type Car struct {
	name  string
	model int
	color string
}

func (car1 Car) enginestart() {
	fmt.Printf("The %v %v colour %v has just started.\n", car1.model, car1.color, car1.name)
}

func (car2 Car) highestSpeed() {
	fmt.Printf("The %v %v is the fastest car in the world\n", car2.model, car2.name)
}

func main() {
	var fastestCar Car
	fastestCar.name = "Bugatti Chiron Supersport"
	fastestCar.model = 2016
	fastestCar.color = "Exposed Blue Carbon Fiber With Silve"

	fastestCar.highestSpeed()
	fastestCar.enginestart()
}
