# gep-go
 💡 **用Golang实现的GEP *[基因表达式编程]* 算法**
# 数据集

[testData.go文件](https://github.com/EricsmOOn/gep-go/blob/master/gep/testData.go)

# 参数列表
 ```
 global.go

 //头部长度
 var HeadLength

 //基因产生个数
 var PopulationsSize

 //选择范围
 var SelectRang

 //选择精度
 var Precision

 //染色体含有基因数
 var NumOfGenes

 //连接函数
 var LinkFun

 //函数集
 var FunSet

 //终点集
 var TermSet

 //直接变异率
 var DcMutationRate

 //单点重组率
 var OnePointRecombinationRate

 //双点重组率
 var TwoPointRecombinationRate

 //重组率
 var RecombinationRate

 //插入转座率
 var ISTranspositionRate

 //插入转座元素长度
 var ISElementsLength

 //根转座率
 var RISTranspositionRate

 //根转座元素长度
 var RISElementsLength

 //转座概率
 var GeneTranspositionRate

 //最大操作数(参数个数)
 var MaxFactorNum

 //基因尾部长度
 var TailLength

 //基因长度
 var GeneLength

 //选择结果
 var ResultRang

 //精英策略个数
 var EliteNum

 //不变异精英策略个数
 var NonEliteNum

 //转盘赌个数
 var TurnNum

 ```
 # 进化图解
 ```
  进化统计图使用
  程序找到最优解后通过浏览器进入 http://localhost:8081/
  * 显示进化中的适应度变化趋势
 ```
