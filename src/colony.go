package main

import (
	"aco"
	"fmt"
)

func main() {
	fmt.Println("Start program.")

	fmt.Println("Do ACO")
	aco.AntSystem("att48.tsp")

	//fmt.Println("Resutls")

	fmt.Println("End program.")
}
