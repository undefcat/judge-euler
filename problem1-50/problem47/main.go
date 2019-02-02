package main

import (
	"fmt"
	"time"
)

func primeFactorization(n int) int {
	ret := 1
	i := 2
	for {
		if i*i > n {
			break
		}
		if n%i == 0 {
			n /= i
			for n%i == 0 {
				n /= i
			}
			ret++
		} else {
			i++
		}
	}

	return ret
}

func main() {
	begin := time.Now()
	counter := 0
	start := 0
	n := 210
	for counter != 4 {
		if primeFactorization(n) == 4 {
			if counter == 0 {
				start = n
			}
			counter++

		} else {
			counter = 0
			start = 0
		}
		n++
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(start)
}
