package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type policy struct {
	min    int
	max    int
	letter byte
}

func (p policy) verify(password string) bool {
	count := 0
	for _, char := range password {
		if byte(char) == p.letter {
			count++
		}
		if count > p.max {
			return false
		}
	}
	return count >= p.min
}

var input string

func getPolicies() (policies []policy) {
	for _, line := range strings.Split(input, "\n") {
		policystr := strings.Split(line, ":")[0]

		if len(strings.Split(policystr, " ")) < 2 {
			continue
		}

		var p policy
		minmax := strings.Split(policystr, " ")[0]
		p.letter = []byte(strings.Split(policystr, " ")[1])[0]

		min := strings.Split(minmax, "-")[0]
		max := strings.Split(minmax, "-")[1]

		p.min, _ = strconv.Atoi(min)
		p.max, _ = strconv.Atoi(max)
		policies = append(policies, p)
	}
	return policies
}

func getPasswords() (passwords []string) {
	for _, line := range strings.Split(input, "\n") {
		if len(strings.Split(line, ":")) < 2 {
			continue
		}
		password := strings.Trim(strings.Split(line, ":")[1], " ")
		passwords = append(passwords, password)
	}
	return passwords
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)
	policies := getPolicies()
	passwords := getPasswords()
	count := 0
	for i, password := range passwords {
		if policies[i].verify(password) {
			count++
		}
	}
	fmt.Println(count)
}
