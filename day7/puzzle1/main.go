package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var input string

type Bag struct {
	Contains []string
}

var bags = make(map[string]Bag)

func searchBag(name string) bool {
	bag := bags[name]
	for _, contains := range bag.Contains {
		if contains == "shiny gold" {
			return true
		}
		if searchBag(contains) {
			return true
		}
	}

	return false
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
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
				strings.Split(c, " ")[1:3],
				" ")
			bag.Contains = append(bag.Contains, containsType)
		}
		bags[bagType] = bag
	}

	count := 0

	for name := range bags {
		if searchBag(name) {
			count++
		}
	}

	fmt.Println(count)
}
