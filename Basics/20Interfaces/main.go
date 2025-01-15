package main

import "fmt"

// Defining an interface `Vehicle`
type Vehicle interface {
	enginestart()  // Method to start the engine
	highestSpeed() // Method to declare the highest speed
}

// Defining the Car struct
type Car struct {
	name  string
	model int
	color string
}

// `Car` implements the `enginestart` method
func (car Car) enginestart() {
	fmt.Printf("The %v %v colour %v has just started.\n", car.model, car.color, car.name)
}

// `Car` implements the `highestSpeed` method
func (car Car) highestSpeed() {
	fmt.Printf("The %v %v is the fastest car in the world\n", car.model, car.name)
}

// Main function
func main() {
	// Creating a Car instance
	var fastestCar Car
	fastestCar.name = "Bugatti Chiron Supersport"
	fastestCar.model = 2016
	fastestCar.color = "Exposed Blue Carbon Fiber With Silver"

	// Declare a variable of type `Vehicle` and assign `fastestCar`
	var myVehicle Vehicle = fastestCar

	// Call methods using the interface
	myVehicle.highestSpeed()
	myVehicle.enginestart()
}
