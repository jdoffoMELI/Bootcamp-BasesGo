package main

import (
	"fmt"
	"math"
)

type StadisticOp uint8

const (
	minimo StadisticOp = iota
	maximo
	promedio
)

func main() {
	minFunc := operation(minimo)
	averageFunc := operation(promedio)
	maxFunc := operation(maximo)

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("Valor minimo: %.2f \nValor Maximo: %.2f\nPromedio: %.2f\n", minValue, maxValue, averageValue)
}

func operation(operator StadisticOp) func(values ...float32) float32 {
	switch operator {
	case minimo:
		return minFunc
	case maximo:
		return maxFunc
	case promedio:
		return meanFunc
	default:
		return nil
	}
}

func minFunc(values ...float32) float32 {
	var minimo float32 = math.MaxFloat32
	for _, value := range values {
		if minimo > value {
			minimo = value
		}
	}
	return minimo
}

func maxFunc(values ...float32) float32 {
	var maximo float32 = math.SmallestNonzeroFloat32
	for _, value := range values {
		if maximo < value {
			maximo = value
		}
	}
	return maximo
}

func meanFunc(values ...float32) float32 {
	var media float32 = 0
	for _, value := range values {
		media += value
	}
	return media / float32(len(values))
}
