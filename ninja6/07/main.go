package main

// Hands-on exercise #7
// ● Assign a func to a variable, then call that func

import (
	"fmt"
)

var f func()

func main() {
	f := func() {
		fmt.Println("This is a func")
	}

	f()
	fmt.Printf("%T\n", f)
}
