package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input := string(file)

	var nums []int

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
		if len(nums) > 26 {
			// Trim old numbers
			nums = nums[len(nums)-26:]
		}

		if len(nums) >= 26 {
			previous := nums[:len(nums)-1]
			valid := false
			for _, i := range previous {
				for _, j := range previous {
					if i+j == num {
						valid = true
					}
				}
			}
			if !valid {
				fmt.Println(num)
			}
		}
	}
}
