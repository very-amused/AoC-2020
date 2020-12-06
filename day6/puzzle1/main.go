package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var input string

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	group := make(map[rune]bool)
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			count := 0
			for char := 'a'; char <= 'z'; char++ {
				if group[char] {
					count++
				}
			}
			sum += count
			// reset map
			group = make(map[rune]bool)
		}

		for _, char := range line {
			group[char] = true
		}
	}

	fmt.Println(sum)
}
