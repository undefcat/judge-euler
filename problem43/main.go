package main

import (
	"fmt"
	"time"
)

var ans int

func heapPermutation(a []int, size int) {
	if size == 1 {
		if check(a) {
			ans += sum(a)
		}
		return
	}

	for i := 0; i < size; i++ {
		dec := size-1
		heapPermutation(a, dec)
		if size%2 == 1 {
			a[0], a[dec] = a[dec], a[0]
		} else {
			a[i], a[dec] = a[dec], a[i]
		}
	}
}

func check(arr []int) bool {
	if arr[0] == 0 {
		return false
	}

	if arr[3]%2 != 0 {
		return false
	}

	if (arr[2] + arr[3] + arr[4])%3 != 0 {
		return false
	}

	if arr[5]%5 != 0 {
		return false
	}

	if (arr[4]*100 + arr[5]*10 + arr[6])%7 != 0 {
		return false
	}

	if (arr[5]*100 + arr[6]*10 + arr[7])%11 != 0 {
		return false
	}

	if (arr[6]*100 + arr[7]*10 + arr[8])%13 != 0 {
		return false
	}

	if (arr[7]*100 + arr[8]*10 + arr[9])%17 != 0 {
		return false
	}

	return true
}

func sum(a []int) int {
	ans := 0
	for _, v := range a {
		ans *= 10
		ans += v
	}

	return ans
}

func main() {
	start := time.Now()
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heapPermutation(arr, len(arr))
	fmt.Println(time.Since(start).String())
	fmt.Println(ans)
}
