package main

import (
	"fmt"
	"gen"
	"time"
)

var (
	digit5 = [][]int{
		{0, 0, 0, 1, 1},
		{0, 0, 1, 0, 1},
		{0, 1, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}

	digit6 = [][]int{
		{0, 0, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 1},
		{0, 0, 1, 1, 0, 1},
		{0, 1, 0, 0, 1, 1},
		{0, 1, 0, 1, 0, 1},
		{0, 1, 1, 0, 0, 1},
		{1, 0, 0, 0, 1, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 1, 0, 0, 0, 1},
	}
)

var isPrime = gen.Sieve(1000000)

func main() {
	begin := time.Now()
	ans := (1<<63)-1
	// 각 패턴에 10의자리, 100의자리를 쪼개서 넣는다.
	// 이 때 1의 자리는 1, 3, 7, 9여야만 한다.
	for n := 11; n < 1000; n += 2 {
		var pattern [][]int

		if n%5 == 0 {
			continue
		}

		if n < 100 {
			pattern = digit5
		} else {
			pattern = digit6
		}

		for i := 0; i < len(pattern); i++ {
			// 반복되는 수는 0, 1, 2 셋 중 하나로 시작해야 한다. 그래야 8개의 prime number가
			// 나올 수 있으므로!
			for k := 0; k <= 2; k++ {
				// 만약 첫자리에 수를 채우는데
				// 그 수가 0이면 패스
				if pattern[i][0] == 0 && k == 0 {
					continue
				}

				// n을 쪼개서 패턴에 수를 채움
				filledPattern := fillPattern(pattern[i], n)

				// 가장 작은 후보수를 만든다.
				candidate := generateNumber(filledPattern, k)

				// 만약 소수면 걔의 family size를 알아내본다.
				if isPrime[candidate] {
					fz := getFamilySize(filledPattern, k+1)
					if fz == 8 {
						if ans > candidate {
							ans = candidate
						}
						break
					}
				}
			}
		}
	}
	fmt.Println(time.Since(begin).String())
	fmt.Print(ans)
}

func fillPattern(p []int, n int) []int {
	ret := make([]int, len(p))

	for i := len(p)-1; i >= 0; i-- {
		if p[i] == 1 {
			ret[i] = n%10
			n /= 10
		} else {
			// -1로 마크한 뒤, 여기에 반복되는 수를 넣자.
			ret[i] = -1
		}
	}

	return ret
}

func generateNumber(p []int, n int) int {
	ret := 0
	for i := 0; i < len(p); i++ {
		ret *= 10
		if p[i] == -1 {
			ret += n
		} else {
			ret += p[i]
		}
	}

	return ret
}

func getFamilySize(p []int, repeat int) int {
	ret := 1
	for i := repeat; i < 10; i++ {
		num := generateNumber(p, i)
		if isPrime[num] {
			ret++
		} else if 9-i+ret < 8 {
			return -1
		}
	}

	return ret
}