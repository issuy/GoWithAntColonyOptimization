package aco

import(
    "fmt"
    "errors"
)

type Pheromone struct {
     init   float64
     dtau   float64
     length int
     tau    [][]float64
}

func InitPheromone(init float64, dtau float64, length int) (pheromone Pheromone) {
     pheromone.init   = init
     pheromone.dtau   = dtau
     pheromone.length = length

     tau := make([][]float64, length)
     for key1, _ := range tau {
         tmp := make([]float64, length)
         go func() {
            for key2, _ := range tmp {
                tmp[key2] = init
            }
         } ()
         tau[key1] = tmp
     }
     pheromone.tau = tau
     return
}

func CheckPheromoneInit(pheromone Pheromone) (error){
     fmt.Printf("[Check pheromone start] init:%f dtau:%f length:%d\n", pheromone.init, pheromone.dtau, pheromone.length)

     if(len(pheromone.tau) == pheromone.length) {
         for _, value1 := range pheromone.tau {
             if len(value1) != pheromone.length {
                 fmt.Println("[Check pheromone end] Init is failed.")
                 return errors.New("Invalid argument")
             }
             for _, value2 := range value1 {
                 if value2 != pheromone.init {
                    fmt.Println("[Check pheromone end] Init is failed.")
                    return errors.New("Invalid argument")
                 }
             }
         }
     } else {
       fmt.Println("[Check pheromone end] Init is failed.")
       return errors.New("Invalid argument")
     }
     fmt.Println("[Check pheromone end] Init is succeeded.")
     return nil
}