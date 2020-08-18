package primenumbers

// Sieve returns an array of prime numbers.
// It uses the sieve of Eratosthenes algorithm.
func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	// Create a bool array for each number from 2 to limit (0 and 1 can be skipped)
	size := limit - 2 + 1
	primeBools := make([]bool, size)

	primeBools[0] = true
	for i := 1; i < size; i *= 2 {
		copy(primeBools[i:], primeBools[:i])
	}

	result := make([]int, 0)

	for i := 0; i < size; i++ {
		if !primeBools[i] {
			continue
		}

		val := i + 2
		result = append(result, val)
		for j := val*val - 2; j < size; j += val {
			primeBools[j] = false
		}

	}

	return result
}
