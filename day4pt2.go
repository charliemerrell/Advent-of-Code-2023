package main

import (
	"aoc2023/lines"
	"slices"
	"strings"
)

func Day4Pt2() int {
	lines := lines.GetLines("./day4.txt")

	copies := make([]int, len(lines))
	for i, _ := range copies {
		copies[i] = 1
	}

	for i, line := range lines {
		line := strings.Split(line, ": ")[1]
		sections := strings.Split(line, " | ")
		winning := strings.Fields(sections[0])
		actual := strings.Fields(sections[1])
		matched := 0
		for _, num := range winning {
			if slices.Contains(actual, num) {
				matched++
			}
		}
		for j := 1; j <= matched; j++ {
			copies[i + j] += 1 * copies[i]
		}
	}
	sum := 0
	for _, count := range copies {
		sum += count
	}
	return sum
}
