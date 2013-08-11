package aco

import (
	"testing"
)

func TestPheromone(t *testing.T) {
	init, dtau := 1.0, 1.0
	length := 10
	pheromone := NewPheromone(init, dtau, length)

	pInit := pheromone.Init()
	if init != pInit {
		t.Errorf("pheromone.init = %.3d, want %.3d.", pInit, init)
	}

	pDtau := pheromone.Dtau()
	if dtau != pDtau {
		t.Errorf("pheromone.dtau = %.3d, want %.3d.", pDtau, dtau)
	}

	pLength := pheromone.Length()
	if length != pLength {
		t.Errorf("pheromone.length = %.3d, want %.3d.", pLength, length)
	}

	tau := *pheromone.Tau()
	for i := range tau {
		for j := range tau[i] {
			if init != tau[i][j] {
				t.Errorf("pheromone.tau[%d][%d] = %.3f, want %.3f.", i, j, tau[i][j], init)
			}
		}
	}
}
