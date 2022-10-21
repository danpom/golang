package main

// Hands-on exercise #6
// Create a program that prints out your OS and ARCH. Use the following commands to run it
// ● go run
// ● go build
// ● go install

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Os: ", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
}
