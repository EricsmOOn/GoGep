package gep

const (
	Open     = true
	Close    = false
	Detailed = 0
	Simple   = 1
	Simplest = 2
)

var TermSetAll = []byte{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

/*
	这里是全局参数控制中心
*/

/*
	运行参数配置
*/
//控制台输出方式 Detailed - 详细,Simple - 简略,Simplest - 最简略
var ViewStyle = Simplest

//开启图表
var Chart = Open

//图表端口
var ChartPort = 8081

//图表记录跨度 0为每次变异记录
var ChartInterval = 0

//最高运行代数 0 - 不限制
var MaxGenerations = 0

//函数计时器开关
var FuncTimer = false

/*
	GEP参数配置
*/
//头部长度 7 6 简单/a3
var HeadLength = 7

//基因产生个数 20 30
var PopulationsSize = 100

//选择范围
var SelectRang float64 = 1000

//选择精度
var Precision float64 = 800

//染色体含有基因数 3 4
var NumOfGenes = 3

//连接函数
var LinkFun = byte('+')

//函数集
var FunSet = []byte{'+', '-', '*', '/'}

//终点集
var TermSet = TermSetAll[:GetVarSetNum()]

//直接变异率 0.144 0.0385
var DcMutationRate = 0.0385

//单点重组率 0.4 0.3
var OnePointRecombinationRate = 0.3

//双点重组率 0.2 0.3
var TwoPointRecombinationRate = 0.3

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

/*
	高阶GEP参数配置
*/
//自适应变异率开关
var DynamicDcMutation = Open

//自适应变异率基数
var DynamicDcMutationRate = 1.0

//精英策略个数
var EliteNum = 0

//不变异精英策略个数
var NonEliteNum = 5

//转盘赌个数
var TurnNum = PopulationsSize - EliteNum - NonEliteNum
