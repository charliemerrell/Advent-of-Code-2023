package main

func Day6Pt1() int {
	times := []int{49, 78, 79, 80}
	distances := []int{298, 1185, 1066, 1181}
	out := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distances := distances[i]
		ways := 0
		for t := 0; t <= time; t++ {
			d := t * (time - t)
			if d > distances {
				ways++
			}
		}
		out *= ways
	}
	return out
}
