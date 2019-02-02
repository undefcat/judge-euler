package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func checkPermutation(a, b int) bool {
	num := make([]int, 10)
	for a != 0 {
		num[a%10]++
		a /= 10
	}

	for b != 0 {
		if num[b%10] == 0 {
			return false
		}

		num[b%10]--
		b /= 10
	}

	return true
}

func main() {
	begin := time.Now()
	primeArr := make([]int, 0, 5000)

	for n := 3; n < 10000; n += 2 {
		sqrtN := int(math.Sqrt(float64(n)))
		isPrime := true
		for i := 3; i <= sqrtN; i += 2 {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime && n >= 1000 {
			primeArr = append(primeArr, n)
		}
	}

	ans := make([]int, 0, 10)
	for i := 0; i < len(primeArr); i++ {
		for j := i+1; j < len(primeArr); j++ {
			k := primeArr[j] + (primeArr[j]-primeArr[i])
			if k >= 10000 {
				break
			}

			idx := sort.SearchInts(primeArr, k)
			if len(primeArr) != idx && primeArr[idx] == k {
				a, b, c := primeArr[i], primeArr[j], k

				if checkPermutation(a, b) && checkPermutation(a, c) {
					ans = append(ans, a*100000000 + b*10000 + c)
				}
			}
		}
	}

	fmt.Println(time.Since(begin).String())
	for _, v := range ans {
		fmt.Println(v)
	}
}
