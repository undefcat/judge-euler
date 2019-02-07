package main

import (
	"fmt"
	"is"
	"math/big"
	"strconv"
	"time"
)

var (
	cached = make([]int, 10001)
)

func isLychrelNumber(n int) bool {
	if cached[n] != 0 {
		if cached[n] == 1 {
			return true
		} else {
			return false
		}
	}

	num := strconv.Itoa(n)
	reversed, _ := strconv.Atoi(reverse(num))

	var a, b *big.Int

	a = new(big.Int)
	a.SetString(num, 10)

	for i := 0; i < 50; i++ {
		b = new(big.Int)
		b.SetString(reverse(a.String()), 10)

		a = a.Add(a, b)
		if is.PalindromeString(a.String()) {
			cached[reversed] = -1
			return false
		}
	}

	cached[reversed] = 1
	return true
}

func reverse(s string) string {
	l := len(s)
	last := l-1
	tmp := make([]byte, l)

	for i := 0; i < l; i++ {
		tmp[i] = s[last-i]
	}

	return string(tmp)
}

func main() {
	begin := time.Now()
	ans := 0

	for n := 0; n < 10000; n++ {
		if isLychrelNumber(n) {
			ans++
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}