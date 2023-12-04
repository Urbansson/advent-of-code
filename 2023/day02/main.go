package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	re := regexp.MustCompile(`(\d+) (\w+)`)

	sum := 0
	power := 0

	for i, l := range lines {
		r := strings.Split(l, ":")[1]
		lu := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(r, -1) {
			n, _ := strconv.Atoi(m[1])
			lu[m[2]] = aoc.Max(n, lu[m[2]])
		}

		if lu["red"] <= 12 && lu["green"] <= 13 && lu["blue"] <= 14 {
			sum += i + 1
		}
		power += lu["red"] * lu["green"] * lu["blue"]
	}
	fmt.Println(sum, power)
}
