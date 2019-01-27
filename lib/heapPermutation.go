package lib

import "fmt"

// 사전순으로 출력되는 것은 아니다.
func heapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
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