package main

import(
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type City struct {
     num int
     x   float64
     y   float64
}

func main() {
     fmt.Println("Start Program.")
     
     fmt.Println("Input root data.")
     var cities [100]City
     file, err := os.Open("att48.tsp")
     if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        return
     }
     
     lineNumber := 0
     lineReader := bufio.NewReaderSize(file, 20)
     for line, _, e := lineReader.ReadLine(); e==nil; line, _, e = lineReader.ReadLine() {
         row := strings.Split(string(line), " ")
         go func () {
                  cities[lineNumber].num, _ = strconv.Atoi(row[0])
                  cities[lineNumber].x, _   = strconv.ParseFloat(row[1], 64)
                  cities[lineNumber].y, _   = strconv.ParseFloat(row[2], 64)
                  if lineNumber<10 {
                     fmt.Printf("%.3d: %3d %f %f\n", lineNumber, cities[lineNumber].num, cities[lineNumber].x, cities[lineNumber].y)
                  } else {
                       fmt.Printf("%.3d, ", lineNumber)
                  }
         } ()
         lineNumber++
     }
     fmt.Println("...completed\ninput done.")

    //fmt.Println("Do ACO")

    //fmt.Println("Resutls")

    fmt.Println("End Program.")
   
}
