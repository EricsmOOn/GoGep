package chart

import (
	"github.com/EricsmOOn/GoGep/gep"
	"github.com/chenjiandongx/go-echarts/charts"
	"log"
	"math"
	"net/http"
	"os"
)

var MaxFitness = make([]float64, 0)

var AvaFitness = make([]float64, 0)

var XValue = make([]float64, 0)

var XValueTestPred = make([]float64, 0)

var YValueTestPred []float64

var YValueTestPredError = make([]float64, 0)

var XValueSamplePred = make([]float64, 0)

var YValueSamplePred []float64

var YValueSamplePredError = make([]float64, 0)

var Gene gep.Gene

var MaxPrinter = 0.0

func GetPredictResult(genes gep.Gene) {
	Gene = genes
	data := gep.ReadTestData()
	for i := 1; i <= len(data); i++ {
		XValueTestPred = append(XValueTestPred, float64(i+9))
	}
	YValueTestPred = gep.GetPredictTestResult(gep.GetEffectGenes(Gene))
	for i, r := range YValueTestPred {
		if data[i].Result != 0 {
			YValueTestPredError = append(YValueTestPredError, (math.Abs(data[i].Result-r)/data[i].Result)*100)
		} else {
			YValueTestPredError = append(YValueTestPredError, math.Abs(data[i].Result-r)*100)
		}
	}
	data = gep.ReadSampleData()
	for i := 1; i <= len(data); i++ {
		XValueSamplePred = append(XValueSamplePred, float64(i))
	}
	YValueSamplePred = gep.GetPredictSampleResult(gep.GetEffectGenes(Gene))
	for i, r := range YValueSamplePred {
		if data[i].Result != 0 {
			YValueSamplePredError = append(YValueSamplePredError, (math.Abs(data[i].Result-r)/data[i].Result)*100)
		} else {
			YValueSamplePredError = append(YValueSamplePredError, math.Abs(data[i].Result-r)*100)
		}
		//if data[i].Result == 0{
		//	YValueSamplePredError = append(YValueSamplePredError,data[i].Result-r)
		//}
		//f, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Abs(data[i].Result-r)/data[i].Result), 64)
		//YValueSamplePredError = append(YValueSamplePredError,f)
	}
}

func GetChartData(genes []*gep.Gene) {
	var interval int
	if gep.ChartInterval == 0 {
		max := 0.0
		sum := 0.0
		for _, g := range genes {
			sum += g.Fitness
			if g.Fitness > max {
				max = g.Fitness
			}
		}
		if max > MaxPrinter {
			MaxPrinter = max
			MaxFitness = append(MaxFitness, max)
			AvaFitness = append(AvaFitness, sum/(float64(len(genes))))
			XValue = append(XValue, float64(genes[0].Generation))
		}
		return
	} else {
		interval = gep.ChartInterval
	}
	if genes[0].Generation%interval == 0 {
		max := 0.0
		sum := 0.0
		for _, g := range genes {
			sum += g.Fitness
			if g.Fitness > max {
				max = g.Fitness
			}
		}
		MaxFitness = append(MaxFitness, max)
		AvaFitness = append(AvaFitness, sum/(float64(len(genes))))
		XValue = append(XValue, float64(genes[0].Generation))
	}
}

func lineBase1() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "进化详细"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "进化代数"})
	line.AddXAxis(XValue).
		AddYAxis("每代最大适应度", MaxFitness, charts.LineOpts{Smooth: false}).
		AddYAxis("每代平均适应度", AvaFitness, charts.LineOpts{Smooth: false})
	return line
}

func lineBase2() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "预测测试集"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "样本"})
	line.AddXAxis(XValueTestPred).AddYAxis("测试值", gep.GetTestResult(), charts.LineOpts{Smooth: false}).
		AddYAxis("预测值", YValueTestPred, charts.LineOpts{Smooth: false})

	line2 := charts.NewLine()
	line2.AddXAxis(XValueTestPred).AddYAxis("误差值", YValueTestPredError)
	line2.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "平均值", Type: "average"},
		charts.LineOpts{Smooth: false},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"}},
	)
	line.Overlap(line2)
	return line
}

func lineBase3() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "预测样本集"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "样本"})
	line.AddXAxis(XValueSamplePred).AddYAxis("样本值", gep.GetSampleResult(), charts.LineOpts{Smooth: false}).
		AddYAxis("预测值", YValueSamplePred, charts.LineOpts{Smooth: false})

	line2 := charts.NewLine()
	line2.AddXAxis(XValueSamplePred).AddYAxis("误差值", YValueSamplePredError)
	line2.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "平均值", Type: "average"},
		charts.LineOpts{Smooth: false},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"}},
	)
	line.Overlap(line2)
	return line
}

func Handler(w http.ResponseWriter, _ *http.Request) {
	p := charts.NewPage()
	if gep.Chart {
		p.Add(lineBase1(), lineBase2(), lineBase3())
	} else {
		p.Add(lineBase2(), lineBase3())
	}
	f, err := os.Create("page.html")
	if err != nil {
		log.Println(err)
	}
	err = p.Render(w, f)
	if err != nil {
		log.Println(err)
	}
}
