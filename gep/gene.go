package gep

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

type Gene struct {
	Gene            []byte   //基因序列
	InfixExpression [][]byte //中缀表达式
	Fitness         float64  //适应度
	Generation      int      //代数
}

//显示自己
func PrintSelf(genes []*Gene) {
	fmt.Printf("Generation - [%3d]\n", genes[0].Generation)
	for n, gene := range genes {
		fmt.Printf("%s - [%2d] = %.4f \n", *(*string)(unsafe.Pointer(&gene.Gene)), n, gene.Fitness) //高效转换byte到String

		////显示中缀表达式
		//k := gene.InfixExpression
		//for i := 0; i < len(k); i++ {
		//	fmt.Print(string(k[i]))
		//	if i < len(k)-1 {
		//		fmt.Printf(" %s ", string(ConnectFun))
		//	}
		//}
		//fmt.Println()

	}

}

//随机创建个体
func creatRandomGene(generation int) *Gene {
	funSetNum := len(FunSet)
	termSetNum := len(TermSet)

	set := append(FunSet, TermSet...)

	gene := Gene{make([]byte, 0, NumOfGenes*GeneLength), make([][]byte, 0), 0, generation}

	for k := 0; k < NumOfGenes; k++ {

		for i := 0; i < HeadLength; i++ {
			gene.Gene = append(gene.Gene, set[R.Intn(funSetNum+termSetNum)])
		}

		//fmt.Println(*(*string)(unsafe.Pointer(&gene.Gene)))

		for i := 0; i < TailLength; i++ {
			gene.Gene = append(gene.Gene, TermSet[R.Intn(termSetNum)])
		}
	}

	return &gene
}

//随机创建种群
func CreatGenes() []*Gene {

	var genes []*Gene
	var gene *Gene
	for i := 0; i < PopulationsSize; i++ {
		gene = creatRandomGene(0)
		genes = append(genes, gene)
	}
	return genes
}

//个体遗传
func Evolution(dad Gene, dads []*Gene) Gene {
	//深拷贝
	genes := make([]byte, len(dad.Gene))
	copy(genes, dad.Gene)
	son := Gene{genes, make([][]byte, 0), 0, dad.Generation + 1}
	//变异
	Change(&son, dads)
	//返回
	return son
}

//转盘赌
func Select(genes []*Gene) Gene {
	fitness := float64(0)
	for _, gene := range genes {
		fitness += gene.Fitness
	}
	f := R.Float64() * fitness
	for _, gene := range genes {
		f -= gene.Fitness
		if f <= 0 {
			return *gene
		}
	}
	fmt.Print("Error!")
	return *genes[0]
}
