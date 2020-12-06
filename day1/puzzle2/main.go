package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func loadInput() (nums []int) {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	strnums := strings.Split(string(file), "\n")
	for i, str := range strnums {
		if i == len(strnums)-1 {
			continue
		}
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func main() {
	nums := loadInput()
outer:
	for _, i := range nums {
		for _, j := range nums {
			for _, l := range nums {
				if i+j+l == 2020 {
					fmt.Println(i * j * l)
					break outer
				}
			}
		}
	}
}
