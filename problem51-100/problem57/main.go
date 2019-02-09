package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	begin := time.Now()
	memo := make([]*big.Rat, 1000)
	memo[0] = big.NewRat(5, 2)

	one := big.NewRat(1, 1)
	two := big.NewRat(2, 1)

	for i := 1; i < 1000; i++ {
		rat := new(big.Rat).Inv(memo[i-1])
		rat = rat.Add(rat, two)
		memo[i] = rat
	}

	ans := 0
	for i := 0; i < 1000; i++ {
		rat := memo[i].Sub(memo[i], one)
		n, m := rat.Num().String(), rat.Denom().String()

		if len(n) > len(m) {
			ans++
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}
