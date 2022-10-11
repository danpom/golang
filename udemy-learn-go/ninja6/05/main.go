package main

// Hands-on exercise #5
// ● create a type SQUARE
// ● create a type CIRCLE
// ● attach a method to each that calculates AREA and returns it
// ○ circle area= π r 2
// ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle

import (
	"fmt"
	"math"
)

type square struct {
	l float64
}

type circle struct {
	r float64
}

func (s square) area() float64 {
	return s.l * s.l
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		l: 3.24,
	}

	c1 := circle{
		r: 4,
	}

	info(c1)
	info(s1)

}
