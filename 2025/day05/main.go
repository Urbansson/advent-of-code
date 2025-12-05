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

	dividerIndex := 0
	for i, line := range lines {
		if line == "" {
			dividerIndex = i
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
		return ranges[i].value < ranges[j].value
	})

	fresh := 0
	for _, line := range lines[dividerIndex+1:] {
		checkValue := aoc.Atoi(line)

		ar := 0
		for _, point := range ranges {
			if point.value >= checkValue {
				break
			}

			if point.isStart {
				ar++
			} else {
				ar--
			}
		}

		if ar > 0 {
			fresh++
		}
	}
	fmt.Println(fresh)
}
