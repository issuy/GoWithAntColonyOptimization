package aco

import (
	"bufio"
	"fmt"
)

func AntSystem(lineReader *bufio.Reader) {

	fmt.Print("Input city data.")
	cities := CreateCities(lineReader)
	fmt.Println("..done")

	_ = NewPheromone(0.09, 0.01, len(cities))

}
