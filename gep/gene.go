package gep

import (
	"math/rand"
	"time"
)

type Gene struct {
	Gene []byte //基因序列
	Fitness int //适应度
}

func creatRandomGene(headLength int, tailLength int,funSet []byte,termSet []byte) Gene {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))//初始化随机数

	funSetNum := len(funSet)
	termSetNum := len(termSet)

	set := append(funSet,termSet...)

	gene := &Gene{make([]byte,0),0}

	for i:=0;i<headLength;i++{
		gene.Gene = append(gene.Gene, set[r.Intn(funSetNum + termSetNum)])
	}

	//fmt.Println(*(*string)(unsafe.Pointer(&gene.Gene)))

	for i:=0;i<tailLength;i++{
		gene.Gene = append(gene.Gene, termSet[r.Intn(termSetNum)])
	}

	return *gene
}

func CreatGenes(num int ,headLength int, tailLength int,funSet []byte,termSet []byte) (genes []Gene) {

	for i := 0; i < num; i++ {
		gene := creatRandomGene(headLength, tailLength, funSet, termSet)
		genes = append(genes, gene)
	}
	return genes
}