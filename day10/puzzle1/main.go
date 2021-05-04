package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var adapters []int

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input := string(file)

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

	voltage := 0
	oneVDiff := 0
	threeVDiff := 0
	for _, adapter := range adapters {
		if adapter-voltage == 1 {
			oneVDiff++
		} else if adapter-voltage == 3 {
			threeVDiff++
		}
		voltage = adapter
	}
	// Add the final constant diff of 3
	threeVDiff++

	fmt.Println(oneVDiff * threeVDiff)
}
