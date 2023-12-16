package main

import (
	"aoc2023/lines"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Day1Pt1() int {
	lines := lines.GetLines("./day1.txt")
	sum := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		num := getLineSum(line)
		sum += num
	}
	return sum
}

func getLineSum(line string) int {
	digitStart := '_'
	digitEnd := '0'
	for j := 0; j < len(line); j++ {
		c, _ := utf8.DecodeRuneInString(line[j:])
		if unicode.IsDigit(c) {
			if digitStart == '_' {
				digitStart = c
			}
			digitEnd = c
		}
	}
	bothDigits := string(digitStart) + string(digitEnd)
	num, _ := strconv.Atoi(bothDigits)
	return num
}
