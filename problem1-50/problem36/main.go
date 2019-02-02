package main

import (
	"fmt"
	"time"
)

func isPalindromic(n, digit int) bool {
	return isDecimalPalindromic(n, digit) && isBinaryPalindromic(n)
}

func isDecimalPalindromic(n, digit int) bool {
	arr := make([]int, digit)
	m := n
	for i := 0; i < digit; i++ {
		arr[i] = m%10
		m /= 10
	}

	mid, last := digit/2, digit-1
	for i := 0; i < mid; i++ {
		if arr[i] != arr[last-i] {
			return false
		}
	}
	return true
}

func isBinaryPalindromic(n int) bool {
	arr := make([]int, 0, 20)
	for i := 0; n != 0; i++ {
		arr = append(arr, n%2)
		n /= 2
	}

	mid, last := len(arr)/2, len(arr)-1
	for i := 0; i < mid; i++ {
		if arr[i] != arr[last-i] {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	ans := 0
	digit := 1
	exp10 := 10
	for n := 1; n < 1000000; n+=2 {
		if n >= exp10 {
			digit++
			exp10 *= 10
		}

		if isPalindromic(n, digit) {
			ans += n
		}
	}

	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}
