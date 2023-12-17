package main

import (
	"aoc2023/lines"
	"sort"
	"strconv"
	"strings"
)

func Day7Pt1() int {
	hands := lines.GetLines("./day7.txt")
	sort.Slice(hands, func(i, j int) bool { return greater(hands[j], hands[i]) })
	sum := 0
	for i, hand := range hands {
		bid, _ := strconv.Atoi(strings.Split(hand, " ")[1])
		sum += bid * (i + 1)
	}
	return sum
}

func greater(hand1 string, hand2 string) bool {
	cards1 := strings.Split(hand1, " ")[0]
	cards2 := strings.Split(hand2, " ")[0]
	rank1 := getRank(cards1)
	rank2 := getRank(cards2)
	if rank1 == rank2 {
		for i := 0; i < len(cards1); i ++ {
			rank1 = strings.Index("AKQJT98765432", string(cards1[i]))
			rank2 = strings.Index("AKQJT98765432", string(cards2[i]))
			if (rank1 != rank2) {
				return rank1 < rank2
			}
		}
	} else {
		return rank1 < rank2
	}
	return true // shouldn't hit this
}

func getRank(hand string) int {
	counts := map[rune]int {}
	for _, c := range hand {
		counts[c]++ 
	}
	has5 := false
	has4 := false 
	has3 := false
	has2 := false
	has2pt2 := false
	for _, c := range counts {
		if c == 5 { has5 = true }
		if c == 4 { has4 = true }
		if c == 3 { has3 = true }
		if c == 2 { 
			if has2 { has2pt2 = true }
			has2 = true 
		}
	}
	if has5 { return 1 }
	if has4 { return 2 }
	if has3 && has2 { return 3 }
	if has3 { return 4 }
	if has2pt2 { return 5 }
	if has2 { return 6 }
	return 7
}