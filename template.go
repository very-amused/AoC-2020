package main

import (
	"io/ioutil"
	"strings"
)

var input string

func main() {
	file, _ := ioutil.ReadFile("{input}")
	input = string(file)

	for _, line := range strings.Split(input, "\n") {

	}
}
