package main

import (
	"fmt"
	"math"
	"time"
)

func pentagonal(n int) int {
	return n*(3*n-1)/2
}

func isPentagonal(n int) bool {
	sqrt := math.Sqrt(float64(24*n)+1)
	result := (sqrt+1)/6
	return result == float64(int(result))
}

func main() {
	start := time.Now()
	ans := 0
	k := 4
out:
	for {
		pk := pentagonal(k)
		for j := k-1; j >= 2; j-- {
			pj := pentagonal(j)
			if isPentagonal(pk+pj) && isPentagonal(pk-pj) {
				ans = pk-pj
				break out
			}
		}
		k++
	}

	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}