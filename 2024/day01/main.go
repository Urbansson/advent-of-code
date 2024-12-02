package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	ll := make([]int, 0, len(lines))
	m := make(map[int]int)

	for _, l := range lines {
		p := strings.Split(l, "   ")
		ll = append(ll, aoc.Atoi(p[0]))
		mv := m[aoc.Atoi(p[1])]
		m[aoc.Atoi(p[1])] = mv + 1
	}

	sum := 0
	for _, l := range ll {
		sum += l * m[l]
	}

	fmt.Printf("sum: %d\n", int(sum))
}
