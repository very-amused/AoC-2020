package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	lines := strings.Split(input, "\n")
	minTime, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}

	// Get a list of bus IDs
	var buses []int
	for _, str := range strings.Split(lines[1], ",") {
		if str == "x" {
			continue
		}
		id, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		buses = append(buses, id)
	}
	// Figure out which bus to take
	busID := buses[0]
	waitTime := (((minTime / buses[0]) + 1) * buses[0]) % minTime
	for _, bus := range buses {
		w := (((minTime / bus) + 1) * bus) % minTime
		if w < waitTime {
			waitTime = w
			busID = bus
		}
	}
	fmt.Println(busID * waitTime)
}
