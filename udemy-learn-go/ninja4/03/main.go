package main

// Hands-on exercise #3
// Using the code from the previous example, use SLICING to create the following new slices
// which are then printed:
// ● [42 43 44 45 46]
// ● [47 48 49 50 51]
// ● [44 45 46 47 48]
// ● [43 44 45 46 47]
import "fmt"

func main() {

	intSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x := intSlice[0:5]

	for _, number := range x {
		fmt.Printf("%d\n", number)
	}
	fmt.Println("")
	y := intSlice[5:10]

	for _, number := range y {
		fmt.Printf("%d\n", number)
	}
	fmt.Println("")
	z := intSlice[2:7]

	for _, number := range z {
		fmt.Printf("%d\n", number)
	}
	fmt.Println("")
	a := intSlice[1:6]

	for _, number := range a {
		fmt.Printf("%d\n", number)
	}
	fmt.Println("")
	//fmt.Printf("%T", x)
}
