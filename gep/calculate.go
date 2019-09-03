package gep

import (
	"errors"
	"github.com/EricsmOOn/gep-go/util/timer"
	"math"
	"strconv"
	"strings"
)

//计算种群适应度
func CalculateFitness(genes []*Gene) {
	defer timer.TimeCount()()
	for _, gene := range genes {

		calculateFit(gene, GetInfixExpressions(*gene))
	}
}

//简化式子后计算个体适应度
func calculateFit(gene *Gene, easyEquation [][]byte) {
	testData := ReadTestData()

	//注入中缀
	g := easyEquation
	gene.InfixExpression = g

	for _, td := range testData {
		//逐个读入测试数据
		//求表达 result
		result := calculatePerFit(g, td.TermVarSet)
		fi := SelectRang - math.Abs(result-td.Result)
		if fi > 0 {
			gene.Fitness += fi
		}
	}
}

//计算个体对一组数据的适应度
func calculatePerFit(g [][]byte, v []float64) float64 {

	var result float64
	for i := 0; i < len(g); i++ {
		slice := make([]interface{}, 0)
		for _, value := range g[i] {
			slice = append(slice, value)
		}

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

		f, e := calculate(infix2ToPostfix(slice))
		if e != nil {
			return 0
		}
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
	return result
}

// 逆波兰表达式求个体对一条样例的适应度
// 1.添加操作数 2.除以零报错 3.Q使用的是结果向下取整
func calculate(postfix string) (float64, error) {
	stack := ItemStack{}
	split := strings.Fields(postfix)

	fixLen := len(split)

	for i := 0; i < fixLen; i++ {
		nextChar := string(split[i])
		if GetOperationFactorNum([]byte(nextChar)[0]) == 1 && len(nextChar) == 1 {
			num, _ := strconv.ParseFloat(stack.Pop(), 64)
			switch nextChar {
			case "Q":
				if num < 0 {
					return 0, errors.New("开方错误(小于零)")
				}
				stack.Push(strconv.FormatFloat(math.Sqrt(num), 'f', -1, 64))
			case "N":
				stack.Push(strconv.FormatFloat(-num, 'f', -1, 64))
			}
		} else if GetOperationFactorNum([]byte(nextChar)[0]) == 2 && len(nextChar) == 1 {
			// 操作符：取出两个数字计算值，再将结果压栈
			num2, _ := strconv.ParseFloat(stack.Pop(), 64)
			num1, _ := strconv.ParseFloat(stack.Pop(), 64)
			switch nextChar {
			case "+":
				stack.Push(strconv.FormatFloat(num1+num2, 'f', -1, 64))
			case "-":
				stack.Push(strconv.FormatFloat(num1-num2, 'f', -1, 64))
			case "*":
				stack.Push(strconv.FormatFloat(num1*num2, 'f', -1, 64))
			case "/":
				if num2 == 0 {
					return 0, errors.New("除零错误")
				}
				stack.Push(strconv.FormatFloat(num1/num2, 'f', -1, 64))
			}
		} else {
			stack.Push(nextChar)
		}
	}
	result, _ := strconv.ParseFloat(stack.Top(), 64)
	return result, nil
}

// 中缀表达式转后缀表达式
func infix2ToPostfix(exp []interface{}) string {
	stack := ItemStack{}
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {
		switch exp[i].(type) {
		case byte:
			char := string(exp[i].(byte))

			switch char {
			case " ":
				continue
			case "(":
				// 左括号直接入栈
				stack.Push("(")
			case ")":
				// 右括号则弹出元素直到遇到左括号
				for !stack.IsEmpty() {
					preChar := stack.Top()
					if preChar == "(" {
						stack.Pop() // 弹出 "("
						break
					}
					postfix += preChar
					stack.Pop()
				}

			default:
				// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
				for !stack.IsEmpty() {
					top := stack.Top()
					if top == "(" || isLower(top, char) {
						break
					}
					postfix += top
					stack.Pop()
				}
				// 低优先级的运算符入栈
				stack.Push(char)
			}
		case float64:
			// 数字则直接输出
			postfix += strconv.FormatFloat(exp[i].(float64), 'f', -1, 64)
		}
		postfix += " "
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}

// 比较运算符栈栈顶 top 和新运算符 newTop 的优先级高低
func isLower(top string, newTop string) bool {

	topb := []byte(top)[0]
	newTopb := []byte(newTop)[0]

	return GetOperationPriority(topb) > GetOperationPriority(newTopb)
}
