package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

type Bag struct {
	Size     int
	Contains []string
}

var bags = make(map[string]Bag)

func searchBag(name string) int {
	bag := bags[name]
	count := 0
	count += bag.Size

	for _, contains := range bag.Contains {
		parts := strings.Split(contains, " ")
		numContained, _ := strconv.Atoi(parts[0])
		name := strings.Join(parts[1:], " ")
		count += numContained * searchBag(name)
	}

	return count
}

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input = string(file)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " bags contain ")
		if len(parts) < 2 {
			continue
		}
		bagType := parts[0]
		var bag Bag
		contains := strings.Split(parts[1], ", ")
		for _, c := range contains {
			if c == "no other bags." {
				break
			}

			containsType := strings.Join(
				strings.Split(c, " ")[0:3],
				" ")
			size, _ := strconv.Atoi(strings.Split(c, " ")[0])
			bag.Size += size
			bag.Contains = append(bag.Contains, containsType)
		}
		bags[bagType] = bag
	}

	fmt.Println(searchBag("shiny gold"))
}
