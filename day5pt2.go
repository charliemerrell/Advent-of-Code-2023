package main

import (
	"aoc2023/lines"
	"fmt"
	"math"
	"strings"
)

func Day5Pt2() int {
	lines := lines.GetLines("./day5.txt")
	seedToSoil := getMaps(lines, 4, 34)
	soilToFert := getMaps(lines, 37, 83)
	fertToWater := getMaps(lines, 86, 130)
	waterToLight := getMaps(lines, 133, 175)
	lightToTemp := getMaps(lines, 178, 197)
	tempToHumid := getMaps(lines, 200, 216)
	humidToLoc := getMaps(lines, 219, 236)
	minLoc := math.MaxInt
	seedRanges := nums(strings.Fields(strings.TrimPrefix(lines[0], "seeds: ")))
	for i := 0; i < len(seedRanges); i += 2 {
		fmt.Println("i: ", i)
		start := seedRanges[i]
		length := seedRanges[i + 1]
		for j := 0; j < length; j++ {
			seed := start + j
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
	}
	return minLoc
}
