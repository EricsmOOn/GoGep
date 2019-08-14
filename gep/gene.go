package gep

import (
	"math"
	"math/rand"
	"time"
)

type Gene struct {
	Gene    []byte  //基因序列
	Fitness float64 //适应度
}

func creatRandomGene(numOfGenes int, headLength int, tailLength int, funSet []byte, termSet []byte) *Gene {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //初始化随机数

	funSetNum := len(funSet)
	termSetNum := len(termSet)

	set := append(funSet, termSet...)

	gene := Gene{make([]byte, 0), 0}

	for k := 0; k < numOfGenes; k++ {

		for i := 0; i < headLength; i++ {
			gene.Gene = append(gene.Gene, set[r.Intn(funSetNum+termSetNum)])
		}

		//fmt.Println(*(*string)(unsafe.Pointer(&gene.Gene)))

		for i := 0; i < tailLength; i++ {
			gene.Gene = append(gene.Gene, termSet[r.Intn(termSetNum)])
		}
	}

	return &gene
}

func CreatGenes(numOfGenes int, num int, headLength int, tailLength int, funSet []byte, termSet []byte) []*Gene {

	var genes []*Gene
	var gene *Gene
	for i := 0; i < num; i++ {
		gene = creatRandomGene(numOfGenes, headLength, tailLength, funSet, termSet)
		genes = append(genes, gene)
	}
	return genes
}

func CalculateFitness(genes []*Gene, termSet []byte, mM int) {
	testData := ReadTestData()
	var termVarSet []float64

	for _, gene := range genes {
		//解码
		g := Operate(*gene, termSet)

		for _, td := range testData {
			//求个体适应度

			//逐个读入测试数据
			termVarSet = td.TermVarSet

			//求表达 result
			result, err := Calculate(g, termVarSet, termSet)
			if err == nil {
				fi := float64(mM) - (math.Abs((result-td.Result)/td.Result) * 100)
				if fi > 0 {
					gene.Fitness += fi
				}
				//fmt.Print(" - ")
				//fmt.Println(err)
			}
		}
	}

}
