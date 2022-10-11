package main

import "fmt"

// Hands-on exercise #6
// Create a slice to store the names of all of the states in the United States of America. Use make
// and append to do this. Goal: do not have the array that underlies the slice created more than
// once. What is the length of your slice? What is the capacity? Print out all of the values, along
// with their index position, without using the range clause. Here is a list of the 50 states:
// ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, `
// Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, `
// Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, `
// Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`,
// ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`,
// ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, `
// Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
// solution:
// ● my first take on this project, which was an INCORRECT solution
// ○ https://go.dev/play/p/vE9BCWMA2ET
// ● instructive proof from a student which showed me the error of my ways :)
// ○ https://go.dev/play/p/szSlBwJG2S-
// ● HERE IS THE CORRECT SOLUTION
// ○ https://go.dev/play/p/ZcyFCyo7_XH
// It is good to think about the LEN and CAP. It is also good to think about the underlying
// array. You want to think about this so that you're thinking of memory use. Also, it is good
// to think about what got stored in what position. And it is good to remember that slices
// are dynamic - they can grow in size.
// FROM THE SPEC:
// The length of a slice s can be discovered by the built-in function len; unlike with arrays it
// may change during execution. The elements can be addressed by integer indices 0
// Todd McLeod - Learn To Code Go on Udemy - Part 1 - Page 44through len(s)-1. The slice index of a given element may be less than the index of the
// same element in the underlying array.
// A slice, once initialized, is always associated with an underlying array that holds its
// elements. A slice therefore shares storage with its array and with other slices of the same
// array; by contrast, distinct arrays always represent distinct storage.
// The array underlying a slice may extend past the end of the slice. The capacity is a
// measure of that extent: it is the sum of the length of the slice and the length of the array
// beyond the slice; a slice of length up to that capacity can be created by slicing a new one
// from the original slice. The capacity of a slice a can be discovered using the built-in
// function cap(a).
// https://go.dev/ref/spec#Slice_types
// video: 076

func main() {
	//var my_slice_1 = []string{"Geeks", "for", "Geeks"}
	states := make([]string, 50)
	states = append(states, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`,
		`Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`,
		`Missouri`, `Montana`, `Nebraska`, `Nevada`, ` New Hampshire`, ` New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`,
		`Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`,
		`Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for state := 0; state < len(states); state++ {
		fmt.Printf("%d : %s\n", state, states[state])
	}

}
