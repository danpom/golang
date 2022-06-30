package main

import "fmt"

//Print every number from 1 to 10,000

func main() {
	var input = []int{1, 2, 3, 17, 50}
	answer := runningSum(input)
	fmt.Printf("[")
	for _, str := range answer {
		fmt.Printf("%d ", str)
	}
	fmt.Printf("]")
}

func runningSum(nums []int) []int {
	//declare an output[] with the same length as nums
	//using length of nums more than once so store in a variable
	size := len(nums)

	output := make([]int, size)

	//declare an int variable to store the runningSum. initialise as 0
	var iterativeSum int = 0

	//iterate through nums[]; iterativeSum +=current value
	for i := 0; i < size; i++ {
		iterativeSum += nums[i]
		output[i] = iterativeSum
	}
	return output
}
