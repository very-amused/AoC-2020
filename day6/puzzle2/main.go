package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var input string

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input = string(file)

	// Map to number of ppl who answered question
	group := make(map[rune]int)
	sum := 0
	people := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			count := 0
			for char := 'a'; char <= 'z'; char++ {
				if group[char] == people {
					count++
				}
			}
			sum += count
			// reset map + headcount
			group = make(map[rune]int)
			people = 0
			continue
		}

		for _, char := range line {
			group[char]++
		}
		// Add person to group count
		people++
	}

	fmt.Println(sum)
}
