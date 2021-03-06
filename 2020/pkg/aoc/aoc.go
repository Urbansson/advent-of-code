package aoc

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}

func ReadStdin() string {
	d, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(d))
}

// IntList parses a list of ints. devided by spaces or newlines
func IntList(s string) []int {
	fs := strings.Fields(s)
	ints := make([]int, len(fs))
	for i, n := range fs {
		ints[i] = Atoi(n)
	}
	return ints
}

func ExtractInts(s string) []int {
	fs := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '-'
	})
	ints := make([]int, 0, len(fs))
	for _, w := range fs {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

// Atoi is a passthrough for strconv.Atoi that panics upon failure.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// ExtractLines splits input on newlines
func ExtractLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

// Digits splits a int to its digits
func Digits(i int) []int {
	var digits []int
	for i != 0 {
		digits = append([]int{i % 10}, digits...)
		i /= 10
	}
	return digits
}

// MinMax returns min max values from a slice.
func MinMax(array []int) (int, int) {
	max := array[0]
	min := array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
