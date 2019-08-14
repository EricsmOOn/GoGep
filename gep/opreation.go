package gep

import (
	"container/list"
)

type Operation struct {
	symbol    byte //符号
	factorNum int  //因子个数
	priority  int  //优先级 越大优先级越高
}

func GetOperationFactorNum(operationName byte) int {
	num := -1
	switch operationName {
	case '+', '-', '*', '/', '%':
		num = 2
	case 'Q', 'N':
		num = 1
	default:
		num = -1
	}
	return num
}

func GetOperationPriority(operationName byte) int {
	priority := -1
	switch operationName {
	case '+', '-':
		priority = 4
	case '*', '/', '%':
		priority = 3
	case '(':
		priority = 0
	case 'Q', 'N':
		priority = 6
	default:
		priority = -1
	}
	return priority
}

func GetMaxFactorNum(operationNames []byte) int {
	maxNum := 0
	for _, name := range operationNames {
		i := GetOperationFactorNum(name)
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}

func isTerm(factor byte, termSet []byte) bool {
	for _, value := range termSet {
		if value == factor {
			return true
		}
	}
	return false
}

func Operate(geneLength int, numOfGenes int, gene Gene, termSet []byte) (k [][]byte) {
	for i := 0; i < numOfGenes; i++ {
		k = append(k, operate(gene.Gene[i*geneLength:(i+1)*geneLength], termSet))
	}
	return k
}

func operate(s []byte, termSet []byte) []byte {
	//将基因缩短为有效长度

	si := 0
	oi := 1

	if isTerm(s[0], termSet) {
		s = s[:1]
	}

	for {
		if si == oi && si != 0 {
			break
		}
		if !isTerm(s[si], termSet) {
			oi += GetOperationFactorNum(s[si])
		}
		si++
	}

	s = s[:si]

	//用序号代表基因所在位置 用双向链表构建解码表达式 表示括号:-1->( -2->)
	link := list.New()
	i := 0
	j := 0
	pos := link.PushBack(i)
	if !isTerm(s[i], termSet) {
		for i < len(s) {
			//查找i节点的位置
			pos = search(link, i)
			if GetOperationFactorNum(s[i]) == 1 {
				j++
				if !isTerm(s[j], termSet) {
					pos1 := link.InsertAfter(-1, pos)
					pos1 = link.InsertAfter(j, pos1)
					pos1 = link.InsertAfter(-2, pos1)
				} else {
					link.InsertAfter(j, pos)
				}
			} else if GetOperationFactorNum(s[i]) == 2 {
				j++
				if !isTerm(s[j], termSet) {
					pos1 := link.InsertBefore(-2, pos)
					pos1 = link.InsertBefore(j, pos1)
					pos1 = link.InsertBefore(-1, pos1)
				} else {
					link.InsertBefore(j, pos)
				}
				j++
				if !isTerm(s[j], termSet) {
					pos2 := link.InsertAfter(-1, pos)
					pos2 = link.InsertAfter(j, pos2)
					pos2 = link.InsertAfter(-2, pos2)
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

	//fmt.Println(*(*string)(unsafe.Pointer(&result)))
	return result
}

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
