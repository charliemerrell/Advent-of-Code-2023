package main

import "aoc2023/lines"

func Day10Pt1() int {
	lines := lines.GetLines("./day10.txt")
	s := s(lines)
	tile1a := s
	tile2a := s
	tile1b := tile{lines[s.y][s.x+1], s.y, s.x + 1}
	tile2b := tile{lines[s.y][s.x-1], s.y, s.x - 1}
	distance := 1
	for {
		tmp1 := next(lines, tile1a, tile1b)
		tile1a, tile1b = tile1b, tmp1
		tmp2 := next(lines, tile2a, tile2b)
		tile2a, tile2b = tile2b, tmp2
		distance++
		if tile1b.x == tile2b.x && tile1b.y == tile2b.y {
			break
		}
	}
	return distance
}

func s(lines []string) tile {
	for y, line := range lines {
		for x, c := range line {
			if c == 'S' {
				return tile{'S', y, x}
			}
		}
	}
	return tile { '_', -1, -1 }
}

func next(lines []string, t1, t2 tile) tile {
	newY := t2.y
	newX := t2.x
	if t2.pipe == '-' {
		if newX+1 == t1.x {
			newX--
		} else {
			newX++
		}
	} else if t2.pipe == 'F' {
		if newX+1 == t1.x {
			newY++
		} else {
			newX++
		}
	} else if t2.pipe == '7' {
		if newX-1 == t1.x {
			newY++
		} else {
			newX--
		}
	} else if t2.pipe == 'J' {
		if newX-1 == t1.x {
			newY--
		} else {
			newX--
		}
	} else if t2.pipe == 'L' {
		if newX+1 == t1.x {
			newY--
		} else {
			newX++
		}
	} else if t2.pipe == '|' {
		if newY+1 == t1.y {
			newY--
		} else {
			newY++
		}
	}
	return tile{lines[newY][newX], newY, newX}
}

type tile struct {
	pipe byte
	y    int
	x    int
}
