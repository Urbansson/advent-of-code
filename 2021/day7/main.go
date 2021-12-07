package main

import (
	"fmt"
	"math"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	crabs := aoc.ExtractInts(lines[0])

	min, max := aoc.MinMax(crabs)

	minFuel := math.MaxInt
	for i := min; i <= max; i++ {
		currFuel := 0
		for _, crab := range crabs {
			v := int(math.Abs(float64(i - crab)))
			currFuel += v
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}
	fmt.Println(minFuel)
}
