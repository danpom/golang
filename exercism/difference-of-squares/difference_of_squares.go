package diffsquares

//SquareOfSum  returns the square of the sum of n numbers from 1 to n
func SquareOfSum(n int) int {
	result := 0
	result = SquareOfSumRec((n), result)
	result = result * result

	return result
}

//SquareOfSumRec is a helper function which uses recursion to calculate the square of a sum of numbers from 1 to n
func SquareOfSumRec(n, sum int) int {
	if n > 0 {
		sum += n
		return SquareOfSumRec((n - 1), sum)
	}
	return sum
}

//SumOfSquares returns the sum of the square of all numbers from 1 to n
func SumOfSquares(n int) int {
	result := 0
	result = SumOfSquaresRec(n, result)

	return result
}

//SumOfSquaresRec is a helper function which uses recursion to calculate the sum of the square of each number from 1 to n
func SumOfSquaresRec(n, sum int) int {
	if n > 0 {
		sum += n * n
		return SumOfSquaresRec((n - 1), sum)
	}
	return sum
}

//Difference returns the difference between the SquareOfSum and the SumOfSquares of a number n
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
