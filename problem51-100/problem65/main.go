package main

import (
	"fmt"
	"math/big"
	"time"
)

func reverse(s []*big.Rat) {
	mid := len(s)/2
	last := len(s)-1

	for i := 0; i < mid; i++ {
		s[i], s[last-i] = s[last-i], s[i]
	}
}

func main() {
	begin := time.Now()

	fraction := make([]*big.Rat, 99)
	fraction[0] = big.NewRat(1, 1)

	for i := 1; i < len(fraction); i *= 2 {
		copy(fraction[i:], fraction[:i])
	}

	for i := 1; 3*i-2 < len(fraction); i++ {
		rat := big.NewRat(int64(2*i), 1)
		fraction[3*i-2] = rat
	}

	reverse(fraction)
	fraction = append(fraction, big.NewRat(2, 1))

	before := new(big.Rat)
	before.Add(fraction[1], fraction[0].Inv(fraction[0]))

	for i := 2; i < len(fraction); i++ {
		before.Add(fraction[i], before.Inv(before))
	}

	num := before.Num().String()
	sum := 0
	for _, v := range num {
		sum += int(v-'0')
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(sum)
}
