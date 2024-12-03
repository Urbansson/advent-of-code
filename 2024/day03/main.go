package main

import (
	"fmt"
	"regexp"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	ri := regexp.MustCompile(`\d{1,3}`)
	sum := 0
	for _, l := range r.FindAllString(data, -1) {
		i := ri.FindAllString(l, -1)
		v1 := aoc.Atoi(i[0])
		v2 := aoc.Atoi(i[1])
		sum += v1 * v2
	}
	fmt.Println(sum)
}
