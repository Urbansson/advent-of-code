package main

import (
	"fmt"
	"sort"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	elves := []int{}
	temp := 0
	for _, l := range lines {
		if l == "" {
			elves = append(elves, temp)
			temp = 0
		} else {
			temp += aoc.Atoi(l)
		}
	}
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
