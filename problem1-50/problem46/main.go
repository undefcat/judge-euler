package main

import (
	"fmt"
	"math"
	"time"
)

var primes = make([]int, 0, 10001)

func conjection(n int) bool {
	for i := 0; i < len(primes); i++ {
		if primes[i] > n {
			return false
		}

		for j := 1; j < n; j++ {
			sum := primes[i] + 2*j*j
			if sum > n {
				break
			}

			if sum == n {
				return true
			}
		}
	}
	panic("need more prime")
}

func main() {
	start := time.Now()
	n := 3
	for {
		sqrt := int(math.Sqrt(float64(n)))
		isPrime := true
		for i := 3; i <= sqrt; i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, n)
			if len(primes) == cap(primes) {
				break
			}
		}
		n += 2
	}

	ans := 0
out:
	for a := 0; a < len(primes); a++ {
		for b := a; b >= 0; b-- {
			comp := primes[a]*primes[b]

			if !conjection(comp) {
				ans = comp
				break out
			}
		}
	}

	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}