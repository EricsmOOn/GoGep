package gep

import (
	"math"
	"math/rand"
	"time"
)

var R *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type Gene struct {
	Gene       []byte  //基因序列
	Fitness    float64 //适应度
	Generation int     //代数
}

func creatRandomGene(generation int, numOfGenes int, headLength int, tailLength int, funSet []byte, termSet []byte) *Gene {
	funSetNum := len(funSet)
	termSetNum := len(termSet)

	set := append(funSet, termSet...)

	gene := Gene{make([]byte, 0), 0, generation}

	for k := 0; k < numOfGenes; k++ {

		for i := 0; i < headLength; i++ {
			gene.Gene = append(gene.Gene, set[R.Intn(funSetNum+termSetNum)])
		}

		//fmt.Println(*(*string)(unsafe.Pointer(&gene.Gene)))

		for i := 0; i < tailLength; i++ {
			gene.Gene = append(gene.Gene, termSet[R.Intn(termSetNum)])
		}
	}

	return &gene
}

func CreatGenes(numOfGenes int, num int, headLength int, tailLength int, funSet []byte, termSet []byte) []*Gene {

	var genes []*Gene
	var gene *Gene
	for i := 0; i < num; i++ {
		gene = creatRandomGene(0, numOfGenes, headLength, tailLength, funSet, termSet)
		genes = append(genes, gene)
	}
	return genes
}

func Evolution(genetic *Gene, numOfGenes int, num int, headLength int, tailLength int, funSet []byte, termSet []byte) []*Gene {
	var genes []*Gene
	var gene *Gene
	genetic.Generation += 1
	genetic.Fitness = 0
	genes = append(genes, genetic)
	for i := 0; i < num-1; i++ {
		gene = creatRandomGene(genetic.Generation, numOfGenes, headLength, tailLength, funSet, termSet)
		genes = append(genes, gene)
	}
	return genes
}

func Select(genes []*Gene) *Gene {
	fitness := float64(0)
	for _, gene := range genes {
		fitness += gene.Fitness
	}
	f := R.Float64() * fitness
	for _, gene := range genes {
		f -= gene.Fitness
		if f <= 0 {
			return gene
		}
	}
	return nil
}

func Mutation(mutationRate float64, gene *Gene, headLength, geneLength int, funSet, termSet []byte) *Gene {
	funSet = append(funSet, termSet...)
	if R.Float64() < mutationRate {
		intn := R.Intn(len(gene.Gene))
		if intn%geneLength < headLength {
			gene.Gene[intn] = funSet[R.Intn(len(funSet))]
		} else {
			gene.Gene[intn] = termSet[R.Intn(len(termSet))]
		}
	}
	return gene
}

func CalculateFitness(connectFun byte, geneLength int, numOfGenes int, genes []*Gene, termSet []byte, mM float64) {
	testData := ReadTestData()
	var termVarSet []float64

	for _, gene := range genes {
		//解码
		g := Operate(geneLength, numOfGenes, *gene, termSet)

		for _, td := range testData {
			//求个体适应度

			//逐个读入测试数据
			termVarSet = td.TermVarSet

			//求表达 result
			result := Calculate(connectFun, g, termVarSet, termSet)
			fi := mM - math.Abs(result-td.Result)
			if fi > 0 {
				gene.Fitness += fi
			}
			//fmt.Print(" - ")
			//fmt.Println(err)
		}
	}

}
