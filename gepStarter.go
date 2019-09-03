package main

import (
	"fmt"
	"github.com/EricsmOOn/gep-go/gep"
	"unsafe"
)

func main() {

	genes := gep.CreatGenes()

	for true {
		//计算父代适应度
		gep.CalculateFitnessOpt(genes)
		//计算父代适应度(逆波兰) 优化 - 10倍 速度
		//gep.CalculateFitness(genes)
		//图表获取数据
		//chart.GetChartData(genes)
		//显示每一代数据
		//gep.PrintSelf(genes)
		//显示简易数据
		//gep.PrintSelfEasy(genes)
		//显示最简数据
		gep.PrintMostEasy(genes)
		//终止条件(genes,最大运行代数(可选))
		if isEnd(genes) {
			//展示函数耗时情况
			//timer.PrintTimer()
			//展示图表 http://localhost:8081/
			//http.HandleFunc("/", chart.Handler)
			//http.ListenAndServe(":8081", nil)
			return
		}
		//进化
		sons := gep.Evolution(genes)
		//替换为父代
		copy(genes, sons)
		//清空子代
		sons = sons[:0]
	}
}

func isEnd(genes []*gep.Gene, maxGenerations ...int) bool {
	if len(maxGenerations) != 0 {
		//封顶
		if genes[0].Generation == maxGenerations[0] {
			g := genes[0]
			for _, i := range genes {
				if i.Fitness > g.Fitness {
					g = i
				}
			}
			gep.CalculateFit(g, gep.GetInfixExpressions(*g))
			fmt.Printf("\n最优解:  %s - [%d] - [%5f]\n", *(*string)(unsafe.Pointer(&g.Gene)), g.Generation, g.Fitness)
			fmt.Print("中缀式:  ")
			for t := 0; t < gep.NumOfGenes; t++ {
				fmt.Printf("%s", *(*string)(unsafe.Pointer(&g.InfixExpression[t])))
				if t < gep.NumOfGenes-1 {
					fmt.Printf(string(gep.LinkFun))
				}
			}
			fmt.Println()
			return true
		}
	}
	for _, i := range genes {
		if i.Fitness > gep.ResultRang {
			gep.CalculateFit(i, gep.GetInfixExpressions(*i))
			fmt.Printf("\n最优解:  %s - [Generation:%d] - [Fitness:%5f]\n", *(*string)(unsafe.Pointer(&i.Gene)), i.Generation, i.Fitness)
			fmt.Print("中缀式:  ")
			for t := 0; t < gep.NumOfGenes; t++ {
				fmt.Printf("%s", *(*string)(unsafe.Pointer(&i.InfixExpression[t])))
				if t < gep.NumOfGenes-1 {
					fmt.Printf(string(gep.LinkFun))
				}
			}
			fmt.Println()
			return true
		}
	}
	return false
}
