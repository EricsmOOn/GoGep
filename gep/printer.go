package gep

import (
	"fmt"
)

//显示自己
func PrintSelf(genes []*Gene) {
	fmt.Printf("Generation - [%3d]\n", genes[0].Generation)
	for n, gene := range genes {
		fmt.Printf("%s - [%6d] = %.4f \n", gene.Gene, n, gene.Fitness)
	}
}

//显示最大适应度
func PrintSelfEasy(genes []*Gene) {
	g := genes[0]
	for _, i := range genes {
		if i.Fitness > g.Fitness {
			g = i
		}
	}
	fmt.Printf("[%6d] - [%s] - [%.4f]\n", g.Generation, g.Gene, g.Fitness)
}

var MaxFitness float64 = 0

//显示最大适应度
func PrintMostEasy(genes []*Gene) {
	g := genes[0]
	flag := false
	for _, i := range genes {
		if i.Fitness > MaxFitness {
			MaxFitness = i.Fitness
			g = i
			flag = true
		}
	}
	if flag {
		fmt.Printf("[%6d] - [%s] - [%.4f]\n", g.Generation, g.Gene, g.Fitness)
		//g.InfixExpression = GetInfixExpressions(*g)
		//fmt.Print("中缀式:  ")
		//for t := 0; t < NumOfGenes; t++ {
		//	fmt.Printf("%s", g.InfixExpression[t])
		//	if t < NumOfGenes-1 {
		//		fmt.Print(string(LinkFun))
		//	}
		//}
		//fmt.Println()
		//
	}
}

func PrintGreat(g *Gene) {
	g.InfixExpression = GetInfixExpressions(*g)
	fmt.Printf("\n最优解:  %s - [%d] - [%5f]\n", g.Gene, g.Generation, g.Fitness)
	fmt.Print("中缀式:  ")
	for t := 0; t < NumOfGenes; t++ {
		fmt.Printf("%s", g.InfixExpression[t])
		if t < NumOfGenes-1 {
			fmt.Print(string(LinkFun))
		}
	}
	fmt.Println()

}
