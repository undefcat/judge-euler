package main

import (
	"fmt"
	"gen"
	"is"
	"time"
)

var sieve gen.Sieve

func isPrime(n int) bool {
	if n < 10000000 {
		return sieve.IsPrime(n)
	}

	return is.Prime(n)
}

func main() {
	begin := time.Now()

	sieve = gen.NewSieve(10000000)
	primes := 0
	n := 2
out:
	for {
		k := (2*n)-1
		kk := k*k

		for i := 3; i >= 0; i-- {
			edgeNum := kk-(k-1)*i
			if isPrime(edgeNum) {
				primes++
			}
		}

		ratio := float64(primes)/float64(2*k-1)
		if ratio < 0.10 {
			break out
		}

		n++
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(2*n-1)
}
