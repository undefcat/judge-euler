package main

import (
	"fmt"
	"math"
	"time"
)

func isPentagonal(n int) bool {
	numerator := math.Sqrt(float64(24*n+1))+1
	result := numerator/6
	return result == float64(int(result))
}

func isHexagonal(n int) bool {
	numerator := math.Sqrt(float64(8*n+1))+1
	result := numerator/4

	return result == float64(int(result))
}

func main() {
	start := time.Now()
	i := 286
	ans := 0
	ansI := 0
	for {
		n := i*(i+1)/2
		if isPentagonal(n) && isHexagonal(n) {
			ans = n
			ansI = i
			break
		}
		i++
	}

	fmt.Println(time.Since(start).String())
	fmt.Printf("Tn(%d) == %d\n", ansI, ans)
}
