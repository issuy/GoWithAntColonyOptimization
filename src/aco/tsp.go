package aco

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type City struct {
	num int
	x   float64
	y   float64
}

func CreateCities(filename string) (point []City) {
	var cities [100]City

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	lineNumber := 0
	lineReader := bufio.NewReaderSize(file, 20)
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		num := lineNumber
		row := strings.Split(string(line), " ")
		go func() {
			cities[num].num, _ = strconv.Atoi(row[0])
			cities[num].x, _ = strconv.ParseFloat(row[1], 64)
			cities[num].y, _ = strconv.ParseFloat(row[2], 64)
		}()
		lineNumber++
	}

	return cities[:lineNumber]
}

func CheckCities(cities []City) {
	for _, value := range cities {
		fmt.Println(value)
	}
	fmt.Println("Check cities done.")
}
