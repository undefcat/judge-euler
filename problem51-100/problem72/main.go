package main

import (
	"fmt"
	"time"
)

func main() {
	begin := time.Now()
	limit := 1000000
	phi := make([]int, limit+1)

	for i := 0; i <= limit; i++ {
		phi[i] = i
	}

	result := 0
	for i := 2; i <= limit; i++ {
		if phi[i] == i {
			for j := i; j <= limit; j += i {
				phi[j] = phi[j]*(i-1)/i
			}
		}
		result += phi[i]
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(result)
}
