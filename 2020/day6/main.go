package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	s := aoc.ReadInput("./input.txt")
	set := make(map[int32]int)
	count := 0
	p := 0
	for _, r := range append(aoc.ExtractLines(s), "") {
		if len(r) == 0 {
			for _, v := range set {
				if v == p {
					count++
				}
			}
			p = 0
			set = make(map[int32]int)
			continue
		}
		for _, c := range r {
			set[c] = set[c] + 1
		}
		p++
	}
	fmt.Println(count)
}
