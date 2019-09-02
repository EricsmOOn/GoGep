package main

import (
	"fmt"
	"github.com/EricsmOOn/gep-go/chart"
	"github.com/EricsmOOn/gep-go/gep"
	"unsafe"
)

func main() {

	genes := gep.CreatGenes()
	sons := make([]*gep.Gene, 0, gep.PopulationsSize)

	for true {
		//计算父代适应度
		gep.CalculateFitness(genes)
		//插入图标
		chart.GetChartData(genes)
		//显示自己
		//gep.PrintSelf(genes)
		//终止条件
		for _, i := range genes {
			if i.Fitness > gep.ResultRang {
				fmt.Printf("\n最优解:  %s\n", *(*string)(unsafe.Pointer(&i.Gene)))
				fmt.Print("中缀式:  ")
				for t := 0; t < gep.NumOfGenes; t++ {
					fmt.Printf("%s", *(*string)(unsafe.Pointer(&i.InfixExpression[t])))
					if t < gep.NumOfGenes-1 {
						fmt.Print(string(gep.LinkFun))
					}
				}
				//展示图表
				chart.PrintChart()
				return
			}
		}
		//进化
		for i := 0; i < gep.PopulationsSize; i++ {
			//选择进化
			son := gep.Evolution(gep.Select(genes), genes)
			//加入下一代
			sons = append(sons, &son)
		}
		//替换为父代
		copy(genes, sons)
		//清空子代
		sons = sons[:0]
	}
}
