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

var MaxPrinter = 0.0

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

func Handler(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "进化详细"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Name: "适应度", Scale: true, SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "进化代数"})
	line.AddXAxis(XValue).
		AddYAxis("每代最大适应度", MaxFitness, charts.LineOpts{Smooth: true}).
		AddYAxis("每代平均适应度", AvaFitness, charts.LineOpts{Smooth: true})
	//line.SetSeriesOptions(
	//	charts.MLNameTypeItem{Name: "平均值", Type: "average"},
	//	charts.LineOpts{Smooth: true},
	//	charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"}},
	//)
	f, err := os.Create("result.html")
	if err != nil {
		log.Println(err)
	}
	err = line.Render(w, f) // Render 可接收多个 io.Writer 接口
	if err != nil {
		log.Println(err)
	}
}
