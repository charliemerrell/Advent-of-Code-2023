package main

import (
	"aoc2023/lines"
	"math"
	"strconv"
	"strings"
)

func Day5Pt1() int {
	lines := lines.GetLines("./day5.txt")
	seeds := nums(strings.Fields(strings.TrimPrefix(lines[0], "seeds: ")))
	seedToSoil := getMaps(lines, 4, 34)
	soilToFert := getMaps(lines, 37, 83)
	fertToWater := getMaps(lines, 86, 130)
	waterToLight := getMaps(lines, 133, 175)
	lightToTemp := getMaps(lines, 178, 197)
	tempToHumid := getMaps(lines, 200, 216)
	humidToLoc := getMaps(lines, 219, 236)
	minLoc := math.MaxInt
	for _, seed := range seeds {
		soil := mapVal(seed, seedToSoil)
		fert := mapVal(soil, soilToFert)
		water := mapVal(fert, fertToWater)
		light := mapVal(water, waterToLight)
		temp := mapVal(light, lightToTemp)
		humid := mapVal(temp, tempToHumid)
		loc := mapVal(humid, humidToLoc)
		if loc < minLoc {
			minLoc = loc
		}
	}
	return minLoc
}

func mapVal(inp int, maps [][3]int) int {
	for _, mp := range maps {
		dest := mp[0]
		source := mp[1]
		ran := mp[2]
		if inp >= source && inp <= source + ran {
			return dest + inp - source
		}
	}
	return inp
}

func nums(strNums []string) []int {
	out := make([]int, len(strNums))
	for i, str := range strNums {
		out[i], _ = strconv.Atoi(str)
	}
	return out
}

func getMaps(lines []string, lineStart int, lineEnd int) [][3]int {
	maps := make([][3]int, (lineEnd-lineStart)+1)
	for i := lineStart; i <= lineEnd; i++ {
		vals := nums(strings.Fields(lines[i-1]))
		maps[i-lineStart][0] = vals[0]
		maps[i-lineStart][1] = vals[1]
		maps[i-lineStart][2] = vals[2]
	}
	return maps
}
