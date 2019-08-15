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
	selectRang := float64(100)
	//染色体含有基因数
	numOfGenes := 3
	//连接函数
	connectFun := byte('+')
	//函数集
	funSet := []byte{'+', '-', '*', '/'}
	//终点集
	termSet := []byte{'a'}
	//变异率
	mutationRate := 0.5
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := headLength*(maxFactorNum-1) + 1
	//基因长度
	geneLength := headLength + tailLength

	//1.寻找的最小适应度
	//test(999,populationsSize,numOfGenes,geneLength,headLength,tailLength,funSet,termSet,connectFun,mutationRate,selectRang)

	//1.寻找的最小适应度2.执行多少代以内
	test2(999, 10, populationsSize, numOfGenes, geneLength, headLength, tailLength, funSet, termSet, connectFun, mutationRate, selectRang)

	//1.最小适应度 2.最大适应度 3.需要的个数
	//findFintness(999, 1000, 10, populationsSize, numOfGenes, geneLength, headLength, tailLength, funSet, termSet, connectFun, selectRang)

}

func test(fitness float64, populationsSize, numOfGenes, geneLength, headLength, tailLength int, funSet, termSet []byte, connectFun byte, mutationRate, selectRang float64) {
	flag := true
	genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
	for {
		gep.CalculateFitness(connectFun, geneLength, numOfGenes, genes, termSet, selectRang)
		fmt.Printf("Geneation %d : \n", genes[0].Generation)
		for n, gene := range genes {
			if gene.Fitness > fitness {
				flag = false
			}
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
		if !flag {
			return
		}
		mutation := gep.Mutation(mutationRate, gep.Select(genes), headLength, geneLength, funSet, termSet)
		//显示遗传个体
		fmt.Println()
		fmt.Print("遗传Gene:  ")
		fmt.Print(*(*string)(unsafe.Pointer(&mutation.Gene))) //高效转换byte到String
		fmt.Print(" = ")                                      //高效转换byte到String
		fmt.Println(mutation.Fitness)
		fmt.Println()
		genes = gep.Evolution(mutation, numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
	}
}

func findFintness(minFitness, maxFitness float64, num, populationsSize, numOfGenes, geneLength, headLength, tailLength int, funSet, termSet []byte, connectFun byte, selectRang float64) {
	l := 0
	for {
		genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
		gep.CalculateFitness(connectFun, geneLength, numOfGenes, genes, termSet, selectRang)
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

func test2(fitness float64, geneation, populationsSize, numOfGenes, geneLength, headLength, tailLength int, funSet, termSet []byte, connectFun byte, mutationRate, selectRang float64) {
outfor:
	for {
		genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
		for {
			gep.CalculateFitness(connectFun, geneLength, numOfGenes, genes, termSet, selectRang)
			for _, gene := range genes {
				if gene.Generation >= geneation {
					goto outfor
				}
				if gene.Fitness > fitness && gene.Generation < geneation {
					fmt.Printf("Geneation %d : \n", genes[0].Generation)
					for n, gene := range genes {
						fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
						fmt.Printf("-[%2d]", n)
						fmt.Print(" = ")
						fmt.Print(gene.Fitness)
						fmt.Println()
					}
					return
				}
				//显示结果
			}
			mutation := gep.Mutation(mutationRate, gep.Select(genes), headLength, geneLength, funSet, termSet)
			genes = gep.Evolution(mutation, numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)
		}
	}
}
