package main

import "fmt"

func main() {
	check := validateUsername("")
	handleError(check)

	check = validateUsername("John Doe")
	handleError(check)
}

func validateUsername(username string) error {
	if username == "" {
		return fmt.Errorf("username is empty")
	} else {
		return nil
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Username is valid")
	}
}
