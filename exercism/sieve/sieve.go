package sieve

func Sieve(n int) []int {
	var output []int
	//create a boolean array
	prime := make([]bool, n+1)

	//initalise all values as true
	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If prime[p] unchanged then it is a prime
		if prime[p] {
			//Update all multiples of p
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	//return all prime numbers up to n
	for i := 2; i <= n; i++ {
		if prime[i] {
			output = append(output, i)
		}
	}
	return output
}
