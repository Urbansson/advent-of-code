package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Manhattan(o Coordinate) int {
	return int(math.Abs(float64(c.X-o.X)) + math.Abs(float64(c.Y-o.Y)))
}

func (c Coordinate) Move(d Direction, n int) Coordinate {
	switch d {
	case Up:
		c.Y += n
	case Right:
		c.X += n
	case Down:
		c.Y -= n
	case Left:
		c.X -= n
	}
	return c
}

// Walk calls fn for every step
func (c Coordinate) Walk(d Direction, n int, fn func(Coordinate)) Coordinate {
	for i := 0; i < n; i++ {
		c = c.Move(d, 1)
		fn(c)
	}
	return c
}

type Direction int

const (
	Up Direction = iota + 1
	Right
	Down
	Left
)

func directionFromUPLR(s byte) Direction {
	switch s {
	case 'U':
		return Up
	case 'D':
		return Down
	case 'L':
		return Left
	case 'R':
		return Right
	}
	panic("invalid UPLR")
}

func main() {
	input := utils.ReadInput("day03/input.txt")
	data := utils.ExtractLines(input)
	panel := make(map[Coordinate]int)
	var c1 Coordinate
	steps := 0
	for _, i := range strings.Split(data[0], ",") {
		dir := directionFromUPLR(i[0])
		step := utils.Atoi(i[1:])
		c1 = c1.Walk(dir, step, func(c Coordinate) {
			steps++
			if _, ok := panel[c]; !ok {
				panel[c] = steps
			}
		})
	}
	steps = 0
	intersections := make(map[Coordinate]int)
	var c2 Coordinate
	for _, i := range strings.Split(data[1], ",") {
		dir := directionFromUPLR(i[0])
		step := utils.Atoi(i[1:])
		c2 = c2.Walk(dir, step, func(c Coordinate) {
			steps++
			if val, ok := panel[c]; ok {
				intersections[c] = steps + val
			}
		})
	}
	best := math.MaxInt64
	for i := range intersections {
		best = int(math.Min(float64(best), float64(i.Manhattan(Coordinate{0, 0}))))
	}
	fmt.Println(best)
	bestSteps := math.MaxInt64
	for _, steps := range intersections {
		bestSteps = int(math.Min(float64(bestSteps), float64(steps)))
	}
	fmt.Println(bestSteps)

}
