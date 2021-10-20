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

// #region Ship

type Ship struct {
	x        int
	y        int
	waypoint *Waypoint
}

type Waypoint struct {
	x int
	y int
}

func (s *Ship) turn(op Op) {
	// Get the movement as an absolute positive degree
	degrees := (op.value + 360) % 360
	// yes
	w := s.waypoint
	for d := degrees; d > 0; d -= 90 {
		x := w.x
		y := w.y
		w.x = y
		w.y = -x
	}
}

func (s *Ship) Move(op Op) {
	direction := op.direction

	w := s.waypoint
	v := op.value
	switch direction {
	case forward:
		// Move the ship forward to the waypoint
		s.x += v * w.x
		s.y += v * w.y
	case turn:
		s.turn(op)
	// Move the waypoint
	case north:
		w.y += v
	case east:
		w.x += v
	case south:
		w.y -= v
	case west:
		w.x -= v
	}
}

// #endregion

// #region Op

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

// #endregion

func Solve(ops []Op) {
	var ship Ship
	// The waypoint starts 10 units east and 1 unit north from the ship
	ship.waypoint = &Waypoint{
		x: 10,
		y: 1}
	// the ship start facing east
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
