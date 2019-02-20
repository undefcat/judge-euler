package main

import (
	"fmt"
	"gen"
	"time"
)

var (
	primes []int
)

func getPrimes(lower, upper int) []int {
	primes := make([]int, 0, upper>>4)
	sieve := gen.NewSieve(upper)

	for i, k := lower+1, 2; i < upper; i, k = i+k, 6-k {
		if sieve.IsPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func isPermutation(a, b int) bool {
	nums := make([]int, 10)

	for a != 0 {
		nums[a%10]++
		a /= 10
	}

	for b != 0 {
		nums[b%10]--

		b /= 10
	}

	for _, v := range nums {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	begin := time.Now()
	upper := 10000000
	primes = getPrimes(2000, 5000)
	minRatio := 10e15
	ans := 0
	for i := 1; i < len(primes); i++ {
		for j := 0; j < i; j++ {
			n := primes[i]*primes[j]
			if n > upper {
				break
			}

			pi := (primes[i]-1)*(primes[j]-1)

			if isPermutation(n, pi) {
				ratio := float64(n)/float64(pi)
				if minRatio > ratio {
					minRatio = ratio
					ans = n
				}
			}
		}
	}
	fmt.Println(ans)
	fmt.Println(time.Since(begin).String())
}