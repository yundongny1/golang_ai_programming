package main

import (
	"math"
	"testing"
)

//we're going to need to test a few things:
//test that we get similar results to results from Python and R

//test if input is anascombe dataset

func TestAnscombe1(t *testing.T) {
	python_result := []float64{8.001, 7.00081818, 9.50127273, 7.50090909, 8.50109091,
		10.00136364, 6.00063636, 5.00045455, 9.00118182, 6.50072727, 5.50054545}

	r_result := []float64{8.001000, 7.000818, 9.501273, 7.500909, 8.501091,
		10.001364, 6.000636, 5.000455, 9.001182, 6.500727, 5.500545}

	go_result := anscombe_1()

	if len(python_result) != len(go_result) {
		t.Errorf("Length of Python and Go slices are not the same!")
	} else if len(r_result) != len(go_result) {
		t.Errorf("Length of R and Go slices are not the same!")
	}

	//tolerance for comparing float values
	tolerance := 1e-6

	//compare go and python results
	for i := range python_result {
		if math.Abs(python_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, Python resulted: %v", go_result[i], python_result[i])
		}
		if math.Abs(r_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, R resulted: %v", go_result[i], r_result[i])
		}
	}
}
func TestAnscombe2(t *testing.T) {
	python_result := []float64{8.00090909, 7.00090909, 9.50090909, 7.50090909,
		8.50090909, 10.00090909, 6.00090909, 5.00090909, 9.00090909, 6.50090909, 5.50090909}

	r_result := []float64{8.000909, 7.000909, 9.500909, 7.500909, 8.500909,
		10.000909, 6.000909, 5.000909, 9.000909, 6.500909, 5.500909}

	go_result := anscombe_2()

	if len(python_result) != len(go_result) {
		t.Errorf("Length of Python and Go slices are not the same!")
	} else if len(r_result) != len(go_result) {
		t.Errorf("Length of R and Go slices are not the same!")
	}

	//tolerance for comparing float values
	tolerance := 1e-6

	//compare go and python results
	for i := range python_result {
		if math.Abs(python_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, Python resulted: %v", go_result[i], python_result[i])
		}
		if math.Abs(r_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, R resulted: %v", go_result[i], r_result[i])
		}
	}
}
func TestAnscombe3(t *testing.T) {
	python_result := []float64{7.99972727, 7.00027273, 9.49890909, 7.5,
		8.49945455, 9.99863636, 6.00081818, 5.00136364, 8.99918182, 6.50054545, 5.50109091}

	r_result := []float64{7.999727, 7.000273, 9.498909, 7.500000, 8.499455,
		9.998636, 6.000818, 5.001364, 8.999182, 6.500545, 5.501091}

	go_result := anscombe_3()

	if len(python_result) != len(go_result) {
		t.Errorf("Length of Python and Go slices are not the same!")
	} else if len(r_result) != len(go_result) {
		t.Errorf("Length of R and Go slices are not the same!")
	}

	//tolerance for comparing float values
	tolerance := 1e-6

	//compare go and python results
	for i := range python_result {
		if math.Abs(python_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, Python resulted: %v", go_result[i], python_result[i])
		}
		if math.Abs(r_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, R resulted: %v", go_result[i], r_result[i])
		}
	}
}
func TestAnscombe4(t *testing.T) {
	python_result := []float64{7.001, 7.001, 7.001, 7.001, 7.001, 7.001, 7.001, 12.5, 7.001, 7.001, 7.001}

	r_result := []float64{7.001, 7.001, 7.001, 7.001, 7.001, 7.001, 7.001, 12.500, 7.001, 7.001, 7.001}

	go_result := anscombe_4()

	if len(python_result) != len(go_result) {
		t.Errorf("Length of Python and Go slices are not the same!")
	} else if len(r_result) != len(go_result) {
		t.Errorf("Length of R and Go slices are not the same!")
	}

	//tolerance for comparing float values
	tolerance := 1e-6

	//compare go and python results
	for i := range python_result {
		if math.Abs(python_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, Python resulted: %v", go_result[i], python_result[i])
		}
		if math.Abs(r_result[i]-go_result[i]) > tolerance {
			t.Errorf("Got %v, R resulted: %v", go_result[i], r_result[i])
		}
	}
}
