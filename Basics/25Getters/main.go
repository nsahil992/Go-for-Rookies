package main

import "fmt"

type User struct {
    name string
}

func (u *User) GetName() string {
    return u.name
}

func main() {
    user := User{name: "Sahil"}
    fmt.Println("User Name:", user.GetName())
}