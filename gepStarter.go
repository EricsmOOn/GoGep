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
		gep.PrintSelf(genes)
		//终止条件
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

func isEnd(genes []*gep.Gene) bool {
	for _, i := range genes {
		if i.Fitness > gep.ResultRang {
			gep.CalculateFit(i, gep.GetInfixExpressions(*i))
			fmt.Printf("\n最优解:  %s\n", *(*string)(unsafe.Pointer(&i.Gene)))
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
