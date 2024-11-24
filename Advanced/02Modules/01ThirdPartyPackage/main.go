package main

import (
	"fmt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)
	fmt.Println(nimbus)

	// Encrypt and decrypt

}

func nimbus(str string) string {
	encryptedStr := ""
	for c, _ := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
