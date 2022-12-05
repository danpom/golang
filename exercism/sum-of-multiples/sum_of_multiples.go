package summultiples

import "sort"

func SumMultiples(limit int, divisors ...int) int {
	//var multiples []int
	sum := 0
	if len(divisors) < 1 {
		return 0
	}
	if len(divisors) == 1 && divisors[0] == 0 {
		return 0
	}

	sort.Slice(divisors, func(i, j int) bool {
		return divisors[i] < divisors[j]
	})

	//index of first non zero divisor
	ind := 0
	for i, v := range divisors {
		if v > 0 {
			ind = i
			break
		}
	}

	divisors = divisors[ind:]

	for i := divisors[0]; i < limit; i++ {
		isMultiple := false
		for _, v := range divisors {
			if i%v == 0 {
				isMultiple = true
				break
			}
		}
		if isMultiple {
			//multiples = append(multiples, i)
			sum += i
		}
	}

	return sum
}
