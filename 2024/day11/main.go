package main

import (
	"fmt"
	"slices"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	stones := aoc.ExtractInts(aoc.ReadStdin())
	for range 25 {
		for i := 0; i < len(stones); i++ {
			v := stones[i]
			r := applyRule(v)
			if len(r) == 1 {
				stones = slices.Replace(stones, i, i+1, r...)
			} else {
				stones = slices.Replace(stones, i, i+1, r[0])
				stones = slices.Insert(stones, i+1, r[1])
				i += 1
			}
		}
	}
	fmt.Println("---")
	fmt.Println(len(stones))
}

func applyRule(r int) []int {
	if r == 0 {
		return []int{1}
	}
	d := aoc.Digits(r)
	if len(d)%2 == 0 {
		h := len(d) / 2

		f := aoc.DigitsToInt(d[:h])
		s := aoc.DigitsToInt(d[h:])

		return []int{f, s}
	}
	return []int{r * 2024}
}
