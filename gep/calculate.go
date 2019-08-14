package gep

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

func Calculate(g []byte, v []byte ,termSet []byte) (int,error) {

	for gno,b := range g{
		for ts,e := range termSet{
			if e == b{
				tmp := g[gno+1:]
				g = append(g[:gno],v[ts])
				g = append(g,tmp...)
			}
		}
	}

	//fmt.Println(*(*string)(unsafe.Pointer(&g)))
	return calculate(infix2ToPostfix(string(g)))
}

//1.添加操作数 2.除以零报错 3.Q使用的是结果向下取整
func calculate(postfix string) (int,error) {
	stack := ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else if GetOperationFactorNum([]byte(nextChar)[0]) == 1{
			num, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "Q":
				if num < 0 {return 0,errors.New("开方错误(小于零)")}
				stack.Push(strconv.Itoa(int(math.Floor(math.Sqrt(float64(num))))))
			case "N":
				stack.Push(strconv.Itoa(-num))
			}
		} else if GetOperationFactorNum([]byte(nextChar)[0]) == 2{
			// 操作符：取出两个数字计算值，再将结果压栈
			num2, _ := strconv.Atoi(stack.Pop())
			num1, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "%":
				if num2 == 0{
					return 0,errors.New("除零错误")
				}
				stack.Push(strconv.Itoa(num1 % num2))
			case "/":
				if num2 == 0{
					return 0,errors.New("除零错误")
				}
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result,nil
}

// 中缀表达式转后缀表达式
func infix2ToPostfix(exp string) string {
	stack := ItemStack{}
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

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

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

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