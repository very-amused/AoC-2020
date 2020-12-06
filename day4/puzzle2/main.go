package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var input string

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid"}

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input = string(file)

	fields := make(map[string]*string)
	count := 0
outer:
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			for _, field := range requiredFields {
				if fields[field] == nil {
					fields = make(map[string]*string)
					continue outer
				}
			}

			// Validate
			var p Passport
			p.byr = *fields["byr"]
			p.iyr = *fields["iyr"]
			p.eyr = *fields["eyr"]
			p.hgt = *fields["hgt"]
			p.hcl = *fields["hcl"]
			p.ecl = *fields["ecl"]
			p.pid = *fields["pid"]

			valid, _ := p.validate()
			if valid {
				fmt.Println(p.iyr)
				count++
			} else {
			}

			fields = make(map[string]*string)
			continue outer
		}

		kv := strings.Split(line, " ")
		for _, field := range kv {
			parts := strings.Split(field, ":")
			if len(parts) < 2 {
				continue
			}

			key := parts[0]
			value := parts[1]
			fields[key] = &value
		}
	}

	fmt.Println(count)
}
