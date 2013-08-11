package aco

import (
	"bufio"
	"strconv"
	"strings"
)

type City struct {
	num int
	x   float64
	y   float64
}

func CreateCities(lineReader *bufio.Reader) []City {
	var cities [100]City

	lineNumber := 0
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		row := strings.Split(string(line), " ")
		cities[lineNumber].num, _ = strconv.Atoi(row[0])
		cities[lineNumber].x, _ = strconv.ParseFloat(row[1], 64)
		cities[lineNumber].y, _ = strconv.ParseFloat(row[2], 64)
		lineNumber++
	}

	return cities[:lineNumber]
}
