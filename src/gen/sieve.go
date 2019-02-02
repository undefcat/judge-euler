package gen

import "math"

func Sieve(n int) []bool {
	isPrime := make([]bool, n+1)
	isPrime[0] = true
	for i := 1; i < len(isPrime); i *= 2 {
		copy(isPrime[i:], isPrime[:i])
	}

	isPrime[0], isPrime[1] = false, false

	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if isPrime[i] {
			for j := i*i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}