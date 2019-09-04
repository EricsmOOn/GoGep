package main

import (
	"github.com/EricsmOOn/gep-go/chart"
	"github.com/EricsmOOn/gep-go/gep"
	"github.com/EricsmOOn/gep-go/util/timer"
	"net/http"
)

func main() {

	genes := gep.CreateGenes()

	for true {
		//计算父代适应度
		gep.CalculateFitnessOpt(genes)
		//计算父代适应度(逆波兰) 优化 - 10倍 速度
		//gep.CalculateFitness(genes)
		//图表获取数据
		if gep.Chart {
			chart.GetChartData(genes)
		}
		switch gep.ViewStyle {
		case 0:
			//显示每一代数据
			gep.PrintSelf(genes)
		case 1:
			//显示简易数据
			gep.PrintSelfEasy(genes)
		case 2:
			//显示最简数据
			gep.PrintMostEasy(genes)
		default:
		}
		//终止条件(genes,最大运行代数(可选))
		if isEnd(genes, gep.MaxGenerations) {
			//展示函数耗时情况
			if gep.FuncTimer {
				timer.PrintTimer()
			}
			//展示图表 http://localhost:8081/
			if gep.Chart {
				http.HandleFunc("/", chart.Handler)
				http.ListenAndServe(":8081", nil)
			}
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
	if maxGenerations[0] != 0 {
		//封顶
		if genes[0].Generation == maxGenerations[0] {
			g := genes[0]
			for _, i := range genes {
				if i.Fitness > g.Fitness {
					g = i
				}
			}
			gep.PrintGreat(g)
			return true
		}
	}
	for _, i := range genes {
		if i.Fitness > gep.ResultRang {
			gep.PrintGreat(i)
			return true
		}
	}
	return false
}
