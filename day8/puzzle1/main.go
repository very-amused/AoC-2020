package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

var completed = make(map[int]bool)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	accumulator := 0
	for i := 0; i < len(strings.Split(input, "\n")); i++ {
		line := strings.Split(input, "\n")[i]
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		if completed[i] {
			fmt.Println(accumulator)
			break
		}
		completed[i] = true

		switch parts[0] {
		case "nop":
			continue
		case "acc":
			inst := parts[1]
			op := inst[0]
			num, _ := strconv.Atoi(inst[1:])
			if op == '+' {
				accumulator += num
			} else if op == '-' {
				accumulator -= num
			}
			break
		case "jmp":
			inst := parts[1]
			op := inst[0]
			num, _ := strconv.Atoi(inst[1:])
			if op == '+' {
				i += num - 1
			} else if op == '-' {
				i -= num + 1
			}
			break
		}
	}
}
