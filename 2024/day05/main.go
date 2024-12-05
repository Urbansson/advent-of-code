package main

import (
	"fmt"
	"strings"

	"slices"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	// Contain what pages needs to be printed after the keyed page.
	after := map[int][]int{}

	// COntains what pages needs to be printed before the keyed page.
	before := map[int][]int{}

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
	}

	updates := [][]int{}
	for _, l := range lines[pi:] {
		if l == "" {

			break
		}
		updates = append(updates, aoc.ExtractInts(l))

	}

	// for k, v := range after {
	// 	fmt.Println(k, ": ", v)
	// }
	// fmt.Println("---")
	// for k, v := range before {
	// 	fmt.Println(k, ": ", v)
	// }
	// fmt.Println("---")
	// for _, v := range updates {
	// 	fmt.Println(v)
	// }
	// fmt.Println("---")

	sum := 0
	for _, u := range updates {
		// fmt.Println("Checking if update is valid,", u)
		seenBefore := map[int]bool{}

		valid := true

		for i, update := range u {

			if v, ok := after[update]; ok {

				// fmt.Println("Checking that ", after[update], " is not present before", update, "in", u[:i])
				// If we have seen the update that should be after this it is invalid
				for _, shouldBeAfter := range v {

					if _, ok := seenBefore[shouldBeAfter]; ok {
						// Invalid
						// fmt.Println("Invalid,", shouldBeAfter, " should not be before", update)
						valid = false
					}
				}

			}

			// fmt.Println("----")

			if v, ok := before[update]; ok {
				// fmt.Println("Checking that ", before[update], " is not present in", u[i+1:])

				for _, shouldNotBeAfter := range v {
					if ok := slices.Contains(u[i+1:], shouldNotBeAfter); ok {
						// fmt.Println("Invalid,", shouldNotBeAfter, " should not be after", update)
						valid = false
					}
				}
			}
			seenBefore[update] = true

			// fmt.Println("Checking that ", before[update], " is not present after", update, "in", u[i+1:])

			// if v, ok := before[update]; ok {
			// 	// If we have not seen the update that should be before this it is invalid
			// 	for _, shouldBeBefore := range v {
			// 		if _, ok := seen[shouldBeBefore]; !ok {
			// 			// Invalid
			// 			fmt.Println("Invalid, should have been seen")
			// 			valid = false

			// 		}
			// 	}
			// }
			// fmt.Println("----")

		}

		if valid {

			// fmt.Println("Middle valie", u[len(u)/2])
			sum += u[len(u)/2]
		}

		// fmt.Println("---", valid, ":", u)
		// fmt.Println("---")
	}

	fmt.Println("result:", sum)
}
