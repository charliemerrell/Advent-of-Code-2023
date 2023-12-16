package main

import (
	"aoc2023/lines"
	"unicode"
	"unicode/utf8"
)

func Day1Pt2() int {
	lines := lines.GetLines("./day1.txt")
	sum := 0
	numbers := map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		inc := newFunction(line, numbers)
		sum += inc
	}
	return sum
}

func newFunction(line string, numbers map[string]int) int {
	digitStart := -1
	digitEnd := 0
	for j := 0; j < len(line); j++ {
		c, _ := utf8.DecodeRuneInString(line[j:])
		if unicode.IsDigit(c) {
			if digitStart == -1 {
				digitStart = int(c - '0')
			}
			digitEnd = int(c - '0')
		} else {
			for numStr, numInt := range numbers {
				if j+len(numStr) <= len(line) {
					chars := line[j : j+len(numStr)]
					if chars == numStr {
						if digitStart == -1 {
							digitStart = numInt
						}
						digitEnd = numInt
					}
				}
			}
		}
	}
	inc := (digitStart * 10) + digitEnd
	return inc
}
