package main

import (
	"aoc2023/lines"
	"regexp"
)

func Day8Pt1() int {
	lines := lines.GetLines("./day8.txt")
	instructions := lines[0]
	pattern := regexp.MustCompile(`[A-Z]{3}`)
	nodes := map[string]node8 {}
	for i := 2; i < len(lines); i++ {
		matches := pattern.FindAllString(lines[i], -1)
		if len(matches) != 3 {
			return 0
		}
		node := node8 { matches[0], matches[1], matches[2] }
		nodes[node.key] = node
	}
	currNode := nodes["AAA"]
	i := 0
	for currNode.key != "ZZZ" {
		instruction := instructions[i % len(instructions)]
		if instruction == 'L' {
			currNode = nodes[currNode.left] 
		}
		if instruction == 'R' {
			currNode = nodes[currNode.right]
		}
		i++
	}
	return i
}

type node8 struct {
	key string
	left string
	right string
}