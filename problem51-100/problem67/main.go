package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	triangle [][]int
	dp       [][]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	begin := time.Now()
	inputTxt, _ := os.Open("input.txt")
	sc := bufio.NewScanner(inputTxt)

	triangle = make([][]int, 100)
	dp = make([][]int, 100)

	for i := 0; i < 100; i++ {
		sc.Scan()
		line := sc.Bytes()
		splited := bytes.Fields(line)
		ints := make([]int, len(splited))

		for i := range splited {
			n, _ := strconv.Atoi(string(splited[i]))
			ints[i] = n
		}

		triangle[i] = ints
		dp[i] = make([]int, len(ints))
	}

	dp[0][0] = triangle[0][0]

	for y := 1; y < len(triangle); y++ {
		dp[y][0] += dp[y-1][0] + triangle[y][0]
	}

	for y := 1; y < len(triangle); y++ {
		for x := 1; x < len(triangle[y]); x++ {
			dp[y][x] = max(dp[y][x], dp[y-1][x-1]+triangle[y][x-1])
			if x != len(triangle[y])-1 {
				dp[y][x] = max(dp[y][x], dp[y-1][x]+triangle[y][x])
			}
		}
	}

	ans := 0
	for x := 0; x < len(dp[99]); x++ {
		ans = max(ans, dp[99][x])
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}