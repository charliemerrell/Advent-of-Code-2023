package main

import (
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day9Pt2() int {
	lines := lines.GetLines("./day9.txt")
	sum := 0
	for _, line := range lines {
		sum+= getNext2(line)
	}
	return sum
}

func getNext2(s string) int {
	numStrings := strings.Fields(s)
	rows := [][]int{}
	row := make([]int, len(numStrings))
	for i, str := range numStrings {
		row[i], _ = strconv.Atoi(str)
	}
	rows = append(rows, row)
	for !all0(row) {
		newRow := make([]int, len(row)-1)
		for i, _ := range newRow {
			newRow[i] = row[i+1] - row[i]
		}
		row = newRow
		rows = append(rows, row)
	}
	start := 0 
	for i := len(rows) - 1; i > 0; i-- {
		rowAbove := rows[i - 1]
		start = rowAbove[0] - start
	}
	return start
}

