package primenumbers

// FindPrimeNumbers returns an array of all prime numbers up to the given number.
func FindPrimeNumbers(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	primes := []int{2}

outer:
	for i := 3; i <= limit; i++ {
		for _, num := range primes {
			if float64(i)/float64(num) == float64(i/num) {
				continue outer
			}
		}
		primes = append(primes, i)
	}

	return primes
}
