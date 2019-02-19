package main

import (
	"fmt"
	"gen"
	"time"
)

func main() {
	begin := time.Now()
	primes := make([]int, 1, 100)
	primes[0] = 2
	sieve := gen.NewSieve(100)

	for n := 3; n <= 100; n += 2 {
		if sieve.IsPrime(n) {
			primes = append(primes, n)
		}
	}

	sum := 1
	for _, prime := range primes {
		val := sum*prime
		if val > 1000000 {
			break
		}
		sum = val
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(sum)
}
