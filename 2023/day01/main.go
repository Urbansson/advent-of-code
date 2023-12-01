package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	// fix edgecase of oneight == 18, not 1ight.
	r := strings.NewReplacer(
		"one", "1e",
		"two", "2o",
		"three", "3e",
		"four", "4r",
		"five", "5e",
		"six", "6x",
		"seven", "7n",
		"eight", "8t",
		"nine", "9e",
	)
	for _, l := range lines {
		// Run replacer twice to fix the updated values.
		nl := r.Replace(r.Replace(l))
		first := 0
		last := 0
		for _, c := range nl {
			if ok := unicode.IsDigit(c); ok {
				if first == 0 {
					first = int(c-'0') * 10
				}
				last = int(c - '0')
			}
		}
		sum += first + last
	}

	fmt.Println(sum)
}
