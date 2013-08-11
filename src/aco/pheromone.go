package aco

type Pheromone struct {
	init   float64
	dtau   float64
	length int
	tau    [][]float64
}

func NewPheromone(init float64, dtau float64, length int) *Pheromone {
	tau := make([][]float64, length)
	for i := range tau {
		tau[i] = make([]float64, length)
		for j := range tau {
			tau[i][j] = init
		}
	}

	return &Pheromone{init, dtau, length, tau}
}

func (p *Pheromone) Init() float64 {
	return p.init
}

func (p *Pheromone) Dtau() float64 {
	return p.dtau
}

func (p *Pheromone) Length() int {
	return p.length
}

func (p *Pheromone) Tau() *[][]float64 {
	return &p.tau
}
