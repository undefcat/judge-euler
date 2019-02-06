package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const (
	Jack = 11 + iota
	Queen
	King
	Ace
)

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeKind
	Straight
	Flush
	FullHouse
	FourKind
	StraightFlush
)

const (
	Spade = 1 + iota
	Diamond
	Heart
	Clover
)

type Card struct {
	Suit  int
	Value int
	Group int
}

type Deck []*Card

type Info struct {
	Type       int
	HighNumber []int
}

func (a *Info) win(b *Info) bool {
	if a.Type > b.Type {
		return true
	} else if a.Type < b.Type {
		return false
	} else {
		for i := 0; i < len(a.HighNumber); i++ {
			if a.HighNumber[i] != b.HighNumber[i] {
				return a.HighNumber[i] > b.HighNumber[i]
			}
		}
	}

	panic("no-case")
}

func (it Deck) info() *Info {
	infoArr := make([]int, 9)

	beforeVal := it[0].Value
	beforeSuit := it[0].Suit

	isStraight := true
	isFlush := true
	sameVal := 1

	for i := 1; i < len(it); i++ {
		if isStraight && it[i].Value != beforeVal+1 {
			isStraight = false
		}

		if isFlush && it[i].Suit != beforeSuit {
			isFlush = false
		}

		if it[i].Value == beforeVal {
			sameVal++

		} else {
			group := HighCard

			switch sameVal {
			case 2:
				infoArr[OnePair]++
				group = OnePair

			case 3:
				infoArr[ThreeKind]++
				group = ThreeKind

			case 4:
				infoArr[FourKind]++
				group = FourKind
			}

			for j := 1; j <= sameVal; j++ {
				it[i-j].Group = group
			}

			sameVal = 1
		}

		beforeVal = it[i].Value
	}

	// 5번 다 돌았을 때에
	// 분류를 해줘야 한다.
	group := HighCard

	if isFlush {
		infoArr[Flush]++
		group = Flush
	}

	if isStraight {
		infoArr[Straight]++
		group = Straight
	}

	if infoArr[Straight]*infoArr[Flush] > 0 {
		infoArr[StraightFlush]++
		infoArr[Straight]--
		infoArr[Flush]--
		group = StraightFlush

	} else {
		switch {
		case sameVal == 4:
			infoArr[FourKind]++
			group = FourKind

		case infoArr[OnePair] == 2:
			infoArr[OnePair] = 0
			infoArr[TwoPair]++

		case infoArr[OnePair] == 1 && sameVal == 2:
			infoArr[TwoPair]++
			infoArr[OnePair]--
			group = OnePair

		case infoArr[TwoPair] == 1 && sameVal == 3:
			infoArr[TwoPair]--
			infoArr[FullHouse]++
			group = ThreeKind

		case infoArr[ThreeKind] == 1 && sameVal == 2:
			infoArr[ThreeKind]--
			infoArr[FullHouse]++
			group = OnePair
		}
	}

	lastIdx := len(it)-1

	if isStraight || isFlush {
		for i := 0; i < len(it); i++ {
			it[i].Group = group
		}
	} else {
		for i := 0; i < sameVal; i++ {
			it[lastIdx-i].Group = group
		}
	}

	info := new(Info)
	for i, val := range infoArr {
		if val != 0 {
			info.Type = i
			break
		}
	}

	// 5장의 카드가 하나의 패가 되는 경우
	// HighNumber를 따로 지정해줘야 한다.
	// 그 이외의 경우에는 남은 HighCard의 가장 높은 번호로 지정하면 된다.
	highNumber := make([]int, 0, len(it))

	switch info.Type {
	case StraightFlush, Straight, Flush:
		for i := lastIdx; i >= 0; i-- {
			highNumber = append(highNumber, it[i].Value)
		}

	case FullHouse:
		for i := lastIdx; i >= 0; i-- {
			if it[i].Group == ThreeKind {
				highNumber = append(highNumber, it[i].Value)
				break
			}
		}

	default:
		for i := lastIdx; i >= 0; i-- {
			if it[i].Group == HighCard {
				highNumber = append(highNumber, it[i].Value)
			}
		}
	}

	info.HighNumber = highNumber

	return info
}

func (it Deck) Len() int {
	return len(it)
}

func (it Deck) Less(i, j int) bool {
	if it[i].Value != it[j].Value {
		return it[i].Value < it[j].Value
	}

	return it[i].Suit < it[j].Suit
}

func (it Deck) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func makeDeck(b [][]byte) Deck {
	deck := make(Deck, 5)

	for i := range deck {
		deck[i] = makeCard(b[i])
	}

	sort.Sort(deck)
	return deck
}

func makeCard(b []byte) *Card {
	var value int

	switch b[0] {
	case 'T':
		value = 10

	case 'A':
		value = Ace

	case 'J':
		value = Jack

	case 'Q':
		value = Queen

	case 'K':
		value = King

	default:
		value = int(b[0] - '0')
	}

	var suit int

	switch b[1] {
	case 'C':
		suit = Clover

	case 'S':
		suit = Spade

	case 'D':
		suit = Diamond

	case 'H':
		suit = Heart
	}

	return &Card{
		Value: value,
		Suit:  suit,
	}
}

func main() {
	begin := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	ans := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		input := sc.Bytes()
		data := bytes.Fields(input)
		p1 := makeDeck(data[:5])
		p2 := makeDeck(data[5:])
		p1Info := p1.info()
		p2Info := p2.info()

		if p1Info.win(p2Info) {
			ans++
		}
	}

	fmt.Println(time.Since(begin).String())
	fmt.Println(ans)
}
