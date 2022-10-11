package main

// Hands-on exercise #1
// ○ create a func with the identifier foo that returns an int
// ○ create a func with the identifier bar that returns an int and a string
// ○ call both funcs
// ○ print out their results

import (
	"fmt"
)

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 27
}

func bar() (int, string) {
	return 27, "twenty seven"
}
