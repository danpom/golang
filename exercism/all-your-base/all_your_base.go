package allyourbase

import (
	"errors"
	"math"
)

//ConvertToBase converts inputDigits (a slice of int representing a number of base "inputBase". Each element is a digit of the number)
//to slice of int representing the same number in base "outputBase"
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	/*
		steps:
		convert to base 10
		convert from base 10 to new base
	*/

	switch {
	case inputBase < 2:
		return []int{0}, errors.New("input base must be >= 2")
	case outputBase < 2:
		return []int{0}, errors.New("output base must be >= 2")
	case !validDigit(inputDigits, inputBase):
		return []int{0}, errors.New("all digits must satisfy 0 <= d < input base")
	default:
		return From10ToBaseOB(ToBase10(inputBase, inputDigits), outputBase, []int{}), nil

	}
}

//ToBase10 converts id (a slice of int representing a number of base "ib". Each element is a digit of the number) to base 10
func ToBase10(ib int, id []int) int {
	var res int

	c := float64(len(id) - 1)
	for _, v := range id {
		res += (v * int(math.Pow(float64(ib), c)))
		c--
	}
	//return IntToSlice(res, []int{})
	return res
}

//From10ToBaseOB is used to convert a base 10 int n to an array of digits representing the same number converted to base "ob".
// This is a recursive function so it should be called with output as "[]int{}"
func From10ToBaseOB(n, ob int, output []int) []int {
	if n != 0 {
		output = append([]int{n % ob}, output...)
		return From10ToBaseOB(n/ob, ob, output)
	}

	if len(output) == 0 {
		return []int{0}
	} else {
		return output
	}
}

// validDigit checks if all digits in a slice of int are positive and less than or equal int n
func validDigit(s []int, n int) bool {
	for _, v := range s {
		if v < 0 || v >= n {
			return false
		}
	}

	return true
}

//Note: This function wasn't needed but leaving here for posterity
//IntToSlice converts an int n into a slice of int where each element is a digit of n
// func IntToSlice(n int, sequence []int) []int {
// 	if n != 0 {
// 		i := n % 10
// 		sequence = append([]int{i}, sequence...)
// 		return IntToSlice(n/10, sequence)
// 	}
// 	return sequence
// }
