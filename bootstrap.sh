#! /bin/sh

mkdir -p $1/puzzle1
mkdir -p $1/puzzle2

cat template.go | sed 's/{input}/.\/input.txt/' > $1/puzzle1/main.go
cat template.go | sed 's/{input}/..\/puzzle1\/input.txt/' > $1/puzzle2/main.go
