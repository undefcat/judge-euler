package main

import (
	"fmt"
	"math/big"
	"time"
)

func bruteForce() int {
	ans := 0
	for i := int64(99); i > 0; i-- {
		a := big.NewInt(i)
		b := big.NewInt(i)

		a = a.Mul(a, b)
		maxDigitSum := 0
		for j := 3; j <= 99; j++ {
			a = a.Mul(a, b)
			digitSum := sum(a.String())
			ans = max(ans, digitSum)
			maxDigitSum = max(maxDigitSum, digitSum)
		}

		// maxDigitSum은 대략적으로 감소하는 그래프일텐데
		// 값이 절반 이하로 떨어지면 더 이상 찾을 필요가 없을 거라고 추측할 수 있다.
		if ans/2 > maxDigitSum {
			break
		}
	}

	return ans
}

func sum(s string) int {
	ret := 0
	for _, v := range s {
		ret += int(v-'0')
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
	begin := time.Now()
	ans := bruteForce()
	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}
