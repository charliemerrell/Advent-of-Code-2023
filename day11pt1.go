package main

import (
	"aoc2023/lines"
)

func Day11Pt1() int {
	grid := lines.GetLines("./day11.txt")
	galaxies := []galaxy{}
	galaxyRows := make([]bool, len(grid))
	galaxyCols := make([]bool, len(grid[0]))
	for y, r := range grid {
		for x, c := range r {
			if c == '#' {
				galaxies = append(galaxies, galaxy{y, x})
				galaxyCols[x] = true
				galaxyRows[y] = true
			}
		}
	}
	sum := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if j <= i {
				continue
			}
			yTop := min(g1.y, g2.y)
			yBot := max(g1.y, g2.y)
			for y := yTop; y < yBot; y++ {
				sum++
				if !galaxyRows[y] {
					sum++
				}
			}
			xRight := max(g1.x, g2.x)
			xLeft := min(g1.x, g2.x)
			for x := xLeft; x < xRight; x++ {
				sum++
				if !galaxyCols[x] {
					sum++
				}
			}
		}
	}
	return sum
}

type galaxy struct {
	y int
	x int
}
