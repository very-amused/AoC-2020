package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var combinations = make(map[int]int)

func solve(chain []int) (c int) {
	combinations[0] = 1
	for i, adapter := range chain {
		combinations[adapter] = combinations[adapter-1] + combinations[adapter-2] + combinations[adapter-3]
		if i == len(chain)-1 {
			c = combinations[adapter]
			break
		}
	}
	return c
}

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input := string(file)

	var adapters []int
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		adapters = append(adapters, num)
	}

	sort.Slice(adapters, func(i, j int) bool {
		return adapters[i] < adapters[j]
	})

	fmt.Println(solve(adapters))
}
