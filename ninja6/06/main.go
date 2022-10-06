package main

// Hands-on exercise #6
// ● Build and use an anonymous func

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("I am an anon func")
	}()
}
