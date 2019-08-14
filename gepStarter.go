package main

import (
	"fmt"
	"gep-go/gep"
	"math"
	"unsafe"
)

//头部长度
const HeadLength int = 3

func main() {

	//基因产生个数
	generateGenesNum := 7
	//选择范围
	mM := 100
	//误差精度

	//函数集
	funSet := []byte{'+', '-', '*', '/'}
	//终点集
	termSet := []byte{'a'}
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := HeadLength*(maxFactorNum-1) + 1

	genes := gep.CreatGenes(generateGenesNum, HeadLength, tailLength, funSet, termSet)

	testData := gep.ReadTestData()
	var termVarSet []float64

	for n, gene := range genes {
		fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
		fmt.Printf("-[%d]", n)

		//解码
		g := gep.Operate(gene, termSet)

		for _, td := range testData {
			//求个体适应度

			//逐个读入测试数据
			termVarSet = td.TermVarSet

			//求表达 result
			result, err := gep.Calculate(g, termVarSet, termSet)
			if err == nil {
				fi := float64(mM) - (math.Abs((result-td.Result)/td.Result) * 100)
				if fi > 0 {
					gene.Fitness += fi
				}
				//fmt.Print(" - ")
				//fmt.Println(err)
			}
			////显示结果
			//fmt.Print(" = ")
			//fmt.Print(result)
			////显示解析出的中缀表达式
			//fmt.Print(" --> ")
			//fmt.Println(*(*string)(unsafe.Pointer(&g)))
		}
		fmt.Print(" = ")
		fmt.Print(gene.Fitness)
		fmt.Print(" --> ")
		fmt.Print(*(*string)(unsafe.Pointer(&g)))
		fmt.Println()
	}

}
