package gep

import (
	"errors"
	"math"
)

//计算种群适应度
func CalculateFitnessOpt(genes []*Gene) {
	for _, gene := range genes {
		calculateFitOpt(gene, GetEffectGenes(*gene))
	}
}

//简化式子后计算个体适应度
func calculateFitOpt(gene *Gene, easyEquation [][]byte) {
	testData := ReadTestData()
	fsum := 0.0
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, err := calculatePerFitOpt(easyEquation, td.TermVarSet)
		//杀死非法表达式
		if err == nil {
			gene.Fitness = 0
		}
		fi := SelectRang - math.Abs(result-td.Result)
		if fi > 0 {
			fsum += fi
		}
	}
	gene.Fitness = fsum
}

//计算个体对一组数据的适应度
func calculatePerFitOpt(g [][]byte, v []float64) (float64, error) {
	var result float64
	for i := 0; i < len(g); i++ {
		//拼接个体的几个基因到 slice
		slice := make([]interface{}, 0)
		for _, value := range g[i] {
			slice = append(slice, value)
		}
		//替换终结符到 slice
		for gno, b := range slice {
			switch b.(type) {
			case byte:
				for ts, e := range TermSet {
					if e == b {
						tmp := slice[gno+1:]
						slice = append(slice[:gno], v[ts])
						slice = append(slice, tmp...)
					}
				}
			default:
			}
		}
		//计算
		f, e := calculateOpt(slice)
		if e != nil {
			return 0, e
		}
		//连接函数
		switch LinkFun {
		case '+':
			result += f
		case '-':
			result -= f
		case '*':
			result = result * f
		case '/':
			result = result / f
		}
	}
	return result, nil
}

//求个体对一条样例的适应度
func calculateOpt(exp []interface{}) (float64, error) {
	for i := len(exp) - 1; i >= 0; i-- {
		switch exp[i].(type) {
		case byte:
			char := exp[i].(byte)
			//一元函数
			if GetOperationFactorNum(char) == 1 {
				x := 0.0
				//找一个终结符
				for i := len(exp) - 1; i > 0; i-- {
					switch exp[i].(type) {
					case float64:
						x = exp[i].(float64)
						break
					}
					break
				}
				//运算
				switch char {
				case 'Q':
					if x < 0 {
						return 0, errors.New("开方错误(小于零)")
					}
					tmp := exp[i+1 : len(exp)-1]
					exp = append(exp[:i], math.Sqrt(x))
					exp = append(exp, tmp...)
				case 'N':
					tmp := exp[i+1 : len(exp)-1]
					exp = append(exp[:i], -x)
					exp = append(exp, tmp...)
				}
			} else if GetOperationFactorNum(char) == 2 {
				//找两个终结符
				x := 0.0
				num1 := 0.0
				num2 := 0.0
				for i := len(exp) - 1; i > 0; i-- {
					switch exp[i].(type) {
					case float64:
						num2 = exp[i].(float64)
						break
					}
					for j := i - 1; j > 0; j-- {
						switch exp[i].(type) {
						case float64:
							num1 = exp[j].(float64)
							break
						}
						break
					}
					break
				}
				switch char {
				case '+':
					x = num1 + num2
				case '-':
					x = num1 - num2
				case '*':
					x = num1 * num2
				case '/':
					if num2 == 0 {
						return 0, errors.New("除零错误")
					}
					x = num1 / num2
				}
				tmp := exp[i+1 : len(exp)-2]
				exp = append(exp[:i], x)
				exp = append(exp, tmp...)
			}
		default:
		}
	}
	return exp[0].(float64), nil
}
