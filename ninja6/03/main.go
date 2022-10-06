package main

import "fmt"

// Hands-on exercise #3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
