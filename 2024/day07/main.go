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
		r := strings.Split(l, ":")
		res := aoc.Atoi(r[0])
		nums := aoc.ExtractInts(r[1])

		if calc(res, nums[0], nums[1:], fmt.Sprintf("%d", nums[0])) {
			sum += res
		}
	}

	fmt.Println(sum)
}

func calc(res, cur int, i []int, f string) bool {
	if len(i) == 0 {
		return cur == res
	}

	next := i[0]
	sum := false

	if calc(res, cur+next, i[1:], fmt.Sprintf("%s+%d", f, next)) {
		sum = true
	}

	if calc(res, cur*next, i[1:], fmt.Sprintf("%s*%d", f, next)) {
		sum = true
	}

	result := aoc.Atoi(fmt.Sprintf("%d%d", cur, next))
	if calc(res, result, i[1:], fmt.Sprintf("%s*%d", f, next)) {
		sum = true
	}

	return sum
}
