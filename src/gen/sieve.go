package gen

import (
	"math"
)

type Sieve interface {
	IsPrime(n int) bool
}

func NewSieve(n int) Sieve {
	return NewBitSieve(n)
}

type boolSieve []bool

func (it boolSieve) IsPrime(n int) bool {
	return it[n]
}

func NewBoolSieve(n int) Sieve {
	sieve := make(boolSieve, n+1)
	sieve[0] = true
	for i := 1; i < len(sieve); i *= 2 {
		copy(sieve[i:], sieve[:i])
	}

	sieve[0], sieve[1] = false, false

	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if sieve[i] {
			for j := i*i; j <= n; j += i {
				sieve[j] = false
			}
		}
	}

	return sieve
}

type bitSieve []uint64

func NewBitSieve(n int) Sieve {
	sieve := make(bitSieve, n/64+1)
	sieve[0] = 1<<64-1
	for i := 1; i < len(sieve); i *= 2 {
		copy(sieve[i:], sieve[:i])
	}

	sieve.setComposite(0)
	sieve.setComposite(1)
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if sieve.IsPrime(i) {
			for j := i*i; j <= n; j += i {
				sieve.setComposite(j)
			}
		}
	}

	return sieve
}

func (it bitSieve) setComposite(n int) {
	it[n>>6] &^= 1 << (uint(n)&63)
}

func (it bitSieve) IsPrime(n int) bool {
	return (it[n>>6] & (1 << (uint(n)&63))) > 0
}
