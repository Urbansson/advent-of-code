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

	// Contain what pages needs to be printed after the keyed page.
	after := map[int][]int{}

	// COntains what pages needs to be printed before the keyed page.
	before := map[int][]int{}

	rules := []rule{}

	pi := 0

	for i, l := range lines {
		if l == "" {
			pi = i + 1
			break
		}
		s := strings.Split(l, "|")
		b := aoc.Atoi(s[0])
		a := aoc.Atoi(s[1])
		after[b] = append(after[b], a)
		before[a] = append(before[a], b)

		rules = append(rules, rule{start: b, end: a})

	}

	updates := [][]int{}
	for _, l := range lines[pi:] {
		if l == "" {

			break
		}
		updates = append(updates, aoc.ExtractInts(l))

	}

	for k, v := range after {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("---")
	for k, v := range before {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("---")
	for _, v := range updates {
		fmt.Println(v)
	}
	fmt.Println("---")

	sum := 0

	for _, update := range updates {
		valid := true
		fmt.Println("Checking if update is valid,", update)

		for _, r := range rules {

			fmt.Println("check rule", r)

			si := contains(update, r.start)

			ei := contains(update, r.end)

			if si < 0 {
				continue
			}

			if ei < 0 {
				continue
			}

			fmt.Println(si, ei)
			if si > ei {
				fmt.Println("invalid")

				valid = false
			}
		}

		fmt.Println(valid, update)
		if valid {

			// fmt.Println("Middle valie", u[len(u)/2])
			sum += update[len(update)/2]
		}
		fmt.Println("---")

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
