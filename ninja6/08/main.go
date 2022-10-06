package main

// Hands-on exercise #8
// ● Create a func which returns a func
// ● assign the returned func to a variable
// ● call the returned func

import (
	"fmt"
)

func main() {
	f := eight()
	f()
}

func eight() func() {

	return func() {
		fmt.Println("This is a func")
	}
}
