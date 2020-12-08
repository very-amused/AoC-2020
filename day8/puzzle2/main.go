package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Instruction - A single instruction for the assembler
type Instruction struct {
	Op     string
	Symbol rune
	Value  int
}

func execute(instructions []Instruction) (terminates bool, accumulator int) {
	done := make(map[int]bool)
	i := 0
	for {
		if i >= len(instructions) {
			return true, accumulator
		} else if done[i] {
			return false, accumulator
		}
		done[i] = true

		ins := instructions[i]

		switch ins.Op {
		case "acc":
			if ins.Symbol == '+' {
				accumulator += ins.Value
			} else if ins.Symbol == '-' {
				accumulator -= ins.Value
			}
			i++
			break
		case "jmp":
			if ins.Symbol == '+' {
				i += ins.Value
			} else if ins.Symbol == '-' {
				i -= ins.Value
			}
			break
		default:
			i++
		}

	}
}

func lexer(input string) (instructions []Instruction) {
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		tokens := strings.Split(line, " ")
		var ins Instruction
		ins.Op = tokens[0]
		ins.Symbol = rune(tokens[1][0])
		var err error
		ins.Value, err = strconv.Atoi(tokens[1][1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, ins)
	}

	return instructions
}

func main() {
	file, _ := ioutil.ReadFile("../puzzle1/input.txt")
	input := string(file)

	instructions := lexer(input)

	// Try bruteforcing instructions
	for i, ins := range instructions {
		cpy := make([]Instruction, len(instructions))
		copy(cpy, instructions)
		if ins.Op == "jmp" {
			cpy[i] = Instruction{
				Op:     "noop",
				Symbol: ins.Symbol,
				Value:  ins.Value}
		} else if ins.Op == "noop" {
			cpy[i] = Instruction{
				Op:     "jmp",
				Symbol: ins.Symbol,
				Value:  ins.Value}
		}

		terminates, accumulator := execute(cpy)
		if terminates {
			fmt.Println(accumulator)
		}
	}
}
