package main

import (
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day2Pt1() int {
	lines := lines.GetLines("./day2.txt")
	sum := 0
	for _, line := range lines {
		splitByColon := strings.Split(line, ":")
		gameSection := splitByColon[0]
		gameStr := strings.Split(gameSection, " ")[1]
		game, _ := strconv.Atoi(gameStr)
		roundsSection := splitByColon[1]
		rounds := strings.Split(roundsSection, ";")
		valid := true
		for _, round := range rounds {
			pickups := strings.Split(round, ", ")
			for _, pickup := range pickups {
				pickup = strings.Trim(pickup, " ")
				num, _ := strconv.Atoi(strings.Split(pickup, " ")[0])
				colour := strings.Split(pickup, " ")[1]
				if colour == "green" && num > 13 {
					valid = false
				}
				if colour == "red" && num > 12 {
					valid = false
				}
				if colour == "blue" && num > 14 {
					valid = false
				}
			}
		}
		if valid {
			sum += game
		}
	}
	return sum
}
