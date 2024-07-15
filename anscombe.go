package main

import (
	"fmt"
	"stats"
	"time"
)

func anscombe_1() []float64 {
	data := []stats.Coordinate{
		{10, 8.04},
		{8, 6.95},
		{13, 7.58},
		{9, 8.81},
		{11, 8.33},
		{14, 9.96},
		{6, 7.24},
		{4, 4.26},
		{12, 10.84},
		{7, 4.82},
		{5, 5.68},
	}

	r, _ := stats.LinearRegression(data)

	// Extracting and collecting Y values
	var yValues []float64
	for _, dp := range r {
		yValues = append(yValues, dp.Y)
	}
	return (yValues)
}
func anscombe_2() []float64 {
	data := []stats.Coordinate{
		{10, 9.14},
		{8, 8.14},
		{13, 8.74},
		{9, 8.77},
		{11, 9.26},
		{14, 8.1},
		{6, 6.13},
		{4, 3.1},
		{12, 9.13},
		{7, 7.26},
		{5, 4.74},
	}

	r, _ := stats.LinearRegression(data)

	// Extracting and collecting Y values
	var yValues []float64
	for _, dp := range r {
		yValues = append(yValues, dp.Y)
	}
	return (yValues)
}
func anscombe_3() []float64 {
	data := []stats.Coordinate{
		{10, 7.46},
		{8, 6.77},
		{13, 12.74},
		{9, 7.11},
		{11, 7.81},
		{14, 8.84},
		{6, 6.08},
		{4, 5.39},
		{12, 8.15},
		{7, 6.42},
		{5, 5.73},
	}

	r, _ := stats.LinearRegression(data)

	// Extracting and collecting Y values
	var yValues []float64
	for _, dp := range r {
		yValues = append(yValues, dp.Y)
	}
	return (yValues)
}
func anscombe_4() []float64 {
	data := []stats.Coordinate{
		{8, 6.58},
		{8, 5.76},
		{8, 7.71},
		{8, 8.84},
		{8, 8.47},
		{8, 7.04},
		{8, 5.25},
		{19, 12.5},
		{8, 5.56},
		{8, 7.91},
		{8, 6.89},
	}

	r, _ := stats.LinearRegression(data)

	// Extracting and collecting Y values
	var yValues []float64
	for _, dp := range r {
		yValues = append(yValues, dp.Y)
	}
	return (yValues)
}
func main() {
	start := time.Now()

	fmt.Println(anscombe_1())
	fmt.Println(anscombe_2())
	fmt.Println(anscombe_3())
	fmt.Println(anscombe_4())

	end := time.Now()
	// Calculate the duration
	duration := end.Sub(start)
	fmt.Printf("Execution time: %v\n", duration)
}
