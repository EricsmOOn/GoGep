# gep-go
   [![Build Status](https://travis-ci.org/EricsmOOn/gep-go.svg?branch=master)](https://travis-ci.org/EricsmOOn/gep-go) [![Go Report Card](https://goreportcard.com/badge/github.com/EricsmOOn/gep-go)](https://goreportcard.com/report/github.com/EricsmOOn/gep-go)
 
💡 **用Golang实现的GEP *[基因表达式编程]* 算法**

GEP(Gene Expression Programming,基因表达式编程) —— 通过人工智能建立的数学模型。这是一种新的进化算法，它可以进化诸如数学表达式、神经网络、决策树、多项式构建、逻辑表达式等多种形态的计算机程序。在过去的科研成果中，该技术已被用于公式发现、函数挖掘、关联规则挖掘、因子分解、太阳黑子预测，并且取得了满意的效果。
# 数据集

ℹ️ Todo

# 参数列表

## 运行参数配置
```
ViewStyle - 控制台输出方式(Detailed - 详细,Simple - 简略,Simplest - 最简略)

Chart - 图表开关

ChartPort - 图表展示端口号

ChartInterval - 图表记录跨度(0 - 每次变异触发记录)

MaxGenerations - 最高运行代数(0 - 不限制)

FuncTimer - 函数计时器开关
```

## GEP基本参数配置
```
HeadLength - 头部长度

PopulationsSize - 基因产生个数

SelectRang - 选择范围

Precision - 选择精度

NumOfGenes - 每条染色体含有基因数

LinkFun - 连接函数('+', '-', '*', '/')

FunSet - 函数集{'+', '-', '*', '/'}

TermSet - 终点集(默认根据数据集自动生成)

DcMutationRate - 直接变异率

OnePointRecombinationRate - 单点重组率

TwoPointRecombinationRate - 双点重组率

RecombinationRate - 基因重组率

ISTranspositionRate - 插入转座率

ISElementsLength - 插入转座元素长度

RISTranspositionRate - 根转座率

RISElementsLength - 根转座元素长度

GeneTranspositionRate - 基因转座概率

MaxFactorNum - 最大操作数(默认根据函数集自动生成)

TailLength - 基因尾部长度(默认根据头部长度自动生成)

GeneLength - 基因长度(默认根据头部长度自动生成)

ResultRang - 选择结果(默认根据选择范围、数据集数据个数、选择精度自动生成)
```

## 高阶GEP参数配置
```
DynamicDcMutation - 自适应变异开关

DynamicDcMutationRate - 自适应变异率基数

EliteNum - 精英策略个数

NonEliteNum - 不变异精英策略个数

TurnNum = 转盘赌个数(默认根据EliteNum、NonEliteNum自动生成)
```

# 进化图解
 ![进化图解](https://github.com/EricsmOOn/gep-go/blob/master/pic/%E8%BF%9B%E5%8C%96%E8%AF%A6%E7%BB%86.png)
 ```
  1. 设置运行参数 Chart = Open ,根据情况调整 ChartPort、ChartInterval 。
  2. 程序找到最优解后通过浏览器进入 http://localhost:ChartPort/
 ```
# 致谢
🌟超好用的Go语言图形库 [go-echarts](https://go-echarts.chenjiandongx.com)
