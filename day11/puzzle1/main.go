package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	empty uint8 = iota
	occupied
	floor
)

type Seat struct {
	x int
	y int

	state uint8
}

// [row][column]
type Cabin map[int]map[int]*Seat

func (c Cabin) clone() (new Cabin) {
	new = make(Cabin)

	for i, row := range c {
		new[i] = make(map[int]*Seat)
		for j, cell := range row {
			new[i][j] = cell
		}
	}

	return new
}

// Return the list of seats adjascent to a given seat
func (c Cabin) adjust(cell Seat) (new Seat) {
	new = cell
	var adj []*Seat

	for y := cell.y - 1; y <= cell.y+1; y += 2 {
		for x := cell.x - 1; x <= cell.x+1; x++ {
			adj = append(adj, c[y][x])
		}
	}
	adj = append(adj, c[cell.y][cell.x-1])
	adj = append(adj, c[cell.y][cell.x+1])

	numOccupied := 0
	for _, a := range adj {
		if a != nil && a.state == occupied {
			numOccupied++
		}
	}
	if cell.state == empty && numOccupied == 0 {
		new.state = occupied
	}
	if cell.state == occupied && numOccupied >= 4 {
		new.state = empty
	}

	return new
}

func solve(c Cabin) (numOccupied int) {
	// Clone the cabin so all changes can be applied without interdependence
	old := c.clone()
	new := c.clone()
	stable := false

	for !stable {
		stable = true

		for y, row := range old {
			for x, cell := range row {
				if cell == nil || cell.state == floor {
					continue
				}

				adjusted := old.adjust(*cell)
				if adjusted.state != cell.state {
					stable = false
				}
				new[y][x] = &adjusted
			}
		}
		old = new.clone()
	}

	for _, row := range old {
		for _, cell := range row {
			if cell.state == occupied {
				numOccupied++
			}
		}
	}

	return numOccupied
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input := string(file)

	lines := strings.Split(input, "\n")
	var seats Cabin = make(Cabin)

	for i, line := range lines {
		if len(line) == 0 {
			break
		}

		row := make(map[int]*Seat)

		for j, char := range line {
			var seat Seat
			seat.y = i
			seat.x = j
			switch char {
			case 'L':
				seat.state = empty
				break

			case '#':
				seat.state = occupied
				break

			case '.':
				seat.state = floor
				break

			default:
				panic(fmt.Sprint("unexpected cell character", char))
			}

			row[j] = &seat
		}
		seats[i] = row
	}

	fmt.Println(solve(seats))
}
