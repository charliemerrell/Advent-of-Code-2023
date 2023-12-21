package main

import (
	// "fmt"
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day12Pt1() int {
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
		go getSum(lines, start, end, ch)
	}
	sum := 0
	for c := 0; c < cores; c++ {
		sum += <-ch
	}
	return sum
}

func getSum(lines []string, start, end int, c chan int) {
	sum := 0
	for i := start; i < end; i++ {
		line := lines[i]
		split := strings.Split(line, " ")
		springs, numStr := split[0], strings.Split(split[1], ",")
		nums := make([]int, len(numStr))
		for i, str := range numStr {
			nums[i], _ = strconv.Atoi(str)
		}
		sum += count(springs, nums)
	}
	c <- sum
}

func count(springs string, nums []int) int {
	for _, c := range springs {
		if c == '?' {
			if !isValidSoFar(springs, nums) {
				return 0
			}
			str1 := strings.Replace(springs, "?", ".", 1)
			str2 := strings.Replace(springs, "?", "#", 1)
			return count(str1, nums) + count(str2, nums)
		}
	}
	if isValid(springs, nums) {
		return 1
	}
	return 0
}

func isValid(springs string, nums []int) bool {
	contigDmg := 0
	i := 0
	for _, c := range springs {
		if c == '#' {
			contigDmg++
		} else {
			if contigDmg != 0 {
				if i == len(nums) {
					return false
				}
				if nums[i] != contigDmg {
					return false
				}
				contigDmg = 0
				i++
			}
		}
	}
	if contigDmg != 0 {
		if i == len(nums) {
			return false
		}
		if nums[i] != contigDmg {
			return false
		}
		contigDmg = 0
		i++
	}
	return i == len(nums)
}

func isValidSoFar(springs string, nums []int) bool {
	contigDmg := 0
	i := 0
	for _, c := range springs {
		if c == '?' {
			break
		}
		if c == '#' {
			contigDmg++
		} else { 
			if contigDmg != 0 {
				if i == len(nums) {
					return false
				}
				if nums[i] != contigDmg {
					return false
				}
				contigDmg = 0
				i++
			}
		}
	}
	return true
}