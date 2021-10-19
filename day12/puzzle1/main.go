package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

// Directions (NESW being in order at the beginning is important)
const (
	north uint8 = iota
	east
	south
	west
	forward
	turn
)

type Ship struct {
	x      int
	y      int
	facing uint8
}

func (s *Ship) turn(op Op) {
	// iota abuse for no reason
	f := int(s.facing)
	f += (op.value / 90)
	f = (f + 4) % 4
	s.facing = uint8(f)
}

func (s *Ship) Move(op Op) {
	direction := op.direction
	if direction == forward {
		direction = s.facing
	} else if direction == turn {
		s.turn(op)
		return
	}
	switch direction {
	case north:
		s.y += op.value
	case east:
		s.x += op.value
	case south:
		s.y -= op.value
	case west:
		s.x -= op.value
	}
}

// If direction == turn, the move is a turn
type Op struct {
	direction uint8
	value     int
}

func turnOp(s string) (op Op) {
	degrees, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}

	if s[0] == 'L' {
		degrees *= -1
	}
	op.direction = turn
	op.value = degrees
	return op
}

func moveOp(s string) (op Op) {
	direction := s[0]
	switch direction {
	case 'F':
		op.direction = forward
	case 'N':
		op.direction = north
	case 'E':
		op.direction = east
	case 'S':
		op.direction = south
	case 'W':
		op.direction = west
	}
	value, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}
	op.value = value
	return op
}

func NewOp(s string) (op Op) {
	if s[0] == 'R' || s[0] == 'L' {
		return turnOp(s)
	} else {
		return moveOp(s)
	}
}

func Solve(ops []Op) {
	var ship Ship
	// the ship start facing east
	ship.facing = east
	for _, op := range ops {
		ship.Move(op)
	}
	// Get the manhattan distance
	var absX, absY int
	absX = ship.x
	absY = ship.y
	if absX < 0 {
		absX *= -1
	}
	if absY < 0 {
		absY *= -1
	}
	distance := absX + absY
	fmt.Println(distance)
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	var ops []Op

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		ops = append(ops, NewOp(line))
	}
	Solve(ops)
}
