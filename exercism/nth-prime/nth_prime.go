package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {

	if n < 1 {
		return 0, errors.New("invalid input - primes cannot be negative")
	} else {
		//p := SieveOfEratosthenes(n * int(math.Sqrt(float64(n))))
		p := SieveOfEratosthenes(n*int(math.Round(math.Sqrt(float64(n)))) + 1)

		return p[n-1], nil
	}
}

//returns a slice of int containing all prime numbers from 2 to n
func SieveOfEratosthenes(n int) []int {
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
