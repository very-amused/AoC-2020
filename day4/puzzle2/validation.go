package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var hclRegex = regexp.MustCompile("^#([0-9]|[a-f]){6}$")
var eclRegex = regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
var pidRegex = regexp.MustCompile("^\\d{9}$")

func (p Passport) validate() (bool, string) {
	if len(p.byr) != 4 {
		return false, ""
	}
	byr, err := strconv.Atoi(p.byr)
	if err != nil || !(byr >= 1920 && byr <= 2002) {
		return false, fmt.Sprintf("byr:%s", p.byr)
	}

	if len(p.iyr) != 4 {
		return false, ""
	}
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil || !(iyr >= 2010 && byr <= 2020) {
		return false, fmt.Sprintf("iyr:%s", p.iyr)
	}

	if len(p.eyr) != 4 {
		return false, ""
	}
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil || !(eyr >= 2020 && eyr <= 2030) {
		return false, fmt.Sprintf("eyr:%s", p.eyr)
	}

	hgtSuffix := p.hgt[len(p.hgt)-2:]
	hgt, err := strconv.Atoi(p.hgt[:len(p.hgt)-2])
	if err != nil {
		return false, fmt.Sprintf("hgt:%s", p.hgt)
	}
	if hgtSuffix == "cm" {
		if !(hgt >= 150 && hgt <= 193) {
			return false, fmt.Sprintf("hgt:%s", p.hgt)
		}
	} else if hgtSuffix == "in" {
		if !(hgt >= 59 && hgt <= 76) {
			return false, fmt.Sprintf("hgt:%s", p.hgt)
		}
	} else {
		return false, ""
	}

	if !hclRegex.Match([]byte(p.hcl)) {
		return false, fmt.Sprintf("hcl:%s", p.hcl)
	}

	if !eclRegex.Match([]byte(p.ecl)) {
		return false, fmt.Sprintf("ecl:%s", p.ecl)
	}

	if !pidRegex.Match([]byte(p.pid)) {
		return false, fmt.Sprintf("pid:%s", p.pid)
	}

	return true, fmt.Sprintf("pid:%s", p.pid)
}
