package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])
}
func main() {
	fmt.Println("Enter your password: ")
	var password string
	_, err := fmt.Scanln(&password)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Password encrypted to: ", encodeWithMD5(password))
}
