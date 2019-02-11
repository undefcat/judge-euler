package main

import (
	"fmt"
	"gen"
	"is"
	"strconv"
	"strings"
	"time"
)

var (
	sieve gen.Sieve
)

const Max = 10000000

type Twin struct {
	a, b int
}

type Twins []*Twin

func getTwin(n int) *Twin {
	str := strconv.Itoa(n)
	for i := 1; i < len(str); i++ {
		a, b := atoi(str[:i]), atoi(str[i:])

		if isPrime(a) && isPrime(b) && isPrime(concatStr(str[i:], str[:i])) {
			return &Twin{a, b}
		}
	}
	return nil
}

func isPossible(a, b int) bool {
	return isPrime(concatInt(a, b)) && isPrime(concatInt(b, a))
}

func atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func concatStr(a, b string) int {
	return atoi(strings.Join([]string{a, b}, ""))
}

func concatInt(a, b int) int {
	bb := b
	for bb != 0 {
		a *= 10
		bb /= 10
	}
	return a + b
}

func isPrime(n int) bool {
	if n < Max {
		return sieve.IsPrime(n)
	}

	return is.MiillerPrime(n, 4)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	begin := time.Now()

	sieve = gen.NewSieve(Max)
	primes := make([]int, 4, 10000)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7

	for n, k := 11, 2; n < 1000000; n, k = n+k, 6-k {
		if isPrime(n) {
			primes = append(primes, n)
		}
	}

	twins := make(Twins, 0, 10000)

	for i := range primes {
		twin := getTwin(primes[i])
		if twin != nil {
			twins = append(twins, twin)
		}
	}

	twinPrimes := make([]bool, 100000)

	for i := range twins {
		twinPrimes[twins[i].a] = true
		twinPrimes[twins[i].b] = true
	}

	primes = make([]int, 0, 100000)

	for i, ok := range twinPrimes {
		if ok {
			primes = append(primes, i)
		}
	}

	ans := (1 << 63) - 1
	for a := 0; a < len(primes); a++ {
		if primes[a]*5 > ans {
			break
		}

		for b := a + 1; b < len(primes); b++ {
			if primes[a]+primes[b]*4 > ans {
				break
			}

			if !isPossible(primes[a], primes[b]) {
				continue
			}

			for c := b + 1; c < len(primes); c++ {
				if primes[a]+primes[b]+primes[c]*3 > ans {
					break
				}

				if !isPossible(primes[a], primes[c]) || !isPossible(primes[b], primes[c]) {
					continue
				}

				for d := c + 1; d < len(primes); d++ {
					if primes[a]+primes[b]+primes[c]+primes[d]*2 > ans {
						break
					}

					if !isPossible(primes[a], primes[d]) || !isPossible(primes[b], primes[d]) || !isPossible(primes[c], primes[d]) {
						continue
					}

					for e := d + 1; e < len(primes); e++ {
						if primes[a]+primes[b]+primes[c]+primes[d]+primes[e] > ans {
							break
						}

						if !isPossible(primes[a], primes[e]) || !isPossible(primes[b], primes[e]) || !isPossible(primes[c], primes[e]) ||
							!isPossible(primes[d], primes[e]) {
							continue
						}

						ans = min(ans, primes[a]+primes[b]+primes[c]+primes[d]+primes[e])
					}
				}
			}
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}
