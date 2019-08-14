package main

import (
	"fmt"
	"gep-go/gep"
	"unsafe"
)

func main() {

	//头部长度
	headLength := 7
	//基因产生个数
	populationsSize := 20
	//选择范围
	mM := float64(100)
	//染色体含有基因数
	numOfGenes := 3
	//连接函数
	connectFun := byte('+')
	//函数集
	funSet := []byte{'+', '-', '*', '/'}
	//终点集
	termSet := []byte{'a'}
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := headLength*(maxFactorNum-1) + 1
	//基因长度
	geneLength := headLength + tailLength

	//test(populationsSize,numOfGenes,geneLength,headLength,tailLength,funSet,termSet,connectFun,mM)

	//1.最小适应度 2.最大适应度 3.需要的个数
	findFintness(999.99, 1000, 10, populationsSize, numOfGenes, geneLength, headLength, tailLength, funSet, termSet, connectFun, mM)

}

func test(populationsSize, numOfGenes, geneLength, headLength, tailLength int, funSet, termSet []byte, connectFun byte, mM float64) {
	genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
	gep.CalculateFitness(connectFun, geneLength, numOfGenes, genes, termSet, mM)
	for n, gene := range genes {
		//显示结果
		fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
		fmt.Printf("-[%2d]", n)
		fmt.Print(" = ")
		fmt.Print(gene.Fitness)

		////显示中缀表达式
		//fmt.Print(" --> ")
		//k := gep.Operate(geneLength, numOfGenes, *gene, termSet)
		//for i := 0; i < len(k); i++ {
		//	fmt.Print(string(k[i]))
		//	if i < len(k)-1 {
		//		fmt.Printf(" %s ", string(connectFun))
		//	}
		//}

		fmt.Println()
	}
}

func findFintness(minFitness, maxFitness float64, num, populationsSize, numOfGenes, geneLength, headLength, tailLength int, funSet, termSet []byte, connectFun byte, mM float64) {
	l := 0
	for {
		genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
		gep.CalculateFitness(connectFun, geneLength, numOfGenes, genes, termSet, mM)
		for n, gene := range genes {
			if gene.Fitness > minFitness && gene.Fitness < maxFitness {
				//显示结果
				fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
				fmt.Printf("-[%2d]", n)
				fmt.Print(" = ")
				fmt.Print(gene.Fitness)

				//显示中缀表达式
				fmt.Print(" --> ")
				k := gep.Operate(geneLength, numOfGenes, *gene, termSet)
				for i := 0; i < len(k); i++ {
					fmt.Print(string(k[i]))
					if i < len(k)-1 {
						fmt.Printf(" %s ", string(connectFun))
					}
				}
				fmt.Println()
				l++
				if l >= num {
					return
				}
			}
		}
	}
}
