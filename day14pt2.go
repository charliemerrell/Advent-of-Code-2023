package main

import (
	"aoc2023/lines"
)

func Day14Pt2() int {
	lines := lines.GetLines("./day14.txt")
	grid := make([][]byte, len(lines))
	for y := range lines {
		grid[y] = make([]byte, len(lines[y]))
		for x := range lines[y] {
			grid[y][x] = lines[y][x]
		}
	}

	sum := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				for newY := y - 1; newY >= 0 && grid[newY][x] == '.'; newY-- {
					grid[newY+1][x] = '.'
					grid[newY][x] = 'O'
				}
			}
		}
	}
	printGrid(grid)

	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'O' {
				sum += len(grid) - y
			}
		}
	}
	return sum
}
