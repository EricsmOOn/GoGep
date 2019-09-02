package chart

import (
	"fmt"
	"github.com/EricsmOOn/gep-go/gep"
)

var Max_fitness []float64 = make([]float64, 0)

var Ava_fitness []float64 = make([]float64, 0)

func PrintChart() {
	fmt.Println()
	fmt.Println(Max_fitness)
	fmt.Println(Ava_fitness)
}

func GetChartData(genes []*gep.Gene) {
	max := 0.0
	sum := 0.0
	for _, g := range genes {
		sum += g.Fitness
		if g.Fitness > max {
			max = g.Fitness
		}
	}
	Max_fitness = append(Max_fitness, max)
	Ava_fitness = append(Ava_fitness, sum/(float64(len(genes))))
}
