package main

import (
	"fmt"
	"gep-go/gep"
	"math"
	"unsafe"
)

//头部长度
const HeadLength int = 5

func main() {

	//函数集
	funSet := []byte{'+', '-', '*', '/','N','Q'}
	//终点集
	termSet := []byte{'a', 'b'}
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := HeadLength*(maxFactorNum-1) + 1

	genes := gep.CreatGenes(10, HeadLength, tailLength, funSet, termSet)

	testData := gep.ReadTestData()
	var termVarSet []byte

	for {
		for n, gene := range genes {
			//fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
			//fmt.Printf("-[%d]", n)

			//解码
			g := gep.Operate(gene, termSet)

			for _, td := range testData {
				//求个体适应度

				//逐个读入测试数据
				termVarSet = td.TermVarSet

				//求表达 result
				result, err := gep.Calculate(g, termVarSet, termSet)
				if err == nil {
					if math.Abs(float64(result-td.Result)) <= 1 {
						gene.Fitness ++
					}
					//fmt.Print(" - ")
					//fmt.Println(err)
				}
				////显示结果
				//fmt.Print(" = ")
				//fmt.Println(result)
				////显示解析出的中缀表达式
				//fmt.Print(" --> ")
				//fmt.Println(*(*string)(unsafe.Pointer(&g)))
			}
			//fmt.Print(" = ")
			//fmt.Println(gene.Fitness)
			if gene.Fitness >= 2{
				fmt.Print(*(*string)(unsafe.Pointer(&gene.Gene))) //高效转换byte到String
				fmt.Printf("-[%d]", n)
				fmt.Print(" = ")
				fmt.Println(gene.Fitness)
				return
			}
		}
	}

}
