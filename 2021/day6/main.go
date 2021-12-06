package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	dig := aoc.ExtractInts(lines[0])

	fishes := make([]lanternfish, len(dig))
	for i, a := range dig {
		fishes[i] = lanternfish{
			age: a,
		}
	}

	for i := 0; i < 80; i++ {
		newFishes := []lanternfish{}
		for i := range fishes {
			fishes[i].age = fishes[i].age - 1
			if fishes[i].age < 0 {
				newFishes = append(newFishes, lanternfish{
					age: 8,
				})
				fishes[i].age = 6
			}
		}
		fishes = append(fishes, newFishes...)
	}
	fmt.Println(len(fishes))
}

type lanternfish struct {
	age int
}
