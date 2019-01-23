package main

import (
	"fmt"
	"math"
	"time"
)

var (
	primes = make(map[int]bool)
)

func isPrime(n int) bool {
	if v, ok := primes[n]; ok {
		return v
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			primes[n] = false
			return false
		}
	}
	primes[n] = true
	return true
}

func isTruncatablePrime(n int) bool {
	if !isPrime(n) {
		return false
	}

	right := n
	left := 0
	mul := 1
	for {
		left += mul*(right%10)
		right /= 10
		if right == 0 {
			break
		}

		if !isPrime(left) || !isPrime(right) {
			return false
		}

		mul *= 10
	}
	return true
}

func main() {
	start := time.Now()
	count := 0
	ans := 0
	ansArr := make([]int, 0, 11)
	primes[1] = false
	primes[2] = true
	primes[3] = true
	primes[5] = true
	primes[7] = true

	for n := 11; count != 11; n += 2 {
		if isTruncatablePrime(n) {
			count++
			ansArr = append(ansArr, n)
			ans += n
		}
	}
	fmt.Println(time.Since(start))
	fmt.Println(ans)
	fmt.Println(ansArr)
}
