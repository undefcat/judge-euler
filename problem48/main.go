package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now()
	sum := 0
	for n := 1; n <= 1000; n++ {
		tmp := n
		for i := 1; i < n; i++ {
			tmp *= n
			tmp %= 10000000000
		}
		sum += tmp
		sum %= 10000000000
	}
	fmt.Println(time.Since(begin).String())
	fmt.Printf("%.10d", sum)
}