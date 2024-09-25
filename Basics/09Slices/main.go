package main
import(
	"fmt"
)

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Print("Slice is: ")
	slice_1 := arr[0:8]
	fmt.Println(slice_1)  // [10 20 30 40 50 60 70 80]

	// Changing value of slice

	slice_1[0] = 11
	fmt.Println("After modification of Slice: ")
	fmt.Println(slice_1)

	cars := []string{"Maserati", "Mini Cooper", "Defender", "Range Rover", "Lexus"}
	car_slice := cars[0:5]
	fmt.Println(car_slice)

	for index, car := range cars {
		fmt.Println(index, "=>", car)
	} 

	// Appending an element

	add_cars := append(cars, "BMW 340i, Mercedes Maybach")
	fmt.Println(add_cars)

	fmt.Println(cap(cars))
}