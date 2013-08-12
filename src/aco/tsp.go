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
	cities := make([]City, 0, 0)

	lineNumber := 0
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		row := strings.Split(string(line), " ")
		var city City
		city.num, _ = strconv.Atoi(row[0])
		city.x, _ = strconv.ParseFloat(row[1], 64)
		city.y, _ = strconv.ParseFloat(row[2], 64)
		cities = append(cities, city)
		lineNumber++
	}

	return cities
}
