package main

import (
	"fmt"
	"math"
	"time"
)

func calc(a0, n, a, b int) (an, c, d int) {
	c = (n-(a*a))/b
	an = (a0+a)/c
	d = an*c-a

	return
}

func main() {
	begin := time.Now()
	isContinue := make([]bool, 10001)
	ans := 0
	for n := 2; n*n <= 10000; n++ {
		isContinue[n*n] = true
	}

	for n := 2; n <= 10000; n++ {
		if isContinue[n] {
			continue
		}

		a0 := int(math.Sqrt(float64(n)))
		a := a0
		b := 1

		peri := 1
		for {
			an, c, d := calc(a0, n, a, b)
			if an == a0*2 {
				if peri%2 == 1 {
					ans++
				}
				break
			}
			a, b = d, c
			peri++
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}
