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
	mM := 100
	//染色体含有基因数
	numOfGenes := 1
	//误差精度

	//函数集
	funSet := []byte{'+', '-', '*', '/'}
	//终点集
	termSet := []byte{'a'}
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := headLength*(maxFactorNum-1) + 1

	genes := gep.CreatGenes(numOfGenes, populationsSize, headLength, tailLength, funSet, termSet)

	gep.CalculateFitness(genes, termSet, mM)

	for n, gene := range genes {
		//显示结果
		fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
		fmt.Printf("-[%2d]", n)
		fmt.Print(" = ")
		fmt.Println(gene.Fitness)
	}
}
