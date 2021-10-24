package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var input string

type Bus struct {
	// a and n are parameters that fulfill the congruence t % n = a
	a int
	n int
	N int // Ni = N / ni
	x int // Nx % ni = 1
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	input = string(file)

	lines := strings.Split(input, "\n")

	// Extract a and n from each bus, variables used for Chinese Remainder Theory calculations
	var buses []*Bus
	for i, str := range strings.Split(lines[1], ",") {
		if str == "x" {
			continue
		}

		// n is the bus' id
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		// calculate a using the following theorem:
		// if (t + i) % n == 0 and i > 0, then t % n == n - (i % n)
		a := 0
		if i > 0 {
			a = n - (i % n)
		}

		buses = append(buses, &Bus{
			a: a,
			n: n})
	}

	// Calculate N (product of all values of n)
	N := 1
	for _, bus := range buses {
		N *= bus.n
	}

	// Because a = 0 for the first element in the list, it can be predicted that a0 * N0 * x0 == 0
	// This has no effect on the solution, and therefore the first element can now be removed/ignored
	buses = buses[1:]

	T := 0
	// Calculate N and x for each bus
	for _, bus := range buses {
		bus.N = N / bus.n
		base := bus.N % bus.n
		for bus.x = 1; (base*bus.x)%bus.n != 1; bus.x++ {
		}
		T += bus.a * bus.N * bus.x
	}
	t := T % N // t = (T % N) is the solution, satisfying all congruences of t % n = a

	fmt.Println(t)
}
