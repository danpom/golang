package main

import "fmt"

func main() {
	//var input = []int{1, 7, 3, 6, 5, 6}
	//var input = []int{1, 2, 3}
	var input = []int{2, 1, -1}
	answer := pivotIndex(input)
	fmt.Println(answer)
}
func pivotIndex(nums []int) int {
	size := len(nums)
	ans := -1
	for i := 0; i < size; i++ {
		left := 0
		right := 0

		for j := i; j > 0; j-- {
			left += nums[j-1]
		}
		for k := i; k < size-1; k++ {
			right += nums[k+1]
		}
		fmt.Printf("%v iteration. Left: %v Right: %v\n", i, left, right)
		//fmt.Printf("%v iteration. Left: %v\n", i, left)
		//fmt.Printf("%v iteration. Right: %v\n", i, right)

		if left == right {
			ans = i
			break
		}

	}
	return ans
}
