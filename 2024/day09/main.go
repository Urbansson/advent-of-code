package main

import (
	"fmt"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()

	d := aoc.DigitList(data)

	disk := []int{}
	id := 0
	for i, v := range d {
		free := i%2 != 0
		for range v {
			if !free {
				disk = append(disk, id)
			} else {
				disk = append(disk, -1)
			}
		}
		if !free {
			id++
		}
	}
	i := defrag(disk)
	cs := 0
	for i, v := range disk[:i] {
		cs += i * v
	}
	fmt.Println(cs)
}

func defrag(disk []int) int {
	lastIndex := len(disk) - 1
	for i := 0; i < lastIndex; i++ {
		if disk[i] == -1 {
			disk[i] = disk[lastIndex]
			disk[lastIndex] = -1
			lastIndex = findLastId(disk, lastIndex)
		}
	}
	return lastIndex + 1
}

func findLastId(disk []int, prev int) int {
	for i := prev; i >= 0; i-- {
		if disk[i] != -1 {
			return i
		}
	}
	return -1
}
