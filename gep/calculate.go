package gep

import (
	"errors"
	"math"
	"sync"
)

var Wg sync.WaitGroup

//计算种群适应度
func CalculateFitnessOpt(genes []*Gene) {
	for _, gene := range genes {
		Wg.Add(1)
		switch FitnessFunc {
		case 0: //普通
			go calculateFitOpt(gene, GetEffectGenes(*gene))
		case 1: //百分比
			go calculateFitOptOther(gene, GetEffectGenes(*gene))
		case 2: //改进普通
			go calculateFitOpt2(gene, GetEffectGenes(*gene))
		case 3:
			go calculateFitOpt3(gene, GetEffectGenes(*gene))
		default:
		}
	}
}

//简化式子后计算个体适应度
func calculateFitOpt(gene *Gene, easyEquation [][]byte) {
	defer Wg.Done()
	testData := ReadSampleData()
	fsum := 0.0
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, err := calculatePerFitOpt(easyEquation, td.TermVarSet)
		//杀死非法表达式
		if err != nil {
			gene.Fitness = 0
			break
		}
		fi := SelectRang - math.Abs(result-td.Result)
		if fi > 0 {
			fsum += fi
		}
	}
	gene.Fitness = fsum
}

//简化式子后计算个体适应度
func calculateFitOpt3(gene *Gene, easyEquation [][]byte) {
	defer Wg.Done()
	testData := ReadSampleData()
	fsum := 0.0
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, err := calculatePerFitOpt(easyEquation, td.TermVarSet)
		//杀死非法表达式
		if err != nil {
			gene.Fitness = 0
			break
		}
		b := 1.0
		if td.Result == 0 {
			b = math.Abs(result - td.Result)
		} else {
			b = math.Abs(result-td.Result) / td.Result
		}
		fi := SelectRang - math.Abs(result-td.Result)*b
		if fi > 0 {
			fsum += fi
		}
	}
	gene.Fitness = fsum
}

//简化式子后计算个体适应度R^2改进
func calculateFitOpt2(gene *Gene, easyEquation [][]byte) {
	defer Wg.Done()
	testData := ReadSampleData()
	fsum := 0.0
	eiup := 0.0
	eidown := 0.0
	ei := 0.0
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, err := calculatePerFitOpt(easyEquation, td.TermVarSet)
		//杀死非法表达式
		if err != nil {
			gene.Fitness = 0
			break
		}
		eValue := math.Abs(result - td.Result)
		eiup += math.Pow(eValue, 2)
		eidown += math.Pow(result-ResultSampleAvg, 2)
		fi := SelectRang - eValue
		if fi > 0 {
			fsum += fi
		}
	}
	ei = eiup / eidown
	gene.Fitness = fsum * (1 - ei)
}

//简化式子后计算个体适应度使用特殊方法
func calculateFitOptOther(gene *Gene, easyEquation [][]byte) {
	defer Wg.Done()
	testData := ReadSampleData()
	eiup := 0.0
	eidown := 0.0
	ei := 0.0
	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result, err := calculatePerFitOpt(easyEquation, td.TermVarSet)
		//杀死非法表达式
		if err == nil {
			gene.Fitness = 0
		}
		eiup += math.Pow(result-td.Result, 2)
		eidown += math.Pow(result-ResultSampleAvg, 2)
	}
	ei = math.Sqrt(eiup / eidown)
	gene.Fitness = SelectRang * 1 / (1 + ei)
}

//计算个体对一条样例的适应度前的预操作
func calculatePerFitOpt(g [][]byte, v []float64) (float64, error) {
	var result float64
	for i := 0; i < len(g); i++ {
		//转换基因到 slice
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
				flag := true
				for i := len(exp) - 1; flag; i-- {
					switch exp[i].(type) {
					case float64:
						x = exp[i].(float64)
						flag = false
					}
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
				case '$':
					tmp := exp[i+1 : len(exp)-1]
					exp = append(exp[:i], math.Sin(x))
					exp = append(exp, tmp...)
				case '@':
					tmp := exp[i+1 : len(exp)-1]
					exp = append(exp[:i], math.Cos(x))
					exp = append(exp, tmp...)
				}
			} else if GetOperationFactorNum(char) == 2 {
				//找两个终结符
				x := 0.0
				num1 := 0.0
				num2 := 0.0
				flag := true
				var j int
				for i := len(exp) - 1; flag; i-- {
					switch exp[i].(type) {
					case float64:
						num2 = exp[i].(float64)
						flag = false
					}
					j = i
				}
				flag = true
				for j = j - 1; flag; j-- {
					switch exp[j].(type) {
					case float64:
						num1 = exp[j].(float64)
						flag = false
					}
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
