package chart

import (
	"fmt"
	"github.com/EricsmOOn/gep-go/gep"
	"github.com/chenjiandongx/go-echarts/charts"
	"log"
	"net/http"
	"os"
)

var Max_fitness = make([]float64, 0)

var Ava_fitness = make([]float64, 0)

var Min_fitness = make([]float64, 0)

var X_value = make([]float64, 0)

func PrintChart() {
	fmt.Println()
	fmt.Println(Max_fitness)
	fmt.Println(Ava_fitness)
	fmt.Println(Min_fitness)
}

func GetChartData(genes []*gep.Gene) {
	//if genes[0].Generation%100 != 0 {
	//	return
	//}
	max := 0.0
	min := genes[0].Fitness
	sum := 0.0
	for _, g := range genes {
		sum += g.Fitness
		if g.Fitness > max {
			max = g.Fitness
		}
		if g.Fitness < min {
			min = g.Fitness
		}
	}
	Max_fitness = append(Max_fitness, max)
	Min_fitness = append(Min_fitness, min)
	Ava_fitness = append(Ava_fitness, sum/(float64(len(genes))))
	X_value = append(X_value, float64(genes[0].Generation))
}

func Handler(w http.ResponseWriter, _ *http.Request) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "进化详细"},
		charts.ToolboxOpts{Show: true},
		charts.YAxisOpts{Scale: true})
	line.AddXAxis(X_value).
		AddYAxis("每代最大适应度", Max_fitness, charts.LineOpts{Smooth: true}).
		AddYAxis("每代平均适应度", Ava_fitness, charts.LineOpts{Smooth: true}).
		AddYAxis("每代最小适应度", Min_fitness, charts.LineOpts{Smooth: true})
	f, err := os.Create("result.html")
	if err != nil {
		log.Println(err)
	}
	line.Render(w, f) // Render 可接收多个 io.Writer 接口
}
