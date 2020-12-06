package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Map to coordinate grid
var grid = make(map[string]int)
var xmax int
var ymax int

const tree = '#'

var input string

func sled(right, down int) (count int) {
	x := 0
	for i := 0; i < len(strings.Split(input, "\n")); i += down {
		line := strings.Split(input, "\n")[i]
		if len(line) == 0 {
			continue
		}

		if line[x] == tree {
			count++
		}
		x += right
		if x > (len(line) - 1) {
			x = x - len(line)
		}
	}

	return count
}

func main() {
	file, err := ioutil.ReadFile("../puzzle1/input.txt")
	if err != nil {
		panic(err)
	}
	input = string(file)
	answer := sled(1, 1) * sled(3, 1) * sled(5, 1) * sled(7, 1) * sled(1, 2)

	fmt.Println(answer)
}
