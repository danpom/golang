package main

import "fmt"

// Hands-on exercise #2
// ● create a func with the identifier foo that
// ○ takes in a variadic parameter of type int
// ○ pass in a value of type []int into your func (unfurl the []int)
// ○ returns the sum of all values of type int passed in
// ● create a func with the identifier bar that
// ○ takes in a parameter of type []int
// ○ returns the sum of all values of type int passed in

func main() {
	yi := []int{2, 6, 4, 7, 4}
	foo1 := foo(yi...)
	fmt.Println(foo1)

	bar1 := bar(yi)
	fmt.Println(bar1)

}

func foo(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
