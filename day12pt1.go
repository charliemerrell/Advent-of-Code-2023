package main

import (
	// "fmt"
	"aoc2023/lines"
	"strconv"
	"strings"
)

func Day12Pt1() int {
	lines := lines.GetLines("./day12.txt")
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, " ")
		springs, numStr := split[0], strings.Split(split[1], ",")
		nums := make([]int, len(numStr))
		for i, str := range numStr {
			nums[i], _ = strconv.Atoi(str)
		}
		sum += count(springs, nums)
	}
	return sum
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