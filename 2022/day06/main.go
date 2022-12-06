package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

const sopLength = 14

func main() {
	data := aoc.ReadStdinRaw()

	for i := 0; i < len(data); i++ {

		sop := data[i : i+sopLength]

		if !hasRepeated(sop) {
			fmt.Println("Found it!, ", i+sopLength, sop)
			break
		}
	}

}

// Checks for repeated characters in a string
func hasRepeated(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return true
			}
		}
	}
	return false
}
