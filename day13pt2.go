package main

import "aoc2023/lines"

func Day13Pt2() int {
	lines := lines.GetLines("./day13.txt")
	yStart := 0
	sum := 0
	for y, str := range lines {
		if str == "" {
			sum += summarySmudged(lines, yStart, y - 1)
			yStart = y + 1
		}
	}
	sum += summarySmudged(lines, yStart, len(lines) - 1)
	return sum
}

func summarySmudged(lines []string, yStart, yEnd int) int {
	// horizontal
	for cut := yStart; cut < yEnd; cut++ {
		topRow := cut
		botRow := cut + 1
		smudges := 0
		for topRow >= yStart && botRow <= yEnd {
			line1 := lines[topRow]
			line2 := lines[botRow]
			smudges += getSmudges(line1, line2)
			topRow--
			botRow++
		}
		if smudges == 1 {
			return (cut - yStart + 1) * 100
		}
	}

	// vertical
	for cut := 0; cut < len(lines[yStart]) - 1; cut++ {
		leftCol := cut
		rightCol := cut + 1
		smudges := 0
		for leftCol >= 0 && rightCol < len(lines[yStart]) {
			for row := yStart; row <= yEnd; row++ {
				if lines[row][leftCol] != lines[row][rightCol] {
					smudges++
				}
			}
			leftCol--
			rightCol++
		}
		if smudges == 1 {
			return cut + 1
		}
	}
	return 0
}

func getSmudges(l1, l2 string) int {
	s := 0
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			s++
		}
	}
	return s
}