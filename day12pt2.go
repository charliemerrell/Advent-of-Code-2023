package main

import (
	"aoc2023/lines"
	"fmt"
	"strconv"
	"strings"
)

func Day12Pt2() int {
	lines := lines.GetLines("./day12.txt")
	sum := 0
	for i, line := range lines {
		split := strings.Split(line, " ")
		springs, conditions := split[0], split[1]
		springs = x5Springs(springs)
		conditions = x5Conditions(conditions)
		fmt.Println(i, springs, conditions)
		numStr := strings.Split(conditions, ",")
		nums := make([]int, len(numStr))
		for i, str := range numStr {
			nums[i], _ = strconv.Atoi(str)
		}
		sum += count(springs, nums)
	}
	return sum
}

func x5Springs(springs string) string {
	newStr := strings.Repeat(springs+"?", 5)
	return newStr[0 : len(newStr)-1]
}

func x5Conditions(conditions string) string {
	newStr := strings.Repeat(conditions+",", 5)
	return newStr[0 : len(newStr)-1]
}
