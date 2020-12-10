package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input := string(file)

	var nums []int
	var invalidNum int

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)

		if len(nums) > 25 {
			previous := nums[len(nums)-26 : len(nums)-1]
			valid := false
			for _, i := range previous {
				for _, j := range previous {
					if i+j == num {
						valid = true
					}
				}
			}
			if !valid {
				invalidNum = num
			}
		}
	}

	// Bruteforce all possible combination lengths of numbers
	for combLen := 2; combLen <= len(nums); combLen++ {
		for i := 0; i+combLen <= len(nums); i++ {
			combination := nums[i : i+combLen]
			sum := 0
			for _, j := range combination {
				sum += j
			}

			if sum == invalidNum {
				// Calculate minimum and maximum number
				min := combination[0]
				max := combination[0]
				for _, num := range combination {
					if num < min {
						min = num
					} else if num > max {
						max = num
					}
				}

				fmt.Println(min + max)
			}
		}
	}
}
