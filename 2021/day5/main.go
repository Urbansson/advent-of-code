package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	vents := [1000][1000]int{}

	for _, l := range lines {
		cords := strings.Split(l, " -> ")
		start := aoc.ExtractInts(cords[0])
		end := aoc.ExtractInts(cords[1])

		addX := 0
		addY := 0
		if start[0] > end[0] {
			addX = -1
		}
		if start[0] < end[0] {
			addX = 1
		}
		if start[1] > end[1] {
			addY = -1
		}
		if start[1] < end[1] {
			addY = 1
		}
		startX := start[0]
		startY := start[1]
		targetX := end[0]
		targetY := end[1]
		for startX != targetX || startY != targetY {
			vents[startY][startX] += 1
			startX += addX
			startY += addY
		}
		vents[startY][startX] += 1
	}
	sum := 0
	for i := range vents {
		for k := range vents[i] {
			if vents[i][k] >= 2 {
				sum += 1
			}
		}
	}
	fmt.Println(sum)
}
