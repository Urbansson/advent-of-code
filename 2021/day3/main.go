package main

import (
	"fmt"
	"strconv"

	"github.com/Urbansson/advent-of-code/pkg/aoc"
)

func main() {
	data := aoc.ReadStdin()
	lines := aoc.ExtractLines(data)
	mc := make([]int, len(lines[0]))
	for _, l := range lines {
		for i, v := range l {
			if v == '1' {
				mc[i] = mc[i] + 1
			}
		}
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
	// Part 1
	fmt.Println(g * e)
	fmt.Println("----")

	res, _ := findOxygenRating(lines, 0)
	res1, _ := findCO2Rating(lines, 0)
	o2, _ := strconv.ParseInt(res, 2, 64)
	co2, _ := strconv.ParseInt(res1, 2, 64)
	// Part 2
	fmt.Println(o2 * co2)
}

func findMostCommon(input []string, index int) int {
	ones := 0
	zeros := 0
	for _, v := range input {
		if v[index] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	if ones == zeros {
		return 2
	} else if ones > zeros {
		return 1
	}
	return 0
}

func findOxygenRating(input []string, index int) (string, error) {
	if len(input) == 1 {
		return input[0], nil
	}
	if len(input) == 0 {
		return "", fmt.Errorf("failed to find")
	}
	found := []string{}
	mc := findMostCommon(input, index)
	for _, l := range input {
		rating := aoc.Atoi(string(l[index]))
		if mc == 2 && rating == 1 {
			found = append(found, l)
		} else if mc == rating {
			found = append(found, l)
		}
	}
	return findOxygenRating(found, index+1)
}

func findCO2Rating(input []string, index int) (string, error) {
	if len(input) == 1 {
		return input[0], nil
	}
	if len(input) == 0 {
		return "", fmt.Errorf("failed to find")
	}
	found := []string{}
	mc := findMostCommon(input, index)
	for _, l := range input {
		rating := aoc.Atoi(string(l[index]))
		if mc == 2 && rating == 0 {
			found = append(found, l)
		} else if mc != rating && mc != 2 {
			found = append(found, l)
		}
	}
	return findCO2Rating(found, index+1)
}
