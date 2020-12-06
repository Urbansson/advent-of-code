package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Urbansson/advent-of-code/2020/pkg/aoc"
)

func main() {
	input := aoc.ReadInput("./input.txt")
	lines := aoc.ExtractLines(input)
	valid := 0
	for _, p := range parse(lines) {
		if p.valid() {
			valid++
		}
	}
	fmt.Println(valid)
}

func parse(lines []string) []passport {
	var ps []passport
	b := strings.Builder{}
	for _, line := range lines {
		if len(line) == 0 {
			ps = append(ps, passwordFromBatchEntry(b.String()))
			b.Reset()
			continue
		}
		b.WriteString(line)
		b.WriteString(" ")
	}
	if b.Len() != 0 {
		ps = append(ps, passwordFromBatchEntry(b.String()))
	}
	return ps
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p passport) valid() bool {
	if !(len(p.byr) > 0 && len(p.iyr) > 0 && len(p.eyr) > 0 && len(p.hgt) > 0 && len(p.hcl) > 0 && len(p.ecl) > 0 && len(p.pid) > 0) {
		return false
	}
	byr := aoc.Atoi(p.byr)
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr := aoc.Atoi(p.iyr)
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr := aoc.Atoi(p.eyr)
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	if strings.HasSuffix(p.hgt, "cm") {
		hgt := aoc.Atoi(strings.TrimSuffix(p.hgt, "cm"))
		if hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.HasSuffix(p.hgt, "in") {
		hgt := aoc.Atoi(strings.TrimSuffix(p.hgt, "in"))
		if hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}
	if ok, _ := regexp.MatchString("^#[0-9a-f]{6}$", p.hcl); !ok {
		return false
	}
	if !(p.ecl == "amb" || p.ecl == "blu" || p.ecl == "brn" || p.ecl == "gry" || p.ecl == "grn" || p.ecl == "hzl" || p.ecl == "oth") {
		return false
	}
	if ok, _ := regexp.MatchString("^[0-9]{9}$", p.pid); !ok {
		return false
	}
	return true
}

func passwordFromBatchEntry(s string) passport {
	fields := strings.Split(strings.TrimSpace(s), " ")
	p := passport{}
	for _, f := range fields {
		kv := strings.Split(f, ":")
		switch kv[0] {
		case "byr":
			p.byr = kv[1]
		case "iyr":
			p.iyr = kv[1]
		case "eyr":
			p.eyr = kv[1]
		case "hgt":
			p.hgt = kv[1]
		case "hcl":
			p.hcl = kv[1]
		case "ecl":
			p.ecl = kv[1]
		case "pid":
			p.pid = kv[1]
		case "cid":
			p.cid = kv[1]
		}
	}
	return p
}
