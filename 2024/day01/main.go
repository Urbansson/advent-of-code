package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	ll := make([]int, 0, len(lines))
	lr := make([]int, 0, len(lines))

	// Parse
	for _, l := range lines {

		p := strings.Split(l, "   ")

		ll = append(ll, aoc.Atoi(p[0]))
		lr = append(lr, aoc.Atoi(p[1]))

	}

	sort.Ints(ll)
	sort.Ints(lr)
	// Compare.

	sum := float64(0)
	for i, l := range ll {

		sum += math.Abs(float64(l - lr[i]))

	}

	fmt.Printf("sum: %d\n", int(sum))
}
