package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	fishes := aoc.ExtractInts(lines[0])

	var (
		M         = 10
		dueDates  = make([]int, M)
		dueDate   int
		fish      = 0
		today     int
		newFish   int
		newFishDD int
		oldFishDD int
	)

	for _, ini := range fishes {
		dueDate = ini + 1
		dueDates[dueDate]++
		fish++
	}

	// simulate births for N days
	for day := 1; day <= 256; day++ {
		today = day % M
		// birth new fish
		newFish = dueDates[today]
		fish += newFish

		// update due dates
		newFishDD = (today + 9) % M
		oldFishDD = (today + 7) % M
		dueDates[newFishDD] += newFish
		dueDates[oldFishDD] += newFish
		dueDates[today] -= newFish
	}

	fmt.Println(fish)
}
