package main

import (
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	n := [][]int{}
	for _, line := range lines[:len(lines)-1] {
		il := aoc.IntList(line)
		n = append(n, il)
	}
	ops := strings.Fields(lines[len(lines)-1])

	ss := 0
	for i := range n[0] {
		op := ops[i]
		sum := 0
		for j := range n {
			f := n[j][i]
			if op == "*" {
				if sum == 0 {
					sum++
				}
				sum *= f
			} else {
				sum += f
			}
		}
		ss += sum
	}
}
