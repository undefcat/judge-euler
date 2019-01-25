package main

import (
	"fmt"
	"time"
)

func mySolution() int {
	maximum := 0
	ans := 0
	for p := 4; p <= 1000; p++ {
		count := 0
		for c := 1; c < p; c++ {
			cc := c*c
			for b := c; b < p-c; b++ {
				a := p-c-b
				if a*a == b*b + cc {
					count++
				}
			}
		}

		if count > maximum {
			maximum = count
			ans = p
		}
	}
	return ans
}

func sol1() int {
	ans := 0
	maxCount := 0
	for p := 4; p <= 1000; p += 2 {
		count := 0
		for a := 2; a < (p/3); a++ {
			if p*(p-2*a) % (2*(p-a)) == 0 {
				count++
			}
		}

		if count > maxCount {
			maxCount = count
			ans = p
		}
	}
	return ans
}

func main() {
	start := time.Now()

	ans := sol1()

	fmt.Println(time.Since(start).String())	// 1ms
	fmt.Println(ans)

	start = time.Now()
	ans = mySolution()

	fmt.Println(time.Since(start).String())	// 100ms
	fmt.Println(ans)
}
