package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

// A X ROCK
// B Y PAPER
// C Z SCISSORS

var lookup = map[string]int{
	"A X": 3 + 1, // DRAW
	"A Y": 6 + 2, // WIN
	"A Z": 0 + 3, // LOSE
	"B X": 0 + 1, // LOSE
	"B Y": 3 + 2, // DRAW
	"B Z": 6 + 3, // WIN
	"C X": 6 + 1, // WIN
	"C Y": 0 + 2, // LOSE
	"C Z": 3 + 3, // DRAW
}

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)

	sum := 0
	for _, l := range lines {
		sum += lookup[l]
	}
	fmt.Println(sum)
}
