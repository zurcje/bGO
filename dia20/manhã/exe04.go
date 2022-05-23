package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func getFunc(tipoCalculo string) (func(values ...float64) (float64, error), error) {

	if tipoCalculo == minimum {
		return calcMin, nil
	}

	if tipoCalculo == maximum {
		return calcMax, nil
	}

	if tipoCalculo == average {
		return calcAvg, nil
	}

	return nil, nil
}

func main() {

	// Chamando Direto
	// calcMin(5, 6, 7, 8, 9, 1, 2, 3, 4)
	// calcMax(5, 6, 7, 8, 9, 1, 2, 3, 4)
	// calcAvg(5, 6, 7, 8, 9, 1, 2, 3, 4)

	avgFunc, _ := getFunc(average)
	maxFunc, _ := getFunc(maximum)
	minFunc, _ := getFunc(minimum)

	min, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

	avg, _ := avgFunc(2, 3, 3, 4, 10, 2, 4, 5)

	max, _ := maxFunc(2, 3, 3, 4, 10, 2, 4, 5)
	
	fmt.Println(min, avg, max)
}

func calcMin(values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	//5 -> values[0] -> o 5 É O MENOR VALOR
	min := values[0]

	for i := 1; i < len(values); i++ {

		if values[i] < min {

			min = values[i]
		}
	}
	return min, nil
}

func calcMax(values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max, nil
}

func calcAvg(values ...float64) (float64, error) {

	if len(values) == 0 {
		return 0.0, errors.New("input is required")
	}

	avg := 0.0
	for _, v := range values {
		avg += v
	}

	return avg / float64(len(values)), nil

}