package main

import (
	"fmt"
	"gen"
	"time"
)

const Max = 1000000

func main() {
	start := time.Now()
	sieve := gen.Sieve(Max)
	primes := make([]int, 0, 10000)
	for n, isPrime := range sieve {
		if isPrime {
			primes = append(primes, n)
		}
	}

	primeSum := make([]int, len(primes))
	primeSum[0] = primes[0]

	for i := 1; i < len(primes); i++ {
		primeSum[i] = primeSum[i-1] + primes[i]
	}

	longest := 0
	ans := 0
out:
	for i := 1; i < len(primes); i++ {
		for j := i-longest-1; j >= 0; j-- {
			ret := primeSum[i] - primeSum[j]
			if ret > Max {
				continue out
			}

			if sieve[ret] && i-j+1 >= longest {
				ans = ret
				longest = i-j+1
			}
		}
	}

	fmt.Println(time.Since(start).String())
	fmt.Printf("term: %d, ans: %d\n", longest, ans)
}
