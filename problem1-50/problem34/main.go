package main

import (
	"fmt"
	"time"
)

var factorials []int
func calc(n int) int {
	sum := 0
	for n != 0 {
		sum += factorials[n%10]
		n /= 10
	}
	return sum
}

func main() {
	start := time.Now()
	factorials = make([]int, 10)
	factorials[0] = 1

	for i := 1; i < len(factorials); i++ {
		factorials[i] = factorials[i-1]*i
	}

	ans := 0
	for i := 3; i <= 2540160; i++ {
		if i == calc(i) {
			ans += i
		}
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}