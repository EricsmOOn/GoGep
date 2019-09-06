<h1 align="center">GoGep</h1>
<p align="center">
    <strong> 🧬 用Golang实现的GEP（Gene Expression Programming，基因表达式编程）算法 </strong>
</p>
<p align="center">
    <a href="https://travis-ci.org/EricsmOOn/GoGep">
        <img src="https://travis-ci.org/EricsmOOn/GoGep.svg?branch=master" alt="Build Status">
    </a>
    <a href="https://goreportcard.com/report/github.com/EricsmOOn/GoGep">
        <img src="https://goreportcard.com/badge/github.com/EricsmOOn/GoGep" alt="Go Report Card">
    </a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
    </a>
</p>
   
## 简介
GEP(Gene Expression Programming，基因表达式编程) —— 通过人工智能建立的数学模型。这是一种新的进化算法，它可以进化诸如数学表达式、神经网络、决策树、多项式构建、逻辑表达式等多种形态的计算机程序。在过去的科研成果中，该技术已被用于公式发现、函数挖掘、关联规则挖掘、因子分解、太阳黑子预测，并且取得了令人满意的效果。

## 📄数据集
[ - 数据集样例(太阳黑子)](https://github.com/EricsmOOn/gep-go/blob/master/sunspots.csv)

```
csv数据集文件格式请按照:
   1.每行为一组数据，每个样本值之间使用逗号进行分隔，每组样本数据之间使用换行符进行分隔。
   2.每组样本数据的最后一个样本值为实际样本结果。
   3.请将目标csv文件放置于主程序根目录下，并在运行参数里配置CsvFileName为目标csv文件名(不带文件名后缀)。
```

## 📋参数列表
`请在 ./gep/global.go文件中配置。`

### 运行参数配置
```
CsvFileName - csv数据文件名称(请放置于根目录下)

ViewStyle - 控制台输出方式(Detailed - 详细,Simple - 简略,Simplest - 最简略)

Chart - 图表开关

ChartPort - 图表展示端口号

ChartInterval - 图表记录跨度(0 - 每次变异触发记录)

MaxGenerations - 最高运行代数(0 - 不限制)
```

### GEP基本参数配置
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

ResultRang - 选择结果<修改无效>(默认根据选择范围、数据集数据个数、选择精度自动生成)
```

### 高阶GEP参数配置
```
DynamicDcMutation - 自适应变异开关

DynamicDcMutationRate - 自适应变异率基数

EliteNum - 精英策略个数

NonEliteNum - 不变异精英策略个数

TurnNum = 转盘赌个数(默认根据EliteNum、NonEliteNum自动生成)
```

## 📈进化图解
 ![进化图解](https://github.com/EricsmOOn/gep-go/blob/master/pic/%E8%BF%9B%E5%8C%96%E8%AF%A6%E7%BB%86.png)
 ```
  1. 设置运行参数 Chart = Open ,根据情况调整 ChartPort、ChartInterval 。
  2. 程序找到最优解后通过浏览器进入 http://localhost:ChartPort/
 ```
## ⭐️致谢
- 超好用的Go语言图形库 [go-echarts](https://go-echarts.chenjiandongx.com)
