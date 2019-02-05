package main

import (
	"fmt"
	"math/big"
	"time"
)

func comb(n, r int) int {
	a := new(big.Int)
	b := new(big.Int)

	a = a.MulRange(int64(r+1), int64(n))
	b = b.MulRange(1, int64(n-r))

	return len(a.Div(a, b).String())
}

func sol1() int {
	cnt := 0
	for n := 1; n <= 100; n++ {
		for r := 0; r <= n; r++ {
			l := comb(n, r)
			if l >= 7 {
				cnt++
			}
		}
	}

	return cnt
}

func sol2() int {
	cnt := 0

	pascalTriangle := make([][]int, 101)
	for i := range pascalTriangle {
		pascalTriangle[i] = make([]int, 101)
		pascalTriangle[i][0] = 1
	}

	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			pascalTriangle[n][r] = pascalTriangle[n-1][r]+
				pascalTriangle[n-1][r-1]

			if pascalTriangle[n][r] > 1000000 {
				pascalTriangle[n][r] = 1000000
				cnt++
			}
		}
	}

	return cnt
}

func main() {
	sol1Start := time.Now()
	ans := sol1()
	fmt.Println(time.Since(sol1Start).String(), ans)

	sol2Start := time.Now()
	ans2 := sol2()
	fmt.Println(time.Since(sol2Start).String(), ans2)
}
