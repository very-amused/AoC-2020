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

func main() {
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	x := 0
	count := 0
	for i := 0; i < len(strings.Split(input, "\n")); i++ {
		line := strings.Split(input, "\n")[i]
		if len(line) == 0 {
			continue
		}

		row := []rune(line)
		if line[x] == tree {
			count++
			row[x] = 'H'
		} else {
			row[x] = 'h'
		}
		x += 3
		if x > (len(line) - 1) {
			x = x - len(line)
		}
		fmt.Println(string(row))
	}
	fmt.Println(count)
}
