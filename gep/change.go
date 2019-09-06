package gep

func Change(gene *Gene, genes []*Gene) *Gene {
	if DynamicDcMutation {
		dynamicMutation(gene, genes)
	} else {
		mutation(gene)
	}
	isTransposition(gene)
	risTransposition(gene)
	geneTransposition(gene)
	onePointRecombination(gene, genes)
	twoPointRecombination(gene, genes)
	geneRecombination(gene, genes)
	return gene
}

func mutation(gene *Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	set := append(FunSet, TermSet...)
	if R.Float64() < DcMutationRate {
		intn := R.Intn(len(gene.Gene))
		if intn%GeneLength < HeadLength {
			gene.Gene[intn] = set[R.Intn(len(set))]
		} else {
			gene.Gene[intn] = TermSet[R.Intn(len(TermSet))]
		}
	}
	return gene
}

func dynamicMutation(gene *Gene, genes []*Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	max := 0.0
	sum := 0.0
	rate := 0.0
	for _, g := range genes {
		sum += g.Fitness
		if g.Fitness > max {
			max = g.Fitness
		}
	}
	avg := sum / float64(len(genes))
	f := gene.Fitness
	if f > avg {
		rate = (max - f) / (max - avg)
	} else {
		rate = 1.0
	}

	d := DynamicDcMutationRate
	//rate += (1 - rate) * (avg / max)

	d = d * rate

	set := append(FunSet, TermSet...)
	if R.Float64() < d {
		intn := R.Intn(len(gene.Gene))
		if intn%GeneLength < HeadLength {
			gene.Gene[intn] = set[R.Intn(len(set))]
		} else {
			gene.Gene[intn] = TermSet[R.Intn(len(TermSet))]
		}
	}
	return gene
}

func onePointRecombination(gene *Gene, genes []*Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	if R.Float64() > OnePointRecombinationRate {
		return gene
	}

	no := R.Intn(PopulationsSize)
	pos := R.Intn(GeneLength)

	gene.Gene[pos] = genes[no].Gene[pos]

	return gene
}

func twoPointRecombination(gene *Gene, genes []*Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	if R.Float64() > TwoPointRecombinationRate {
		return gene
	}

	//深度备份源值
	org := make([]byte, len(gene.Gene))
	copy(org, gene.Gene)

	no := R.Intn(PopulationsSize)
	pos := R.Intn(GeneLength)
	length := pos + R.Intn(GeneLength-pos)

	gene.Gene = append(gene.Gene[:pos], genes[no].Gene[pos:pos+length]...)
	gene.Gene = append(gene.Gene, org[pos+length:]...)
	return gene
}

func geneRecombination(gene *Gene, genes []*Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	if R.Float64() > RecombinationRate {
		return gene
	}

	//深度备份源值
	org := make([]byte, len(gene.Gene))
	copy(org, gene.Gene)

	no := R.Intn(PopulationsSize)
	num := R.Intn(NumOfGenes)
	pos := num * GeneLength

	gene.Gene = append(gene.Gene[:pos], genes[no].Gene[pos:(num+1)*GeneLength]...)
	gene.Gene = append(gene.Gene, org[(num+1)*GeneLength:]...)

	return gene
}

func isTransposition(gene *Gene) *Gene {
	if R.Float64() > ISTranspositionRate {
		return gene
	}

	//深度备份源值
	org := make([]byte, len(gene.Gene))
	copy(org, gene.Gene)

	num := R.Intn(NumOfGenes)
	destPos := num*GeneLength + R.Intn(HeadLength-1) + 1
	pos := num*GeneLength + R.Intn(GeneLength-ISElementsLength)
	end := pos + ISElementsLength

	orgTemp := org[:destPos]
	orgTemp = append(orgTemp, gene.Gene[pos:end]...)
	gene.Gene = append(orgTemp[:num*GeneLength+HeadLength], gene.Gene[num*GeneLength+HeadLength:]...)

	return gene
}

func risTransposition(gene *Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	if R.Float64() > RISTranspositionRate {
		return gene
	}

	//深度备份源值
	org := make([]byte, GeneLength*NumOfGenes*2)
	copy(org, gene.Gene)

	num := R.Intn(NumOfGenes)
	destPos := num * GeneLength

	for i := 0; i < 20; i++ {
		pos := num*GeneLength + R.Intn(GeneLength-RISElementsLength)
		for _, k := range FunSet {
			if org[pos] == k {
				end := pos + RISElementsLength
				orgTemp := org[:destPos]
				orgTemp = append(orgTemp, gene.Gene[pos:end]...)
				gene.Gene = append(orgTemp[:num*GeneLength+HeadLength], gene.Gene[num*GeneLength+HeadLength:]...)

				return gene
			}
		}
	}

	return gene
}

func geneTransposition(gene *Gene) *Gene {
	//if FuncTimer {
	//	defer timer.TimeCount()()
	//}
	if R.Float64() > GeneTranspositionRate {
		return gene
	}

	num1 := R.Intn(NumOfGenes)
	num2 := R.Intn(NumOfGenes)

	for num1 == num2 {
		num2 = R.Intn(NumOfGenes)
	}

	k := make([][]byte, NumOfGenes)
	m := make([]byte, GeneLength)
	for i := 0; i < NumOfGenes; i++ {
		copy(m, gene.Gene[GeneLength*i:GeneLength*(i+1)])
		k[i] = append(k[i], m...)
	}

	gene.Gene = gene.Gene[:0]

	k[num1], k[num2] = k[num2], k[num1]

	for i := 0; i < NumOfGenes; i++ {
		gene.Gene = append(gene.Gene, k[i]...)
	}

	return gene
}
