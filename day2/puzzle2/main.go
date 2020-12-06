package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

type policy struct {
	pos1   int
	pos2   int
	letter byte
}

func (p policy) verify(password string) bool {
	if len(password) < p.pos1 || len(password) < p.pos2 {
		return false
	}

	match1 := byte(password[p.pos1]) == p.letter
	match2 := byte(password[p.pos2]) == p.letter
	return match1 != match2
}

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input = string(file)

	count := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		var p policy
		policystr := parts[0]
		positions := strings.Split(policystr, " ")[0]
		p.letter = []byte(strings.Split(policystr, " ")[1])[0]

		p.pos1, _ = strconv.Atoi(strings.Split(positions, "-")[0])
		p.pos2, _ = strconv.Atoi(strings.Split(positions, "-")[1])

		if p.verify(parts[1]) {
			count++
		}
	}
	fmt.Println(count)
}
