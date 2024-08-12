package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Dataset struct {
	x, y []float64
}

func main() {
	// Define the Anscombe datasets
	datasets := []Dataset{
		{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}},
		{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}},
		{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}},
		{[]float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}, []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}},
	}

	for i, data := range datasets {
		// Perform linear regression
		alpha, beta := stat.LinearRegression(data.x, data.y, nil, false)
		yPred := make([]float64, len(data.x))
		for i := range data.x {
			yPred[i] = alpha + beta*data.x[i]
		}

		rSquared := stat.RSquared(data.y, yPred, nil)
		fmt.Printf("Dataset %d: alpha=%.2f, beta=%.2f, R^2=%.2f\n", i+1, alpha, beta, rSquared)

		// Create scatter plot
		p, err := plot.New()
		if err != nil {
			panic(err)
		}

		p.Title.Text = fmt.Sprintf("Set %d", i+1)
		p.X.Label.Text = fmt.Sprintf("x%d", i+1)
		p.Y.Label.Text = fmt.Sprintf("y%d", i+1)
		p.X.Min, p.X.Max = 2, 20
		p.Y.Min, p.Y.Max = 2, 14

		scatterData := make(plotter.XYs, len(data.x))
		for j := range data.x {
			scatterData[j].X = data.x[j]
			scatterData[j].Y = data.y[j]
		}

		s, err := plotter.NewScatter(scatterData)
		if err != nil {
			panic(err)
		}

		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("set_%d.png", i+1)); err != nil {
			panic(err)
		}
	}
}
