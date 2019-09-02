package chart

import (
	"fmt"
	"github.com/EricsmOOn/gep-go/gep"
	"github.com/wcharczuk/go-chart"
	"net/http"
)

var Max_fitness = make([]float64, 0)

var Ava_fitness = make([]float64, 0)

var X_value = make([]float64, 0)

func PrintChart() {
	fmt.Println()
	fmt.Println(Max_fitness)
	fmt.Println(Ava_fitness)
}

func GetChartData(genes []*gep.Gene) {
	max := 0.0
	sum := 0.0
	for _, g := range genes {
		sum += g.Fitness
		if g.Fitness > max {
			max = g.Fitness
		}
	}
	Max_fitness = append(Max_fitness, max)
	Ava_fitness = append(Ava_fitness, sum/(float64(len(genes))))
	X_value = append(X_value, float64(genes[0].Generation))
}

func DrawChart(res http.ResponseWriter, req *http.Request) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.StyleShow(), //enables / displays the x-axis
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(), //enables / displays the y-axis
		},
		//Max
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: X_value,
				YValues: Max_fitness,
			},
			//Ava
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: X_value,
				YValues: Ava_fitness,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}
