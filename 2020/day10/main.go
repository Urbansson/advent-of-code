package main

import (
	"fmt"
	"sort"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	a := aoc.ExtractInts(data)
	sort.Ints(a)
	a = append(a, a[len(a)-1]+3)
	fmt.Println(findChain(a))
}

func findChain(a []int) int {
	differences := map[int]int{
		1: 0,
		2: 0,
		3: 0,
	}
	previous := 0
	for _, current := range a {
		difference := current - previous
		differences[difference]++
		previous = current
	}
	return differences[1] * differences[3]
}
