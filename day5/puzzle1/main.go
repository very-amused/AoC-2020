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

	highestID := 0
	for _, pass := range strings.Split(input, "\n") {
		if len(pass) != 10 {
			continue
		}
		min := 0
		max := 127
		rowinfo := pass[0:7]
		for _, char := range rowinfo {
			half := ((max - min) + 1) / 2
			if char == 'F' {
				max -= half
			} else if char == 'B' {
				min += half
			}
		}
		// Sanity check
		if min != max {
			fmt.Println("ERROR in BSP:")
			fmt.Println(min, max)
		}
		row := min

		min = 0
		max = 7
		columninfo := pass[7:]
		for _, char := range columninfo {
			half := ((max - min) + 1) / 2
			if char == 'L' {
				max -= half
			} else if char == 'R' {
				min += half
			}
		}
		// Sanity check
		if min != max {
			fmt.Println("ERROR in BSP:")
			fmt.Println(min, max)
		}
		column := min

		id := (row * 8) + column
		if id > highestID {
			highestID = id
		}
	}

	fmt.Println(highestID)
}
