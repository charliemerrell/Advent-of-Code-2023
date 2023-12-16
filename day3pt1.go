package main

import (
	"aoc2023/lines"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Day3Pt1() int {
	sum := 0
	lines := lines.GetLines("./day3.txt")
	for i, line := range lines {
		start := -1
		end := -1
		for j, c := range line {
			if unicode.IsDigit(c) {
				if start == -1 {
					start = j
					end = j
				} else {
					end = j
				}
			} else {
				if start != -1 {
					sum += getValue(lines, i, start, end)
				}
				start = -1
			}
		}
		if start != -1 {
			sum += getValue(lines, i, start, end)
		}
	}
	return sum
}

func getValue(lines []string, row int, start int, end int) int {
	numStr := lines[row][start : end+1]
	startY := row - 1
	endY := row + 1
	startX := start - 1
	endX := end + 1
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if x < 0 || x >= len(lines[0]) {
				continue
			}
			if y < 0 || y >= len(lines) {
				continue
			}
			c, _ := utf8.DecodeRuneInString(lines[y][x:])
			if c != '.' && !unicode.IsDigit(c) {
				num, _ := strconv.Atoi(numStr)
				return num
			}
		}
	}
	return 0
}
