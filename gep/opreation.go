package gep

import (
	"container/list"
)

//获得函数参数个数
func GetOperationFactorNum(operationName byte) int {
	var num int
	switch operationName {
	case '+', '-', '*', '/', '%':
		num = 2
	case 'Q', 'N', '@', '$':
		num = 1
	default:
		num = -1
	}
	return num
}

//获得函数集最大参数个数
func GetMaxFactorNum() int {
	maxNum := 0
	for _, name := range FunSet {
		i := GetOperationFactorNum(name)
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}

//判断函数是否为终点集
func isTerm(factor byte) bool {
	for _, value := range TermSet {
		if value == factor {
			return true
		}
	}
	return false
}

//获得个体中缀表达式
func GetInfixExpressions(gene Gene) (k [][]byte) {
	for i := 0; i < NumOfGenes; i++ {
		k = append(k, GetInfixExpression(gene.Gene[i*GeneLength:(i+1)*GeneLength]))
	}
	return k
}

//获得单个基因中缀表达式
func GetInfixExpression(s []byte) []byte {

	s = GetEffectGene(s)

	//用序号代表基因所在位置 用双向链表构建解码表达式 表示括号: -1 -> ( , -2 -> )
	link := list.New()
	i := 0
	j := 0
	link.PushBack(i)
	if !isTerm(s[i]) {
		for i < len(s) {
			//查找i节点的位置
			pos := search(link, i)
			if GetOperationFactorNum(s[i]) == 1 {
				j++
				if !isTerm(s[j]) {
					pos1 := link.InsertAfter(-1, pos)
					pos1 = link.InsertAfter(j, pos1)
					link.InsertAfter(-2, pos1)
				} else {
					link.InsertAfter(j, pos)
				}
			} else if GetOperationFactorNum(s[i]) == 2 {
				j++
				if !isTerm(s[j]) {
					pos1 := link.InsertBefore(-2, pos)
					pos1 = link.InsertBefore(j, pos1)
					link.InsertBefore(-1, pos1)
				} else {
					link.InsertBefore(j, pos)
				}
				j++
				if !isTerm(s[j]) {
					pos2 := link.InsertAfter(-1, pos)
					pos2 = link.InsertAfter(j, pos2)
					link.InsertAfter(-2, pos2)
				} else {
					link.InsertAfter(j, pos)
				}
			}
			i = i + 1
		}
	}

	//基因解码后序号反代
	var result []byte

	f := link.Front()
	for i := 0; i < link.Len(); i++ {
		if f.Value.(int) >= 0 {
			result = append(result, s[f.Value.(int)])
		} else if f.Value.(int) == -1 {
			result = append(result, byte('('))
		} else if f.Value.(int) == -2 {
			result = append(result, byte(')'))
		}
		f = f.Next()
	}

	return result
}

//寻找双向链表的节点
func search(link *list.List, aim int) *list.Element {
	f := link.Front()
	for i := link.Len(); i > 0; i-- {
		if f.Value.(int) == aim {
			return f
		}
		f = f.Next()
	}
	return f
}

//将基因缩短为有效长度
func GetEffectGene(s []byte) []byte {
	si := 0
	oi := 1

	if isTerm(s[0]) {
		s = s[:1]
	}

	for {
		if si == oi && si != 0 {
			break
		}
		if !isTerm(s[si]) {
			oi += GetOperationFactorNum(s[si])
		}
		si++
	}

	s = s[:si]
	return s
}

//获得个体最简表达式
func GetEffectGenes(gene Gene) (k [][]byte) {
	for i := 0; i < NumOfGenes; i++ {
		k = append(k, GetEffectGene(gene.Gene[i*GeneLength:(i+1)*GeneLength]))
	}
	return k
}
