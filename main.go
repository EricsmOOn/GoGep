package main

import (
	"fmt"
	"github.com/EricsmOOn/GoGep/chart"
	"github.com/EricsmOOn/GoGep/gep"
	"net/http"
	"strconv"
)

func main() {

	//读取数据集
	gep.InitSampleData()
	//初始化种群
	genes := gep.CreateGenes()

	for {
		//计算父代适应度
		gep.CalculateFitnessOpt(genes)
		gep.Wg.Wait()
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
			//展示图表
			http.HandleFunc("/", chart.Handler)
			e := http.ListenAndServe(":"+strconv.Itoa(gep.ChartPort), nil)
			if e != nil {
				fmt.Print(e.Error())
			}
			return
		}
		//进化
		sons := gep.Evolution(genes)
		//替换为父代
		copy(genes, sons)
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
			chart.GetPredictResult(*g)
			gep.PrintGreat(g)
			return true
		}
	}
	for _, i := range genes {
		if i.Fitness >= gep.ResultRang {
			chart.GetPredictResult(*i)
			gep.PrintGreat(i)
			return true
		}
	}
	return false
}
