package main

import (
	"brightentompkins.com/helpers"
	chart "github.com/wcharczuk/go-chart/v2"
	"os"
)

func x(t float64) float64 {
	return t*t - 2*t
}

func y(t float64) float64 {
	return t + 1
}

func a() {
	TS := helpers.Smooth(-2, 4, 1000)
	XValues := func() []float64 {
		XValues := make([]float64, len(TS))
		for i, t := range TS {
			XValues[i] = x(t)
		}

		return XValues
	}
	YValues := func() []float64 {
		YValues := make([]float64, len(TS))
		for i, t := range TS {
			YValues[i] = y(t)
		}

		return YValues
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: XValues(),
				YValues: YValues(),
			},
		},
	}

	f, _ := os.Create("calc_10_1_example_1a.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func b() {
	TS := helpers.Smooth(-2, 4, 1000)
	XValues := func() []float64 {
		XValues := make([]float64, len(TS))
		for i, t := range TS {
			if t >= 0 && t <= 4 {
				XValues[i] = x(t)
			}
		}

		return XValues
	}
	YValues := func() []float64 {
		YValues := make([]float64, len(TS))
		for i, t := range TS {
			if t >= 0 && t <= 4 {
				YValues[i] = y(t)
			}
		}

		return YValues
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: XValues(),
				YValues: YValues(),
			},
		},
	}

	f, _ := os.Create("calc_10_1_example_1b.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

func main() {
	a()
	b()
}
