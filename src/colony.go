package main

import (
	"aco"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start program.")

	file, err := os.Open("att48.tsp")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	lineReader := bufio.NewReaderSize(file, 20)

	fmt.Println("Do ACO")
	aco.AntSystem(lineReader)

	//fmt.Println("Resutls")

	fmt.Println("End program.")
}
