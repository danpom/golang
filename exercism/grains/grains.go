package grains

import (
	"errors"
)

//Square calculates how many grains were on a given square number.
//Note: there is a formula to calculate this "(2^n)-1". That I wasn't initially aware of.
//Should return to this for a second iteration using the fomula to be more efficient
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("a chess board has only 64 squares. Please enter a number between 1 and 64")
	}
	return SquareRec(number, 1), nil
}

//SquareRec is a helper function that uses recursion to calculate
//how many grains were on a given square.
func SquareRec(number int, sum uint64) uint64 {
	if number > 1 {
		return SquareRec(number-1, sum*2)
	}
	return sum
}

//Total calculates how many grains are on a the chessboard at the end.
//This is a static value so perhaps a hardcoded answer would be sufficient but where's the challenge in that?
func Total() uint64 {
	var sum uint64
	for i := 64; i > 0; i-- {
		squareVal, _ := Square(i)
		sum += squareVal
	}
	return sum
}
