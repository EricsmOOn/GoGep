package gep

//头部长度
var HeadLength = 7

//基因产生个数
var PopulationsSize = 20

//选择范围
var SelectRang = float64(100)

//选择精度
var Precision = 0.1

//染色体含有基因数
var NumOfGenes = 3

//连接函数
var LinkFun = byte('+')

//函数集
var FunSet = []byte{'+', '-', '*', '/'}

//终点集
var TermSet = []byte{'a'}

//直接变异率
var DcMutationRate = 0.144

//单点重组率
var OnePointRecombinationRate = 0.4

//双点重组率
var TwoPointRecombinationRate = 0.2

//重组率
var RecombinationRate = 0.1

//插入转座率
var ISTranspositionRate = 0.1

//插入转座元素长度
var ISElementsLength = 2

//根转座率
var RISTranspositionRate = 0.1

//根转座元素长度
var RISElementsLength = 2

//转座概率
var GeneTranspositionRate = 0.1

//最大操作数(参数个数)
var MaxFactorNum = GetMaxFactorNum()

//基因尾部长度
var TailLength = HeadLength*(MaxFactorNum-1) + 1

//基因长度
var GeneLength = HeadLength + TailLength

//选择结果
var ResultRang = SelectRang*float64(GetDataNum()) - Precision
