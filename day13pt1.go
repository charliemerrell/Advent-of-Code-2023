package main

import "aoc2023/lines"

func Day13Pt1() int {
	lines := lines.GetLines("./day13.txt")
	yStart := 0
	sum := 0
	for y, str := range lines {
		if str == "" {
			sum += summary(lines, yStart, y - 1)
			yStart = y + 1
		}
	}
	sum += summary(lines, yStart, len(lines) - 1)
	return sum
}

func summary(lines []string, yStart, yEnd int) int {
	// horizontal
	for cut := yStart; cut < yEnd; cut++ {
		topRow := cut
		botRow := cut + 1
		valid := true
		for topRow >= yStart && botRow <= yEnd {
			if lines[topRow] != lines[botRow] {
				valid = false
			}
			topRow--
			botRow++
		}
		if valid {
			return (cut - yStart + 1) * 100
		}
	}

	// vertical
	for cut := 0; cut < len(lines[yStart]) - 1; cut++ {
		leftCol := cut
		rightCol := cut + 1
		valid := true
		for leftCol >= 0 && rightCol < len(lines[yStart]) {
			for row := yStart; row <= yEnd; row++ {
				if lines[row][leftCol] != lines[row][rightCol] {
					valid = false
				}
			}
			leftCol--
			rightCol++
		}
		if valid {
			return cut + 1
		}
	}
	return 0
}