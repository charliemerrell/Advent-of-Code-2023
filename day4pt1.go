package main

import (
	"aoc2023/lines"
	"math"
	"slices"
	"strings"
)

func Day4Pt1() int {
	lines := lines.GetLines("./day4.txt")
	sum := 0
	for _, line := range lines {
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
		if matched > 0 {
			sum += 1 * (int)(math.Pow(2, (float64)(matched - 1)))
		}
	}
	return sum
}
