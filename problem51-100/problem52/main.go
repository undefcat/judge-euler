package main

import (
	"fmt"
	"time"
)

type Contain []int

func (it Contain) SameDigit(n int) bool {
	tmp := make(Contain, len(it))
	copy(tmp, it)

	for n != 0 {
		num := n%10
		if tmp[num] == 0 {
			return false
		}
		tmp[num]--
		n /= 10
	}

	for i := range tmp {
		if tmp[i] != 0 {
			return false
		}
	}

	return true
}

func NewContain(n int) Contain {
	arr := make(Contain, 10)

	for n != 0 {
		arr[n%10]++
		n /= 10
	}

	return arr
}

func main() {
	begin := time.Now()
	// 17을 6배하면 자리수가 바뀌므로
	// 각 자리수의 17xxx 까지만 확인해보면 된다.
	start := 100
	end := 170
	ans := 0
out:
	for {
		start *= 10
		end *= 10

		for n := start; n < end; n++ {
			contain := NewContain(n)
			if contain.SameDigit(n*2) && contain.SameDigit(n*3) && contain.SameDigit(n*4) &&
				contain.SameDigit(n*5) && contain.SameDigit(n*6) {
				ans = n
				break out
			}
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Print(ans)
}
