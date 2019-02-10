package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

type Data struct {
	char byte
	num int
}

type Datas []*Data

func (it Datas) Len() int {
	return len(it)
}

func (it Datas) Less(i, j int) bool {
	return it[i].num > it[j].num
}

func (it Datas) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func decrypt(m, key []byte) []byte {
	ret := make([]byte, len(m))

	for i, v := range m {
		ret[i] = v^key[i%3]
	}

	return ret
}

func analysis(m []byte) bool {
	maxSize := 0
	for _, v := range m {
		if int(v) > maxSize {
			maxSize = int(v)
		}
	}

	keys := make(Datas, maxSize+1)

	for i := range keys {
		keys[i] = &Data{byte(i), 0}
	}

	for _, v := range m {
		keys[v].num++
	}

	sort.Sort(keys)

	frequently := 0
	for i := 0; i < 6; i++ {
		switch keys[i].char {
		case ' ', 'e', 't', 'a', 'o', 'i', 'n':
			frequently++
		}
	}

	return frequently >= 4
}

func main() {
	begin := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sc := bufio.NewScanner(f)
	sc.Scan()

	tmp := sc.Bytes()
	splited := bytes.Split(tmp, []byte{','})
	byteMsg := make([]byte, len(splited))

	for i := range splited {
		n, _ := strconv.Atoi(string(splited[i]))
		byteMsg[i] = byte(n)
	}

	out, err := os.Create("out.txt")
	if err != nil {
		log.Fatalln(err)
	}

	wr := bufio.NewWriter(out)

	keys := make([]byte, 3)
	for keyA := 'a'; keyA <= 'z'; keyA++ {
		keys[0] = byte(keyA)
		for keyB := 'a'; keyB <= 'z'; keyB++ {
			keys[1] = byte(keyB)
			for keyC := 'a'; keyC <= 'z'; keyC++ {
				keys[2] = byte(keyC)

				decrypted := decrypt(byteMsg, keys)
				if analysis(decrypted) {
					wr.WriteString("key:")
					for _, v := range keys {
						wr.WriteByte(v)
					}
					wr.WriteByte('\n')
					wr.WriteString(string(decrypted))
					wr.WriteByte('\n')
					wr.WriteByte('\n')
				}
			}
		}
	}
	wr.Flush()
	// key is exp

	decrypted := decrypt(byteMsg, []byte{'e', 'x', 'p'})
	sum := 0
	for _, v := range decrypted {
		sum += int(v)
	}
	fmt.Println(time.Since(begin).String())
	fmt.Println(sum)
}