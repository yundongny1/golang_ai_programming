package main

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// Sample data for testing
var datasets = []Dataset{
	{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}},
	{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}},
	{[]float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}},
	{[]float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}, []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}},
}

// TestLinearRegression tests the linear regression results
func TestLinearRegression(t *testing.T) {
	expectedAlphas := []float64{3.00, 3.00, 3.00, 3.00}
	expectedBetas := []float64{0.50, 0.50, 0.50, 0.50}

	for i, data := range datasets {
		alpha, beta := stat.LinearRegression(data.x, data.y, nil, false)
		if !floatEquals(alpha, expectedAlphas[i], 0.01) {
			t.Errorf("Dataset %d: expected alpha=%.2f, got %.2f", i+1, expectedAlphas[i], alpha)
		}
		if !floatEquals(beta, expectedBetas[i], 0.01) {
			t.Errorf("Dataset %d: expected beta=%.2f, got %.2f", i+1, expectedBetas[i], beta)
		}
	}
}

// floatEquals checks if two float64 numbers are approximately equal
func floatEquals(a, b, tolerance float64) bool {
	return (a-b) < tolerance && (b-a) < tolerance
}

// BenchmarkLinearRegression benchmarks the linear regression computation
func BenchmarkLinearRegression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, data := range datasets {
			stat.LinearRegression(data.x, data.y, nil, false)
		}
	}
}

// BenchmarkPlotGeneration benchmarks the plot generation
func BenchmarkPlotGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, data := range datasets {
			p := plot.New()

			p.Title.Text = fmt.Sprintf("Set %d", j+1)
			p.X.Label.Text = fmt.Sprintf("x%d", j+1)
			p.Y.Label.Text = fmt.Sprintf("y%d", j+1)
			p.X.Min, p.X.Max = 2, 20
			p.Y.Min, p.Y.Max = 2, 14

			scatterData := make(plotter.XYs, len(data.x))
			for k := range data.x {
				scatterData[k].X = data.x[k]
				scatterData[k].Y = data.y[k]
			}

			s, err := plotter.NewScatter(scatterData)
			if err != nil {
				b.Fatal(err)
			}

			p.Add(s)
			if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("set_%d.png", j+1)); err != nil {
				b.Fatal(err)
			}
		}
	}
}
