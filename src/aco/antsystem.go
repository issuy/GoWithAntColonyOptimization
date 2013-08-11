package aco

import (
	"fmt"
)

func AntSystem(filename string) {

	fmt.Println("Input city data.")
	cities := CreateCities(filename)
	fmt.Println("Input done.")

	//CheckCities(cities)
	pheromone := InitPheromone(0.09, 0.01, len(cities))
	err := CheckPheromoneInit(pheromone)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
}
