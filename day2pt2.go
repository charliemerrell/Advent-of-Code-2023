package main

import (
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day2Pt2() int {
	lines := lines.GetLines("./day2.txt")
	sum := 0
	for _, line := range lines {
		splitByColon := strings.Split(line, ":")
		roundsSection := splitByColon[1]
		rounds := strings.Split(roundsSection, ";")
		mins := map[string]int {"green": 0, "red": 0, "blue": 0}
		for _, round := range rounds {
			pickups := strings.Split(round, ", ")
			for _, pickup := range pickups {
				pickup = strings.Trim(pickup, " ")
				num, _ := strconv.Atoi(strings.Split(pickup, " ")[0])
				colour := strings.Split(pickup, " ")[1]
				if mins[colour] < num {
					mins[colour] = num
				}
			}
		}
		sum += (mins["green"] * mins["red"] * mins["blue"])
	}
	return sum
}
