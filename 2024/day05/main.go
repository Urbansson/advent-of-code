package main

import (
	"fmt"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type rule struct {
	start int
	end   int
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	rules := []rule{}

	pi := 0
	for i, l := range lines {
		if l == "" {
			pi = i + 1
			break
		}
		s := strings.Split(l, "|")
		rules = append(rules, rule{start: aoc.Atoi(s[0]), end: aoc.Atoi(s[1])})
	}

	updates := [][]int{}
	for _, l := range lines[pi:] {
		updates = append(updates, aoc.ExtractInts(l))
	}

	sum := 0

	for _, update := range updates {
		valid := false
		for i := 0; i < len(rules); i++ {
			r := rules[i]
			si := contains(update, r.start)
			ei := contains(update, r.end)

			if si < 0 {
				continue
			}
			if ei < 0 {
				continue
			}

			if si > ei {
				update[si], update[ei] = update[ei], update[si]
				valid = true
				i = 0
			}
		}

		if valid {
			sum += update[len(update)/2]
		}
	}
	fmt.Println(sum)
}

func contains(s []int, e int) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}
