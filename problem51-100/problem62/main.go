package main

import (
	"fmt"
	"time"
)

type BitArray uint

func (it *BitArray) Set(i, v uint) {
	bit6 := 1<<7-1
	mov := i*6
	*it &^= BitArray(bit6<<mov)
	*it |= BitArray(v<<mov)
}

func (it *BitArray) Get(i uint) int {
	bit6 := 1<<7-1
	mov := i*6
	result := *it & BitArray(bit6<<mov)
	return int(result>>mov)
}

func cube(n int) int {
	return n*n*n
}

func numToBit(n int) int {
	nums := make([]int, 10)

	for n != 0 {
		nums[n%10]++
		n /= 10
	}

	bitArr := new(BitArray)
	for i, v := range nums {
		bitArr.Set(uint(i), uint(v))
	}
	return int(*bitArr)
}

func main() {
	begin := time.Now()
	maps := make(map[int][]int)

	for n := 100; n < 10000; n++ {
		c := cube(n)
		bit := numToBit(c)
		if _, ok := maps[bit]; !ok {
			maps[bit] = make([]int, 0, 5)
		}
		maps[bit] = append(maps[bit], n)
	}

	ans := 1<<63-1
	for _, v := range maps {
		if len(v) == 5 {
			for _, n := range v {
				c := cube(n)
				if ans > c {
					ans = c
				}
			}
		}
	}

	fmt.Println(time.Since(begin))
	fmt.Println(ans)
}