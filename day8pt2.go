package main

import (
	"aoc2023/lines"
	"regexp"
)

func Day8Pt2() uint64 {
	lines := lines.GetLines("./day8.txt")
	instructions := lines[0]
	pattern := regexp.MustCompile(`[A-Z]{3}`)
	nodes := map[string]node8{}
	for i := 2; i < len(lines); i++ {
		matches := pattern.FindAllString(lines[i], -1)
		if len(matches) != 3 {
			return 0
		}
		node := node8{matches[0], matches[1], matches[2]}
		nodes[node.key] = node
	}
	currNodes := []node8{}
	for _, v := range nodes {
		if v.key[2:] == "A" {
			currNodes = append(currNodes, v)
		}
	}
	loopLengths := make([]uint64, len(currNodes))
	for n, currNode := range currNodes {
		i := uint64(0)
		for currNode.key[2:] != "Z" {
			instruction := instructions[i%(uint64)(len(instructions))]
			if instruction == 'L' {
				currNode = nodes[currNode.left]
			}
			if instruction == 'R' {
				currNode = nodes[currNode.right]
			}
			i++
		}
		loopLengths[n] = i
	}
	return lcm(loopLengths)
}

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(numbers []uint64) uint64 {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		result = (result * numbers[i]) / gcd(result, numbers[i])
	}

	return result
}
