package main

import (
	"fmt"
	"time"
)

func gcd(a, b int) int {
	if a < b {
		b, a = a, b
	}

	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	start := time.Now()
	numers, denos := 1, 1
	// 분모의 10의자리
	for d2 := 1; d2 < 10; d2++ {
		// 분모의 1의자리.
		// 30/50 같은 단순한 경우는 제외하기로 했으므로, 0인 경우는 제외한다.
		for d1 := 1; d1 < 10; d1++ {
			deno := d2*10 + d1

			// 분자의 10의자리
			for n2 := 1; n2 <= d2; n2++ {
				// 분자의 1의자리.
				for n1 := 1; n1 < 10; n1++ {
					// 각 1의자리 <-> 10의자리가 같은 값이 없으면 pass
					// 이는 문제의 조건과 맞지 않음
					if n1 != d2 && n2 != d1 {
						continue
					}

					numer := n2*10 + n1

					// 1보다 작은 값에 대해서만 계산한다.
					if numer/deno < 1 {
						if float64(numer)/float64(deno) == float64(n2)/float64(d1) {
							numers *= numer
							denos *= deno
						}
					}
				}
			}
		}
	}
	// 문제에서 말한 4가지 경우는
	// 16/64
	// 26/65
	// 19/95
	// 49/98
	// 위의 4가지 인데 이를 모두 곱한 뒤 약분하여 분모를 출력하면 된다.
	// 각각의 분자/분모를 곱했으므로 이 둘의 최대공약수를 유클리드 알고리즘으로 구한 뒤
	// 분모/최대공약수를 나누면 약분이 된다.
	ans := denos/gcd(denos, numers)
	end := time.Since(start).String()
	fmt.Println(ans)
	fmt.Println(end)
}

// 거의 뽀록으로 맞춘 수준... 제대로 다시 해봐야 함.