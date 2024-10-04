package main

import "fmt"

// Define an enum using iota
type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    today := Tuesday
    fmt.Println(today)        // Output: 2 (Tuesday is the 3rd item, so it gets value 2)

    // You can also use a switch to work with these enums
    switch today {
    case Sunday:
        fmt.Println("It's Sunday!")
    case Monday:
        fmt.Println("It's Monday!")
    case Tuesday:
        fmt.Println("It's Tuesday!")
    default:
        fmt.Println("It's another day!")
    }
}
