package main

import (
	"fmt"
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day12Pt2() int {
	lines := lines.GetLines("./day12.txt")
	const cores = 16
	n := len(lines)
	nPerThread := n / cores
	ch := make(chan int)
	for c := 0; c < cores; c++ {
		start := c * nPerThread
		end := start + nPerThread
		if c == cores - 1 {
			end = n
		}
		go getSumX5(lines, start, end, ch)
	}
	sum := 0
	for c := 0; c < cores; c++ {
		sum += <-ch
	}
	return sum
}

func getSumX5(lines []string, start, end int, c chan int) {
	sum := 0
	for i := start; i < end; i++ {
		line := lines[i]
		split := strings.Split(line, " ")
		springs, conditions := split[0], split[1]
		springs = x5Springs(springs)
		conditions = x5Conditions(conditions)
		fmt.Println(i)
		numStr := strings.Split(conditions, ",")
		nums := make([]int, len(numStr))
		for i, str := range numStr {
			nums[i], _ = strconv.Atoi(str)
		}
		sum += count(springs, nums)
	}
	c <- sum
}

func x5Springs(springs string) string {
	newStr := strings.Repeat(springs+"?", 5)
	return newStr[0 : len(newStr)-1]
}

func x5Conditions(conditions string) string {
	newStr := strings.Repeat(conditions+",", 5)
	return newStr[0 : len(newStr)-1]
}
