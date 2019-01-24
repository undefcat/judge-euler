package main

import (
	"fmt"
	"time"
)

func isPandigital(n int) bool {
	var digits, count, tmp int

	for n > 0 {
		tmp = digits
		digits |= 1<<uint(n-((n/10)*10)-1)
		if tmp == digits {
			return false
		}
		n /= 10
		count++
	}
	return digits == 1<<uint(count)-1
}

func getDigit(n int) int {
	ret := 0
	for n != 0 {
		n /= 10
		ret++
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	start := time.Now()
	ans := 0
	exp10 := make([]int, 10)
	exp10[0] = 1
	for i := 1; i < 10; i++ {
		exp10[i] = exp10[i-1]*10
	}

	for n := 1; n < 10000; n++ {
		digit := getDigit(n)
		sum := n
		for i := 2; i < 10; i++ {
			mul := n*i
			currentDigit := getDigit(mul)
			sum *= exp10[currentDigit]
			sum += mul
			digit += currentDigit
			if digit >= 9 {
				break
			}
		}

		if digit == 9 && isPandigital(sum) {
			ans = max(ans, sum)
		}
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}