package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

type rangePoint struct {
	isStart bool
	value   int
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	ranges := []rangePoint{}
	for _, line := range lines {
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		s, e := parts[0], parts[1]

		ranges = append(ranges, rangePoint{
			isStart: true,
			value:   aoc.Atoi(s),
		})
		ranges = append(ranges, rangePoint{
			isStart: false,
			value:   aoc.Atoi(e),
		})
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].value != ranges[j].value {
			return ranges[i].value < ranges[j].value
		}
		// When values are equal, starts come before ends
		return ranges[i].isStart && !ranges[j].isStart
	})

	sum := 0
	stack := []int{}
	for _, point := range ranges {
		if point.isStart {
			stack = append(stack, point.value)
		} else {
			if len(stack) == 1 {
				sum += point.value - stack[0] + 1
			}
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println(sum)
}
