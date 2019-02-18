package main

import (
	"fmt"
	"gen"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Group [][]int

func grouped(l []int) Group {
	group := make(Group, 5)
	group[0] = []int{10, l[0], l[1]}
	group[1] = []int{l[2], l[1], l[3]}
	group[2] = []int{l[4], l[3], l[5]}
	group[3] = []int{l[6], l[5], l[7]}
	group[4] = []int{l[8], l[7], l[0]}

	return group
}

func isMagicGonRing(group Group) (bool, int) {
	pivot := groupSum(group[0])
	for i := 1; i < len(group); i++ {
		if groupSum(group[i]) != pivot {
			return false, 0
		}
	}
	return true, pivot
}

func groupSum(l []int) int {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return sum
}

func concat(g Group) string {
	minEdge := 1<<4
	minIdx := -1
	for i := range g {
		if g[i][0] < minEdge {
			minEdge = g[i][0]
			minIdx = i
		}
	}

	buf := make([]string, 0, 5)
	for i := 0; i < len(g); i++ {
		for _, v := range g[(i+minIdx)%len(g)] {
			buf = append(buf, strconv.Itoa(v))
		}
	}
	return strings.Join(buf, "")
}

func main() {
	begin := time.Now()
	perm := gen.NewHeapPermutation([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	table := make(map[int][]Group)

	perm.Callback(func (l []int) {
		group := grouped(l)
		if ok, pivot := isMagicGonRing(group); ok {
			copied := make(Group, len(group))
			for i := range copied {
				tmp := make([]int, len(group[i]))
				copy(tmp, group[i])
				copied[i] = tmp
			}

			if _, ok := table[pivot]; !ok {
				table[pivot] = make([]Group, 0, 3)
			}
			table[pivot] = append(table[pivot], group)
		}
	})

	perm.Gen()
	candidate := make([]string, 0, 10)
	for _, v := range table {
		for _, g := range v {
			candidate = append(candidate, concat(g))
		}
	}
	sort.Strings(candidate)

	fmt.Println(time.Since(begin).String())
	fmt.Println(candidate[len(candidate)-1])
}
