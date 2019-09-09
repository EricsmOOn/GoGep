package chart

import (
	"github.com/EricsmOOn/GoGep/gep"
	"github.com/chenjiandongx/go-echarts/charts"
	"log"
	"net/http"
	"os"
)

var MaxFitness = make([]float64, 0)

var AvaFitness = make([]float64, 0)

var XValue = make([]float64, 0)

var XValueTestPred = make([]float64, 0)

var XValueSamplePred = make([]float64, 0)

var Gene gep.Gene

var MaxPrinter = 0.0

func GetPredictResult(genes gep.Gene) {
	Gene = genes
	num := gep.GetTestDataNum()
	for i := 1; i <= num; i++ {
		XValueTestPred = append(XValueTestPred, float64(i))
	}
	num = gep.GetSampleDataNum()
	for i := 1; i <= num; i++ {
		XValueSamplePred = append(XValueSamplePred, float64(i))
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
		AddYAxis("每代最大适应度", MaxFitness, charts.LineOpts{Smooth: true}).
		AddYAxis("每代平均适应度", AvaFitness, charts.LineOpts{Smooth: true})
	return line
}

func lineBase2() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "预测测试集"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "样本"})
	line.AddXAxis(XValueTestPred).AddYAxis("测试值", gep.GetTestResult(), charts.LineOpts{Smooth: true}).
		AddYAxis("预测值", gep.GetPredictTestResult(gep.GetEffectGenes(Gene)), charts.LineOpts{Smooth: true})
	//line.SetSeriesOptions(
	//	charts.MLNameTypeItem{Name: "平均值", Type: "average"},
	//	charts.LineOpts{Smooth: true},
	//	charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"}},
	//)
	return line
}

func lineBase3() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "预测样本集"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "样本"})
	line.AddXAxis(XValueSamplePred).AddYAxis("样本值", gep.GetSampleResult(), charts.LineOpts{Smooth: true}).
		AddYAxis("预测值", gep.GetPredictSampleResult(gep.GetEffectGenes(Gene)), charts.LineOpts{Smooth: true})
	//line.SetSeriesOptions(
	//	charts.MLNameTypeItem{Name: "平均值", Type: "average"},
	//	charts.LineOpts{Smooth: true},
	//	charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"}},
	//)
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
