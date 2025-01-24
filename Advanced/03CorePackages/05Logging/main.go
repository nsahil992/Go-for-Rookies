package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

/*
Logrus supports 7 of log levels
Trace
Info
Warn
Panic
Debug
Fatal
Error
*/
func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	logrus.SetOutput(file)
	logrus.Println("Hello World")

	// Ex

	logrus.Warn("Hello World")
	logrus.Error("Hello World")
	logrus.Panic("Hello World")
}
