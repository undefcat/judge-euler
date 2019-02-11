package is

import (
	"math/rand"
	"time"
)

func Prime(n int) bool {
	switch {
	case n == 1:
		return false

	case n == 2, n == 3:
		return true

	case n%2 == 0, n%3 == 0:
		return false
	}

	i, w := 5, 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}

		i, w = i+w, 6-w
	}
	return true
}

func MiillerPrime(n, k int) bool {
	if n <= 1 || n == 4 {
		return false
	}

	if n <= 3 {
		return true
	}

	d := n - 1

	for d%2 == 0 {
		d /= 2
	}

	for i := 0; i < k; i++ {
		if !miillerTest(d, n) {
			return false
		}
	}

	return true
}

var rd = rand.New(rand.NewSource(time.Now().Unix()))

func miillerTest(d, n int) bool {
	a := 2 + rd.Int()%(n-4)

	x := power(a, d, n)

	if x == 1 || x == n-1 {
		return true
	}

	for d != n-1 {
		x = (x * x) % n
		d *= 2

		if x == 1 {
			return false
		}
		if x == n-1 {
			return true
		}
	}

	return false
}

func power(x, y, p int) int {
	res := 1
	x = x%p

	for y > 0 {
		if y & 1 > 0 {
			res = (res*x)%p
		}
		y = y>>1
		x = (x*x)%p
	}
	return res
}