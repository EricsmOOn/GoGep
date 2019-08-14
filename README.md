# gep-go
 💡 **用Golang实现的GEP *[基因表达式编程]* 算法**
# 样例
**通过数据:**

![](https://github.com/EricsmOOn/gep-go/blob/master/pic/example2.png)

**计算得出:**

![](https://github.com/EricsmOOn/gep-go/blob/master/pic/example.png)

**经过化简式子后可得到原始式子**

# 使用说明
 ```
 参数列表
    //头部长度
	headLength := 7
	//基因产生个数
	populationsSize := 20
	//选择范围
	mM := float64(100)
	//染色体含有基因数
	numOfGenes := 3
	//连接函数
	connectFun := byte('+')
	//函数集
	funSet := []byte{'+', '-', '*', '/'}
	//终点集
	termSet := []byte{'a'}
	//最大操作数(参数个数)
	maxFactorNum := gep.GetMaxFactorNum(funSet)
	//基因尾部长度
	tailLength := headLength*(maxFactorNum-1) + 1
	//基因长度
	geneLength := headLength + tailLength
 ```
 ```
 函数使用
   func test(...) : 
   根据参数进行一轮基因生成,并计算出每一个式子的适应度

   func findFintness(minFitness, maxFitness, num,...) : 
   minFitness为循环选取适应度的最小值
   maxFitness为循环选取适应度的最大值
   num为循环选取符合条件的个数
 ```
