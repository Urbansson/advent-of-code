package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	for _, l := range lines {
		e := strings.Replace(l, "-", " ", -1)
		e = strings.Replace(e, ",", " ", -1)
		ie := aoc.IntList(e)

		r1 := aoc.Range{Start: ie[0], End: ie[1]}
		r2 := aoc.Range{Start: ie[2], End: ie[3]}

		if contained(r1, r2) || contained(r2, r1) {
			sum++
		}

	}
	fmt.Println(sum)
}

func contained(r1, r2 aoc.Range) bool {
	return r1.Start >= r2.Start && r1.End <= r2.End
}
