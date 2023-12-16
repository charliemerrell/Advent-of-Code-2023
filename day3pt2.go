package main

import (
	"aoc2023/lines"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func Day3Pt2() int {
	gearNums := map[string][]int{}
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
					adjacent, num, x, y := gearAdjacent(lines, i, start, end)
					if adjacent {
						key := string(x) + ":" + string(y)
						arr := gearNums[key]
						arr = append(arr, num)
						gearNums[key] = arr
					}
				}
				start = -1
			}
		}
		if start != -1 {
			adjacent, num, x, y := gearAdjacent(lines, i, start, end)
			if adjacent {
				key := string(x) + ":" + string(y)
				arr := gearNums[key]
				arr = append(arr, num)
				gearNums[key] = arr
			}
		}
	}
	sum := 0
	for _, nums := range gearNums {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}
	return sum
}

func gearAdjacent(lines []string, row int, start int, end int) (adjacent bool, num int, gearX int, gearY int) {
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
			if c == '*' {
				num, _ := strconv.Atoi(numStr)
				return true, num, x, y
			}
		}
	}
	return false, 0, 0, 0
}
