package utils

import (
	"bytes"
	"io/ioutil"
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
