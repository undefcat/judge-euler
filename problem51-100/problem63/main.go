package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	begin := time.Now()

	ans := 0
	for n := 1; n < 10; n++ {
		fn := float64(n)
		sum := fn
		e := 0

		for {
			str := fmt.Sprintf("%e", sum)
			splited := strings.Split(str, "e+")
			exp, _ := strconv.Atoi(splited[1])
			if exp != e {
				break
			}
			e++
			sum *= fn
		}
		ans += e
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}