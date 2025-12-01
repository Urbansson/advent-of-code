package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	pass := 0
	dail := 50
	for _, l := range lines {
		op := l[0:1]
		v := aoc.Atoi(l[1:])

		if op == "L" {
			dail -= v
			for dail < 0 {
				dail += 100
			}
		}
		if op == "R" {
			dail += v
			for dail >= 100 {
				dail -= 100
			}
		}
		if dail == 0 {
			pass += 1
		}
	}
	fmt.Printf("pass: %d\n", pass)
}
