package main

import (
	"fmt"
	"time"
)

var (
	dn = make([]int, 1)
	exp10 = make([]int, 10)
)

func getDigit(n int) int {
	for i := 1; i < len(dn); i++ {
		// n은 i자릿수의 수이다.
		if (dn[i-1] < n) && (n <= dn[i]) {
			// n자릿수에서 몇번째 digit인지 (zero-based)
			idx := n-dn[i-1]-1

			// 수 n은 어떤 숫자에 포함되어 있는지
			// n자리의 0번째 수부터 구하면 된다.
			num := exp10[i-1]+(idx/i)

			// 그 수에서 몇번째 digit인지 (zero-based)
			nth := idx%i

			// 그 digit의 숫자를 알아내기
			for j := 1; j < i-nth; j++ {
				num /= 10
			}
			return num%10
		}
	}

	panic("overflow")
}

func main() {
	start := time.Now()
	exp10[0] = 1
	exp10[1] = 10
	for n := 1; n < 10; n++ {
		expression := 9*n*exp10[n-1]
		dn = append(dn, dn[n-1]+expression)
		exp10[n] = exp10[n-1]*10
	}

	sum := 1
	for i := 0; i < 7; i++ {
		sum *= getDigit(exp10[i])
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(sum)
}