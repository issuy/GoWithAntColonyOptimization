package aco

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestCities(t *testing.T) {
	file, err := os.Open("../att48.tsp")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	lineReader := bufio.NewReaderSize(file, 20)

	cities := CreateCities(lineReader)

	for i := range cities {
		if i != cities[i].num-1 {
			t.Errorf("cities[%d].num = %2d, want %2d.", i, cities[i].num, i+1)
		}
	}
}
