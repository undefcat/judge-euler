package main

import (
	"fmt"
	"math"
	"time"
)

type perm []int

var (
	ans int
)

func (it perm) generate() int {
	ret := 0
	for i := range it {
		ret *= 10
		ret += it[i]
	}
	return ret
}

func (it perm) swap(a, b int) {
	it[a], it[b] = it[b], it[a]
}

func heapPermutation(a perm, size int) {
	if size == 1 {
		val := a.generate()
		if isPrime(val) && val > ans {
			ans = val
		}
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(a, size-1)
		if size%2 == 1 {
			a.swap(0, size-1)
		} else {
			a.swap(i, size-1)
		}
	}
}

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	arr := perm{1, 2, 3}
	for i := 4; i < 10; i++ {
		arr = append(arr, i)
		heapPermutation(arr, len(arr))
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}
