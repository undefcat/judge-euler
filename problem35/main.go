package main

import (
	"fmt"
	"time"
)

var primes []int

func isPrime(n int) bool {
	if primes[n] == -1 {
		return false
	} else if primes[n] == 1 {
		return true
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			primes[n] = -1
			return false
		}
	}

	primes[n] = 1
	return true
}

func isCircularPrime(n, digits, exp10 int) bool {
	for i := 0; i < digits; i++ {
		end := n%10
		n /= 10
		n += end*exp10
		if !isPrime(n) {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()
	primes = make([]int, 1000001)
	primes[2] = 1

	count := 0
	digits := 1
	exp10 := 1
	exp100 := 10
	for i := 2; i < 1000000; i++ {
		if i >= exp100 {
			exp10 *= 10
			exp100 *= 10
			digits++
		}

		if isCircularPrime(i, digits, exp10) {
			count++
		}
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(count)
}
