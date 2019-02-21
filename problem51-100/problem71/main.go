package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now()
	a, b := 3, 7
	r, s := 0, 1

	for q := 1000000; q >= 2; q-- {
		p := ((a*q)-1)/b
		if p*s > r*q {
			s, r = q, p
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Printf("%d/%d\n", r, s)
}
