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

	// Map of all taken seats
	var seats = make(map[int]bool)

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

		id := (8 * row) + column
		seats[id] = true
	}

	for row := 0; row <= 127; row++ {
		for column := 0; column <= 7; column++ {
			id := (8 * row) + column
			isFilled := seats[id]
			lastSeat := seats[id-1]
			nextSeat := seats[id+1]
			if !isFilled && lastSeat && nextSeat {
				fmt.Println(id)
			}
		}
	}
}
