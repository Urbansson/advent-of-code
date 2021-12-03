package main

import (
	"fmt"
	"strconv"

	"github.com/Urbansson/advent-of-code/2021/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	mc := make([]int, len(lines[0]))
	for _, l := range lines {
		for i, v := range l {
			print(string(v))
			if v == '1' {
				mc[i] = mc[i] + 1
			}
		}
		println()
	}
	gamma := ""
	epsilon := ""
	for _, v := range mc {
		if v > len(lines)/2 {
			gamma = fmt.Sprintf("%s%d", gamma, 1)
			epsilon = fmt.Sprintf("%s%d", epsilon, 0)
		} else {
			gamma = fmt.Sprintf("%s%d", gamma, 0)
			epsilon = fmt.Sprintf("%s%d", epsilon, 1)
		}
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(g * e)
}
