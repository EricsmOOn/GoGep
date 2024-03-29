package gep

import (
	"math/rand"
	"time"
)

var R = rand.New(rand.NewSource(time.Now().UnixNano()))

type Gene struct {
	Gene            []byte   //基因序列
	InfixExpression [][]byte //中缀表达式
	Fitness         float64  //适应度
	Generation      int      //代数
}

//随机创建个体
func creatRandomGene(generation int) *Gene {
	funSetNum := len(FunSet)
	termSetNum := len(TermSet)

	set := append(FunSet, TermSet...)

	gene := Gene{make([]byte, 0, NumOfGenes*GeneLength), make([][]byte, 0), 0, generation}

	for k := 0; k < NumOfGenes; k++ {
		if MoreFunc {
			for i := 0; i < HeadLength; i++ {
				if R.Float64() > 0.5 {
					gene.Gene = append(gene.Gene, FunSet[R.Intn(funSetNum)])
				} else {
					gene.Gene = append(gene.Gene, TermSet[R.Intn(termSetNum)])
				}
			}
		} else {
			for i := 0; i < HeadLength; i++ {
				gene.Gene = append(gene.Gene, set[R.Intn(funSetNum+termSetNum)])
			}
		}

		//fmt.Println(*(*string)(unsafe.Pointer(&gene.Gene)))

		for i := 0; i < TailLength; i++ {
			gene.Gene = append(gene.Gene, TermSet[R.Intn(termSetNum)])
		}
	}

	return &gene
}

//随机创建种群
func CreateGenes() []*Gene {

	var genes []*Gene
	var gene *Gene
	for i := 0; i < PopulationsSize; i++ {
		gene = creatRandomGene(0)
		genes = append(genes, gene)
	}
	return genes
}

//个体遗传
func Evolution(dads []*Gene) []*Gene {
	sons := make([]*Gene, 0, PopulationsSize)
	sons = Select(dads, sons)
	return sons
}

//选择函数
func Select(genes []*Gene, sons []*Gene) []*Gene {
	//变异转盘赌
	sons = Turn(genes, TurnNum, sons)
	//变异精英
	sons = Elite(genes, EliteNum, sons)
	//不变异精英
	sons = EliteNone(genes, NonEliteNum, sons)
	return sons
}

//转盘赌
func Turn(genes []*Gene, num int, sons []*Gene) []*Gene {
	for i := 0; i < num; i++ {
		fitness := float64(0)
		for _, gene := range genes {
			fitness += gene.Fitness
		}
		f := R.Float64() * fitness
		detg := genes[0]
		for _, gene := range genes {
			f -= gene.Fitness
			if f <= 0 {
				detg = gene
			}
		}
		sons = append(sons, Change(deepCopy(detg), genes))
	}
	return sons
}

//精英策略
func Elite(genes []*Gene, num int, sons []*Gene) []*Gene {
	for i := 0; i < num; i++ {
		elite := genes[0]
		for _, g := range genes {
			if g.Fitness > elite.Fitness {
				elite = g
			}
		}
		sons = append(sons, Change(deepCopy(elite), genes))
	}
	return sons
}

//不变异精英策略
func EliteNone(genes []*Gene, num int, sons []*Gene) []*Gene {
	for i := 0; i < num; i++ {
		elite := genes[0]
		for _, g := range genes {
			if g.Fitness > elite.Fitness {
				elite = g
			}
		}
		sons = append(sons, deepCopy(elite))
	}
	return sons
}

//遗传深拷贝工具
func deepCopy(dad *Gene) *Gene {
	//深拷贝
	genes := make([]byte, len(dad.Gene))
	copy(genes, dad.Gene)
	son := Gene{genes, make([][]byte, 0), 0, dad.Generation + 1}
	return &son
}
