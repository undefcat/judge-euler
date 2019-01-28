package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	triangles = make(map[int]bool)
)

func triangle(n int) int {
	return n*(n+1)/2
}

func check(b []byte) bool {
	sum := 0
	for _, v := range b {
		sum += int(v-'A')+1
	}

	if _, ok := triangles[sum]; ok {
		return true
	}

	return false
}

func main() {
	start := time.Now()
	// 주요 영어 사전에 있는 단어중 가장 긴 단어가 45글자이므로
	// upper-bound는 ZZZZ...인 45*27로 잡으면 된다.
	for n := 1; n <= 1215; n++ {
		triangles[triangle(n)] = true
	}

	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	count := 0
	isEOF := false
	rd := bufio.NewReader(f)
	for {
		word, err := rd.ReadBytes(',')
		if err != nil {
			if err == io.EOF {
				word = append(word, ',')
				isEOF = true
			} else {
				log.Fatalln(err)
			}
		}

		word = word[1:len(word)-2]

		if check(word) {
			count++
		}

		if isEOF {
			break
		}
	}
	fmt.Println(time.Since(start).String())
	fmt.Println(count)
}
