package gep

/*
 A R I M A : 差 分 整 合 移 动 自 回 归
 时间序列预测:
	 //1.改进的适应度函数<拟合程度参与适应度函数计算>
	 //2.误差分析函数
	 //3.预测新的一组数据集
	 4.避免早熟收敛
		I.  增大变异率
		II. 增大种群数量
		III.在进化一定阶段时引入新个体<种群年龄分层>
*/
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
//CSV样本集文件名称(请放置于根目录下)
//var CsvSampleFileName = "vegetables_sample.csv"
var CsvSampleFileName = "sunspots_sample.csv"

//var CsvSampleFileName = "a3_sample.csv"

//控制台输出方式 Detailed - 详细,Simple - 简略,Simplest - 最简略
var ViewStyle = Simplest

//开启 10*10 交叉检验
var TenCheck = Open

//开启生成时函数符比例增多
var MoreFunc = Close

//开启遗传记录图表
var Chart = Close

//图表记录跨度 0为每次变异记录
var ChartInterval = 0

//图表端口
var ChartPort = 8081

//最高运行代数 0 - 不限制
var MaxGenerations = 8000

//适应度函数选择
var FitnessFunc = 0

/*
	GEP参数配置
*/
//头部长度 7 6 简单/a3
var HeadLength = 7

//基因产生个数 20 30
var PopulationsSize = 100

//选择范围
var SelectRang float64 = 1000

//样本结果平均数
var ResultSampleAvg = 0.0

//选择精度
var Precision float64 = 0

//染色体含有基因数 3 4
var NumOfGenes = 4

//连接函数
var LinkFun = byte('+')

//函数集
var FunSet = []byte{'+', '-', '*', '/', '@'} //,'@','$','N','Q'

//终点集
var TermSet = TermSetAll

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
var ISElementsLength = 3

//根转座率
var RISTranspositionRate = 0.1

//根转座元素长度
var RISElementsLength = 3

//转座概率
var GeneTranspositionRate = 0.1

//最大操作数(参数个数)
var MaxFactorNum = GetMaxFactorNum()

//基因尾部长度
var TailLength = HeadLength*(MaxFactorNum-1) + 1

//基因长度
var GeneLength = HeadLength + TailLength

//选择结果
var ResultRang = 0.0

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
var NonEliteNum = 2

//转盘赌个数
var TurnNum = PopulationsSize - EliteNum - NonEliteNum
