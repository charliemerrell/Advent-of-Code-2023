package main

func Day6Pt2() int {
	time := 49787980
	distance := 298118510661181
	ways := 0
	for t := 0; t <= time; t++ {
		if t * (time - t) > distance {
			ways++
		}
	}
	return ways
}