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

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	passport := make(map[string]*string)
	count := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			for _, field := range requiredFields {
				if passport[field] == nil {
					goto end
				}
			}
			count++
		end:
			passport = make(map[string]*string)
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if len(parts) < 2 {
				continue
			}

			key := parts[0]
			value := parts[1]
			passport[key] = &value
		}
	}

	fmt.Println(count)
}
