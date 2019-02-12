package main

import (
	"fmt"
	"time"
)

const (
	Tri = iota
	Squ
	Pen
	Hex
	Hep
	Oct
)

var (
	nums []map[int]bool
	adj  [][][]int
	ans  []int
)

func triangle(n int) int {
	return n * (n + 1) / 2
}

func square(n int) int {
	return n * n
}

func pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func hexagonal(n int) int {
	return n * (2*n - 1)
}

func heptagonal(n int) int {
	return n * (5*n - 3) / 2
}

func octagonal(n int) int {
	return n * (3*n - 2)
}

func split(n int) (int, int) {
	return n / 100, n % 100
}

func makeGraph() {
	// create adj
	adj = make([][][]int, 100)
	for i := range adj {
		adj[i] = make([][]int, 100)
	}

	// connect edges
	for i := range nums {
		for num := range nums[i] {
			a, b := split(num)
			if adj[a][b] == nil {
				adj[a][b] = make([]int, 6)
			}

			adj[a][b][i] = num
		}
	}
}

func traverse() {
	octaPrefix := make([]int, 0, len(nums[Oct]))

	for num := range nums[Oct] {
		a, _ := split(num)
		octaPrefix = append(octaPrefix, a)
	}

	for i := range octaPrefix {
		used := make([]bool, 6)
		path := make([]int, 6)
		if dfs(octaPrefix[i], 6, used, path) {
			break
		}
	}
}

func dfs(here, remain int, used []bool, path []int) bool {
	if remain == 0 {
		_, a := split(path[5])
		b, _ := split(path[0])

		if a == b {
			ans = make([]int, 6)
			copy(ans, path)
			return true
		}
		return false
	}

	for v := range adj[here] {
		if adj[here][v] == nil {
			continue
		}

		for poly, num := range adj[here][v] {
			if !used[poly] && num != 0 {
				used[poly] = true
				path[6-remain] = num
				if !dfs(v, remain-1, used, path) {
					used[poly] = false
				} else {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	begin := time.Now()
	nums = make([]map[int]bool, 6)
	for i := range nums {
		nums[i] = make(map[int]bool)
	}

	funcs := []func(int) int{triangle, square, pentagonal, hexagonal, heptagonal, octagonal}

	for i := range funcs {
		n := 1
		for {
			value := funcs[i](n)
			n++
			if value < 1000 {
				continue
			} else if value >= 10000 {
				break
			}
			nums[i][value] = true
		}
	}

	makeGraph()
	traverse()
	sum := 0
	for _, v := range ans {
		sum += v
	}
	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
	fmt.Println(sum)
}
