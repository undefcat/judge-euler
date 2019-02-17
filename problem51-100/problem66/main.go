package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

func getContinuedFraction(d int) (ret []int) {
	ret = make([]int, 1, 1<<4)
	a0 := int(math.Sqrt(float64(d)))
	ret[0] = a0

	n, a, b := d, 1, a0
	for {
		an, c, d := calc(a0, n, a, b)
		ret = append(ret, an)
		a, b = d, c
		if a0<<1 == an {
			return ret
		}
	}
}

func calc(a0, n, a, b int) (an, c, d int) {
	d = (n-(b*b))/a
	an = (a0+b)/d
	c = an*d-b
	return
}

func calcConRat(i, n int, con []int) *big.Rat {
	a := big.NewRat(int64(con[i%len(con)]), 1)

	if i >= n-1 {
		return a
	}

	b := calcConRat(i+1, n, con)
	return a.Add(a, new(big.Rat).Inv(b))
}

func calcPell(x, y, d *big.Int) *big.Int {
	xSquare := new(big.Int).Mul(x, x)
	ySquare := new(big.Int).Mul(y, y)
	mulD := new(big.Int).Mul(d, ySquare)

	return new(big.Int).Sub(xSquare, mulD)
}

func main() {
	begin := time.Now()
	D := make([]int, 0, 1000)

	for n := 2; n <= 1000; n++ {
		sqrtN := int(math.Sqrt(float64(n)))
		if sqrtN*sqrtN == n {
			continue
		}
		D = append(D, n)
	}

	maxX := big.NewInt(0)
	ans := 0

	bigOne := big.NewInt(1)

	debug := false
	target := 13
	for _, d := range D {
		if debug {
			if d != target {
				continue
			}

			if d > target {
				break
			}
		}

		con := getContinuedFraction(d)
		n := 1
		bigD := big.NewInt(int64(d))

		for {
			rat := calcConRat(0, n, con[1:])
			rat.Add(rat.Inv(rat), big.NewRat(int64(con[0]), 1))
			num := rat.Num()
			denom := rat.Denom()

			pell := calcPell(num, denom, bigD)

			if pell.Cmp(bigOne) == 0 {
				if maxX.Cmp(num) < 0 {
					maxX = num
					ans = d
				}

				break
			}
			n++
		}
	}
	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}