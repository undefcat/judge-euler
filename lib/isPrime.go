package lib

import "math"

// 가장 간단한 소수판별 알고리즘
func isPrime1(n int) bool {
	if n == 1 {
		return false
	}

	if n == 2 {
		return true
	}

	// 짝수면 소수 아님
	if n%2 == 0 {
		return false
	}

	// 제곱근을 구하고
	// 홀수인 수들에 대해서 판별
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
